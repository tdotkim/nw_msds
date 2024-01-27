package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Anscombe Quartet dataset
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Create a matrix for linear regression
	data := mat.NewDense(len(x), 2, nil)
	for i := range x {
		data.Set(i, 0, 1)
		data.Set(i, 1, x[i])
	}

	// Perform linear regression
	var regression mat.Dense
	err := linearRegression(&regression, data, y)
	if err != nil {
		log.Fatal(err)
	}

	// Get the slope, intercept, and R-squared value
	slope := regression.At(1, 0)
	intercept := regression.At(0, 0)

	// Calculate R-squared value
	var residuals mat.VecDense
	residuals.SubVec(mat.NewVecDense(len(y), y), data.MulVec(&regression, mat.NewVecDense(len(x), nil)))
	rSquared := calculateRSquared(&residuals, y)

	// Print the results
	fmt.Printf("Slope: %.4f\n", slope)
	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("R-squared: %.4f\n", rSquared)
}

// linearRegression performs linear regression and returns the coefficients.
func linearRegression(result *mat.Dense, data mat.Matrix, y []float64) error {
	yVec := mat.NewVecDense(len(y), y)
	if err := result.Solve(data, yVec); err != nil {
		return fmt.Errorf("linear regression failed: %v", err)
	}
	return nil
}

// calculateRSquared calculates the R-squared value.
func calculateRSquared(residuals *mat.VecDense, y []float64) float64 {
	// Convert residuals to a slice of float64
	residualsSlice := residuals.RawVector().Data

	// Calculate R-squared value
	yMean := floats.Sum(y) / float64(len(y))
	rSquared := 1.0 - floats.Norm(residualsSlice, 2)/floats.Norm(mat.NewVecDense(len(y), floats.Apply(func(i int, v float64) float64 {
		return v - yMean
	}, y)), 2)
	return rSquared
}
