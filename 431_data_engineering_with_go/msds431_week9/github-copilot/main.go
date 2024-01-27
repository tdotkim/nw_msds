package main

import (
	"fmt"
	"math"
)

type Dataset struct {
	X []float64
	Y []float64
}

func main() {
	// Hardcode the Anscombe Quartet dataset
	datasets := []Dataset{
		{
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		},
		{
			X: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			Y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{
			X: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			Y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
		},
	}

	// Calculate the regression for each dataset
	for i, dataset := range datasets {
		slope, intercept, rSquared := calculateRegression(dataset.X, dataset.Y)

		// Write the results to the terminal
		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("  Slope: %.2f\n", slope)
		fmt.Printf("  Intercept: %.2f\n", intercept)
		fmt.Printf("  R-squared: %.2f\n", rSquared)
	}
}

func calculateRegression(x []float64, y []float64) (slope float64, intercept float64, rSquared float64) {
	// Calculate the mean of x and y
	xMean := mean(x)
	yMean := mean(y)

	// Calculate the slope and intercept of the regression
	var numerator float64
	var denominator float64
	for i := range x {
		numerator += (x[i] - xMean) * (y[i] - yMean)
		denominator += math.Pow(x[i]-xMean, 2)
	}
	slope = numerator / denominator
	intercept = yMean - slope*xMean

	// Calculate the r-squared value of the regression
	var ssr float64
	var sst float64
	for i := range x {
		yHat := slope*x[i] + intercept
		ssr += math.Pow(yHat-yMean, 2)
		sst += math.Pow(y[i]-yMean, 2)
	}
	rSquared = ssr / sst

	return slope, intercept, rSquared
}

func mean(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}
