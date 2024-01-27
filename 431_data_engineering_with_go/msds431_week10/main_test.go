package main

import (
	"testing"
)

func Test_loadMNISTData(t *testing.T) {
	tests := []struct {
		name           string
		wantImages     [][]float64
		wantLabels     []int
		wantTestimages [][]float64
		wantTestlabels []int
	}{
		{"test1", nil, nil, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImages, gotLabels, gotTestimages, gotTestlabels := loadMNISTData()
			if len(gotImages) != 60000 {
				t.Errorf("Incorrect number of images in train")
			}
			if len(gotLabels) != 60000 {
				t.Errorf("Incorrect number of labels in train")
			}
			if len(gotTestimages) != 10000 {
				t.Errorf("Incorrect number of images in test")
			}
			if len(gotTestlabels) != 10000 {
				t.Errorf("Incorrect number of labels in test")
			}
		})
	}
}
