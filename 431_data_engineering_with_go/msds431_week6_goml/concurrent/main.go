package main

//https://medium.com/devthoughts/linear-regression-with-go-ff1701455bcd

import (
	"fmt"
	"math"
	"sync"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func makePreds(model *linear.LeastSquares, testX [][]float64) []float64 {
	var predholder []float64
	for i := range testX {
		prediction, err3 := model.Predict(testX[i])
		if err3 != nil {
			panic("Error Predicting")
		}
		predholder = append(predholder, prediction[0])
	}

	return predholder
}

func sumSlice(aslice []float64) float64 {
	var holder float64
	for _, i := range aslice {
		holder += i
	}
	return holder
}

func getSSR(preds []float64, mean float64) float64 {
	var ssr float64
	for i := range preds {
		ssr += math.Pow(preds[i]-mean, 2)
	}
	return ssr

}

func getSSE(preds []float64, actuals []float64) float64 {
	var sse float64
	for i := range preds {
		sse += math.Pow(preds[i]-actuals[i], 2)
	}
	return sse

}

func makeModel(trainX [][]float64, trainY []float64, i int) ModelWithName {
	model := ModelWithName{
		linear.NewLeastSquares(base.BatchGA, 1e-10, 2, 10000, trainX, trainY),
		i}
	return model
}

func makeModelStoc(trainX [][]float64, trainY []float64, i int) ModelWithName {
	model := ModelWithName{
		linear.NewLeastSquares(base.StochasticGA, 1e-10, 2, 10000, trainX, trainY),
		i}
	return model
}

type ModelWithName struct {
	*linear.LeastSquares
	Name int
}

func main() {
	var wg sync.WaitGroup
	// load train
	trainX, trainY, err1 := base.LoadDataFromCSV("../data/training_set_noheader.csv")
	if err1 != nil {
		panic("Error Loading Training")
	}

	// load test
	testX, testY, err2 := base.LoadDataFromCSV("../data/testing_set_noheader.csv")
	if err2 != nil {
		panic("Error Loading Testing")
	}
	//BatchGA channels
	for i := 1; i < 101; i++ {
		wg.Add(1)
		// init model
		// opt method, learning rate,
		go func(i int) {
			defer wg.Done()
			model := makeModel(trainX, trainY, i)
			err := model.Learn()
			if err != nil {
				panic("Error Learning")
			}
			predholder := makePreds(model.LeastSquares, testX)
			predavg := sumSlice(predholder) / float64(len(predholder))
			ssr := getSSR(predholder, predavg)
			sse := getSSE(predholder, testY)
			sst := ssr + sse
			rsquared := 1 - (ssr / sst)
			mse := ssr / (float64(len(predholder) - 2))
			rsme := math.Sqrt(mse)
			fmt.Printf("\n************DESC %v**************", model.Name)
			fmt.Printf("rsquared = %.2f\n", rsquared)
			fmt.Printf("mse = %.2f\n", mse)
			fmt.Printf("rsme = %.2f\n", rsme)
			fmt.Printf("model num = %v", model.Name)
			fmt.Println("******************************")
		}(i)

		wg.Add(1)
		// init model
		// opt method, learning rate,
		go func(i int) {
			defer wg.Done()
			model := makeModelStoc(trainX, trainY, i)
			err := model.Learn()
			if err != nil {
				panic("Error Learning")
			}
			predholder := makePreds(model.LeastSquares, testX)
			predavg := sumSlice(predholder) / float64(len(predholder))
			ssr := getSSR(predholder, predavg)
			sse := getSSE(predholder, testY)
			sst := ssr + sse
			rsquared := 1 - (ssr / sst)
			mse := ssr / (float64(len(predholder) - 2))
			rsme := math.Sqrt(mse)
			fmt.Printf("\n************STOC %v**************", model.Name)
			fmt.Printf("rsquared = %.2f\n", rsquared)
			fmt.Printf("mse = %.2f\n", mse)
			fmt.Printf("rsme = %.2f\n", rsme)
			fmt.Printf("model num = %v", model.Name)
			fmt.Println("******************************")
		}(i)
	}
	wg.Wait()
}
