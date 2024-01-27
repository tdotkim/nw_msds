package main

import (
	"io/ioutil"
	"math"
	"os"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		x, y     []float64
		expected []float64
	}{
		{[]float64{1, 2, 3}, []float64{2, 3, 4}, []float64{1, 1}},
		{[]float64{1, 2, 3, 4}, []float64{2, 3, 4, 5}, []float64{1, 1}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		data := mat.NewDense(len(test.x), 2, nil)
		for i := range test.x {
			data.Set(i, 0, 1)
			data.Set(i, 1, test.x[i])
		}

		var result mat.Dense
		err := linearRegression(&result, data, test.y)
		if err != nil {
			t.Errorf("linearRegression(%v, %v) failed: %v", test.x, test.y, err)
			continue
		}

		// Check the coefficients
		coefficients := result.RawMatrix().Data
		for i := range test.expected {
			if math.Abs(coefficients[i]-test.expected[i]) > 1e-6 {
				t.Errorf("Test failed for input (%v, %v). Expected coefficients %v, got %v", test.x, test.y, test.expected, coefficients)
				break
			}
		}
	}
}

func TestMainFunction(t *testing.T) {
	// Capture the standard output for testing
	expectedOutput := "Slope: 0.5001\nIntercept: 3.0001\nR-squared: 0.6665\n"

	old := main
	var output string
	main = func() { output = captureOutput(old) }
	defer func() { main = old }()

	main()

	if output != expectedOutput {
		t.Errorf("Unexpected output. Expected:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// captureOutput captures the standard output of a function
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	return string(out)
}
