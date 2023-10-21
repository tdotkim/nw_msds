package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func checkslices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// quartets
var quartet1 = []stats.Coordinate{
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

var quartet2 = []stats.Coordinate{
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

var quartet3 = []stats.Coordinate{
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

var quartet4 = []stats.Coordinate{
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

// expected slices
var sumslice1 = []float64{99, 82.51000000000002, 1001, 797.6, 660.1726999999998}
var sumslice2 = []float64{99, 82.51, 1001, 797.59, 660.1763}
var sumslice3 = []float64{99, 82.50000000000001, 1001, 797.4699999999999, 659.9762000000001}
var sumslice4 = []float64{99, 82.50999999999999, 1001, 797.58, 660.1324999999998}

type sumTest struct {
	arg1     stats.Series
	expected []float64
}

var sumslicetests = []sumTest{
	{arg1: quartet1, expected: sumslice1},
	{arg1: quartet2, expected: sumslice2},
	{arg1: quartet3, expected: sumslice3},
	{arg1: quartet4, expected: sumslice4},
}

func TestSums(t *testing.T) {
	for _, test := range sumslicetests {
		sumresult, sumi, _ := sums(test.arg1)
		result := checkslices(sumresult, test.expected)
		if result == false {
			t.Errorf("\nIncorrect sums\n Got: %v\n Want: %v\n", sumresult, test.expected)
		}

		if sumi != 11 {
			t.Errorf("\nIncorrect obs\n Got: %v\n Want: %v\n", sumi, 11)
		}
	}
}

// testing slope intercept results
type slopeInterceptTest struct {
	arg1              []float64
	expectedSlope     float64
	expectedIntercept float64
}

var SlopeIntTests = []slopeInterceptTest{
	{arg1: sumslice1, expectedSlope: 0.5001, expectedIntercept: 3.0001},
	{arg1: sumslice2, expectedSlope: 0.5, expectedIntercept: 3.0009},
	{arg1: sumslice3, expectedSlope: 0.4997, expectedIntercept: 3.0025},
	{arg1: sumslice4, expectedSlope: 0.4999, expectedIntercept: 3.0017},
}

func TestSlopeIntercept(t *testing.T) {
	for _, test := range SlopeIntTests {
		slope, intercept, _ := slope_intercept(test.arg1, 11)
		if slope != test.expectedSlope && (intercept != test.expectedIntercept) {
			t.Errorf("\nIncorrect slope/int \n Got: %v,%v\n Want: %v, %v\n", slope, intercept, test.expectedSlope, test.expectedIntercept)
		}
	}
}

// testing rsquared results
type rsquareTest struct {
	arg1      stats.Series
	expectedR float64
}

var rSquaredTests = []rsquareTest{
	{arg1: quartet1, expectedR: .6665},
	{arg1: quartet2, expectedR: .6662},
	{arg1: quartet3, expectedR: .6663},
	{arg1: quartet4, expectedR: .6667},
}

func TestRSquared(t *testing.T) {
	for _, test := range rSquaredTests {
		regressions, _ := stats.LinearRegression(test.arg1)
		rsquaredval, _ := get_rsquared(regressions, test.arg1)
		if rsquaredval != test.expectedR {
			t.Errorf("\nIncorrect slope/int \n Got: %v\n Want: %v\n", rsquaredval, test.expectedR)
		}
	}
}
