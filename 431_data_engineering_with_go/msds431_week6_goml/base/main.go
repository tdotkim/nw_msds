package main

//https://medium.com/devthoughts/linear-regression-with-go-ff1701455bcd

import (
	"fmt"
	"math"

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

func main() {
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

	for i := 0; i < 100; i++ {
		// init model
		// opt method, learning rate,
		model := linear.NewLeastSquares(base.BatchGA, 1e-10, 2, 10000, trainX, trainY)
		//train model
		err := model.Learn()
		if err != nil {
			panic("Error Learning")
		}

		predholder := makePreds(model, testX)
		predavg := sumSlice(predholder) / float64(len(predholder))
		ssr := getSSR(predholder, predavg)
		sse := getSSE(predholder, testY)
		sst := ssr + sse
		rsquared := 1 - (ssr / sst)
		mse := ssr / (float64(len(predholder) - 2))
		rsme := math.Sqrt(mse)
		fmt.Printf("rsquared = %.2f\n", rsquared)
		fmt.Printf("mse = %.2f\n", mse)
		fmt.Printf("rsme = %.2f\n", rsme)

		modeltwo := linear.NewLeastSquares(base.StochasticGA, 1e-10, 2, 10000, trainX, trainY)

		//train model
		err8 := modeltwo.Learn()
		if err8 != nil {
			panic("Error Learning")
		}
		predholder2 := makePreds(modeltwo, testX)
		predavg2 := sumSlice(predholder2) / float64(len(predholder2))
		ssr2 := getSSR(predholder2, predavg2)
		sse2 := getSSE(predholder2, testY)
		sst2 := ssr2 + sse2
		rsquared2 := 1 - (ssr2 / sst2)
		mse2 := ssr2 / (float64(len(predholder2) - 2))
		rsme2 := math.Sqrt(mse2)
		fmt.Printf("rsquared = %.2f\n", rsquared2)
		fmt.Printf("mse = %.2f\n", mse2)
		fmt.Printf("rsme = %.2f\n", rsme2)
	}
}
