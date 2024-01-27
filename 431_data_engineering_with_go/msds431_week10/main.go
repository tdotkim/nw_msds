package main

import (
	"fmt"

	randomforest "github.com/malaschitz/randomForest"

	"github.com/petar/GoMNIST"
)

func loadMNISTData() (images [][]float64, labels []int, testimages [][]float64, testlabels []int) {

	train, test, err := GoMNIST.Load("./data")
	if err != nil {
		fmt.Printf("error loading mnist data %v", err)
	}
	images = make([][]float64, len(train.Images))
	labels = make([]int, len(train.Images))
	testimages = make([][]float64, len(test.Images))
	testlabels = make([]int, len(test.Images))

	for i := 0; i < len(train.Images); i++ {
		images[i] = make([]float64, len(train.Images[0]))
		for j := range train.Images[0] {
			images[i][j] = float64(train.Images[i][j])
			labels[i] = int(train.Labels[i])

		}
	}

	for i := 0; i < len(test.Images); i++ {
		testimages[i] = make([]float64, len(test.Images[0]))
		for j := range test.Images[0] {
			testimages[i][j] = float64(test.Images[i][j])
			testlabels[i] = int(test.Labels[i])

		}
	}
	/* fmt.Println("---Length of Data---")
	fmt.Println(len(images))
	fmt.Println(len(labels)) */
	return images, labels, testimages, testlabels

}

func main() {
	images, labels, testimages, testlabels := loadMNISTData()
	trainimages := images[0:50000][:]
	validationimages := images[50001:][:]
	trainlabels := labels[0:50000]
	validationlabels := labels[50001:]

	rf := randomforest.Forest{
		LeafSize: 2000,
		MaxDepth: 2000, // Deep maximum depth
	}

	rf.Data = randomforest.ForestData{X: trainimages, Class: trainlabels}

	//training
	rf.Train(20)
	p := 0

	//validation
	for i := 0; i < len(validationimages); i++ {
		vote := rf.Vote(validationimages[i])
		bestV := 0.0
		bestI := -1
		for j, v := range vote {
			if v > bestV {
				bestV = v
				bestI = j
			}
		}
		if int(validationlabels[i]) == bestI {
			p++
		}
	}
	fmt.Printf("Training Accuracy: %5.1f%%\n", 100.0*float64(p)/float64(len(testimages)))

	//testing
	predlabels := make([]int, 10000)
	p_correct := 0
	p_incorrect := 0
	for i := 0; i < len(testimages); i++ {
		vote := rf.Vote(testimages[i])
		best := 0.0
		for j, v := range vote {
			if v > best {
				best = v
				predlabels[i] = j
			}
		}
		if int(testlabels[i]) == predlabels[i] {
			p_correct++
		} else {
			p_incorrect++
		}
	}

	fmt.Printf("Test Accuracy: %5.1f%%\n", 100.0*float64(p_correct)/float64(len(testimages)))

}
