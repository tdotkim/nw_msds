package main

import (
	"math"
	"testing"
)

func TestCalculateRegression(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 3, 5, 7, 9}

	slope, intercept, rSquared := calculateRegression(x, y)

	expectedSlope := 2.0
	expectedIntercept := -1.0
	expectedRSquared := 1.0

	if math.Abs(slope-expectedSlope) > 0.0001 {
		t.Errorf("Expected slope to be %f, but got %f", expectedSlope, slope)
	}

	if math.Abs(intercept-expectedIntercept) > 0.0001 {
		t.Errorf("Expected intercept to be %f, but got %f", expectedIntercept, intercept)
	}

	if math.Abs(rSquared-expectedRSquared) > 0.0001 {
		t.Errorf("Expected rSquared to be %f, but got %f", expectedRSquared, rSquared)
	}
}

func TestMean(t *testing.T) {
	values := []float64{1, 2, 3, 4, 5}

	expectedMean := 3.0

	if mean(values) != expectedMean {
		t.Errorf("Expected mean to be %f, but got %f", expectedMean, mean(values))
	}
}
