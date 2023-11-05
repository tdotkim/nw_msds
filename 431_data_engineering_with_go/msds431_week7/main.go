package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/e-XpertSolutions/go-iforest/v2/iforest"
	"github.com/petar/GoMNIST"
)

func loadMNISTData() (images [][]float64, labels []int) {

	train, _, err := GoMNIST.Load("./data")
	if err != nil {
		fmt.Printf("error loading mnist data %v", err)
	}
	images = make([][]float64, len(train.Images))
	labels = make([]int, len(train.Images))

	for i := 0; i < len(train.Images); i++ {
		images[i] = make([]float64, len(train.Images[0]))
		for j := range train.Images[0] {
			images[i][j] = float64(train.Images[i][j]) //converts integer pixel values to float64 for input to iforest
			labels[i] = int(train.Labels[i])

		}
	}
	/* fmt.Println("---Length of Data---")
	fmt.Println(len(images))
	fmt.Println(len(labels)) */
	return images, labels

}
func makeScoreSliceTestable(mymap map[int]float64) (scores [][]string) {
	scores = make([][]string, len(mymap)+1)
	for i := 0; i < (len(scores)); i++ {
		scores[i] = make([]string, 2)
		if i == 0 { //adds header in csv for easy import into R dataframes
			scores[0][0] = "id"
			scores[0][1] = "scores"
		}
		if i != 0 {
			score := mymap[i-1]
			scores[i][0] = fmt.Sprintf("%d", i-1)
			scores[i][1] = fmt.Sprintf("%f", score)
		}

	}
	return scores
}

func makeScoreSlice(f *iforest.Forest) (scores [][]string) {
	scores = make([][]string, len(f.AnomalyScores)+1)
	fmt.Println(f.AnomalyScores)
	for i := 0; i < (len(scores)); i++ {
		scores[i] = make([]string, 2)
		if i == 0 { //adds header in csv for easy import into R dataframes
			scores[0][0] = "id"
			scores[0][1] = "scores"
		}
		if i != 0 {
			score := f.AnomalyScores[i-1]
			scores[i][0] = fmt.Sprintf("%d", i-1)
			scores[i][1] = fmt.Sprintf("%f", score)
		}

	}
	return scores
}

func main() {
	images, _ := loadMNISTData()

	// input parameters
	treesNumber := 1000
	subsampleSize := 256
	outliersRatio := 0.01 //v2 adjust this to actually create boolean flags
	routinesNumber := 10

	//model initialization
	forest := iforest.NewForest(treesNumber, subsampleSize, outliersRatio)

	//training
	forest.Train(images)

	//test
	forest.TestParallel(images, routinesNumber)

	//threshold := forest.AnomalyBound
	//anomalyScores := forest.AnomalyScores
	//labelsTest := forest.Labels

	//fmt.Printf("%5.2f ", threshold)
	//fmt.Println(len(forest.AnomalyScores))

	scores := makeScoreSlice(forest)
	/*remember close to 0 = inlier
	close to 1 = outlier
	if everything is around .5, then there is no distinct anomaly
	*/
	file, err := os.Create("./results/go_data.csv")
	if err != nil {
		fmt.Printf("error creating csv %v", err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(scores)

	//fmt.Println(forest.Labels)

}
