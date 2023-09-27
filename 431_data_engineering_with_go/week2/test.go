package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func Testsums(t *testing.T)

func sums(s stats.Series) ([]float64, int, error) {
	if len(s) == 0 {
		return []float64{}, 0, stats.EmptyInputErr
	}

	// sum holder
	sum := make([]float64, 5)

	// cumulative sums
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X          // sum of X
		sum[1] += s[i].Y          // sum of Y
		sum[2] += s[i].X * s[i].X // sum of X squared
		sum[3] += s[i].X * s[i].Y // sum of X * Y
		sum[4] += s[i].Y * s[i].Y // sum of Y squared
	}

	return sum, i, nil
}

func slope_intercept(sum []float64, i int) (float64, float64, error) {
	if len(sum) == 0 {
		return 0, 0, stats.EmptyInputErr
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	gradient_rounded := roundFloat(gradient, 4)

	intercept := (sum[1] / f) - (gradient * sum[0] / f)
	intercept_rounded := roundFloat(intercept, 4)

	return gradient_rounded, intercept_rounded, nil
}

func get_rsquared(regressions stats.Series, quartet stats.Series) (float64, error) {
	if len(regressions) == 0 {
		return 0, stats.EmptyInputErr
	}
	if len(quartet) == 0 {
		return 0, stats.EmptyInputErr
	}

	var ytotal float64 //get total. could probably do this elsewhere
	i := 0
	for ; i < len(regressions); i++ {
		ytotal += quartet[i].Y
	}

	ymean := ytotal / float64(len(regressions))

	var yhat float64
	var ymeandiff float64
	j := 0

	for ; j < len(regressions); j++ {
		yhat += math.Pow((quartet[j].Y - regressions[j].Y), 2) //get y-yhat
		ymeandiff += math.Pow((quartet[j].Y - ymean), 2)       //get y-ymean
	}

	rsquared := 1 - (yhat / ymeandiff)
	rsquared_rounded := roundFloat(rsquared, 4)
	return rsquared_rounded, nil
}

func main() {

	quartet1 := []stats.Coordinate{
		{X: 10.0, Y: 8.04},
		{X: 8.0, Y: 6.95},
		{X: 13.0, Y: 7.58},
		{X: 9.0, Y: 8.81},
		{X: 11.0, Y: 8.33},
		{X: 14.0, Y: 9.96},
		{X: 6.0, Y: 7.24},
		{X: 4.0, Y: 4.26},
		{X: 12.0, Y: 10.84},
		{X: 7.0, Y: 4.82},
		{X: 5.0, Y: 5.68},
	}

	quartet2 := []stats.Coordinate{
		{X: 10.0, Y: 9.14},
		{X: 8.0, Y: 8.14},
		{X: 13.0, Y: 8.74},
		{X: 9.0, Y: 8.77},
		{X: 11.0, Y: 9.26},
		{X: 14.0, Y: 8.10},
		{X: 6.0, Y: 6.13},
		{X: 4.0, Y: 3.10},
		{X: 12.0, Y: 9.13},
		{X: 7.0, Y: 7.26},
		{X: 5.0, Y: 4.74},
	}

	quartet3 := []stats.Coordinate{
		{X: 10.0, Y: 7.46},
		{X: 8.0, Y: 6.77},
		{X: 13.0, Y: 12.74},
		{X: 9.0, Y: 7.11},
		{X: 11.0, Y: 7.81},
		{X: 14.0, Y: 8.84},
		{X: 6.0, Y: 6.08},
		{X: 4.0, Y: 5.39},
		{X: 12.0, Y: 8.15},
		{X: 7.0, Y: 6.42},
		{X: 5.0, Y: 5.73},
	}

	quartet4 := []stats.Coordinate{
		{X: 8.0, Y: 6.58},
		{X: 8.0, Y: 5.76},
		{X: 8.0, Y: 7.71},
		{X: 8.0, Y: 8.84},
		{X: 8.0, Y: 8.47},
		{X: 8.0, Y: 7.04},
		{X: 8.0, Y: 5.25},
		{X: 19.0, Y: 12.50},
		{X: 8.0, Y: 5.56},
		{X: 8.0, Y: 7.91},
		{X: 8.0, Y: 6.89},
	}

	var q1model model
	var q2model model
	var q3model model
	var q4model model

	q1model.regressionVals, _ = stats.LinearRegression(quartet1)
	q1model.sums, q1model.obs, _ = sums(quartet1)
	q1model.slope, q1model.intercept, _ = slope_intercept(q1model.sums, q1model.obs)
	q1model.rsquared, _ = get_rsquared(q1model.regressionVals, quartet1)
	fmt.Printf("---Quartet 1---\n")
	fmt.Printf("Quartet 1 Slope %v\n", q1model.slope)
	fmt.Printf("Quartet 1 Intercept %v\n", q1model.intercept)
	fmt.Printf("Quartet 1 R-Squared %v\n", q1model.rsquared)
	fmt.Printf("----------------\n")

	q2model.regressionVals, _ = stats.LinearRegression(quartet2)
	q2model.sums, q2model.obs, _ = sums(quartet2)
	q2model.slope, q2model.intercept, _ = slope_intercept(q2model.sums, q2model.obs)
	q2model.rsquared, _ = get_rsquared(q2model.regressionVals, quartet2)
	fmt.Printf("---Quartet 2---\n")
	fmt.Printf("Quartet 2 Slope %v\n", q2model.slope)
	fmt.Printf("Quartet 2 Intercept %v\n", q2model.intercept)
	fmt.Printf("Quartet 2 R-Squared %v\n", q2model.rsquared)
	fmt.Printf("----------------\n")

	q3model.regressionVals, _ = stats.LinearRegression(quartet3)
	q3model.sums, q3model.obs, _ = sums(quartet3)
	q3model.slope, q3model.intercept, _ = slope_intercept(q3model.sums, q3model.obs)
	q3model.rsquared, _ = get_rsquared(q3model.regressionVals, quartet3)
	fmt.Printf("---Quartet 3---\n")
	fmt.Printf("Quartet 3 Slope %v\n", q3model.slope)
	fmt.Printf("Quartet 3 Intercept %v\n", q3model.intercept)
	fmt.Printf("Quartet 3 R-Squared %v\n", q3model.rsquared)
	fmt.Printf("----------------\n")

	q4model.regressionVals, _ = stats.LinearRegression(quartet4)
	q4model.sums, q4model.obs, _ = sums(quartet4)
	q4model.slope, q4model.intercept, _ = slope_intercept(q4model.sums, q4model.obs)
	q4model.rsquared, _ = get_rsquared(q4model.regressionVals, quartet4)
	fmt.Printf("---Quartet 4---\n")
	fmt.Printf("Quartet 4 Slope %v\n", q4model.slope)
	fmt.Printf("Quartet 4 Intercept %v\n", q4model.intercept)
	fmt.Printf("Quartet 4 R-Squared %v\n", q4model.rsquared)
	fmt.Printf("----------------\n")

}
