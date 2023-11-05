package main

import (
	"reflect"
	"testing"
)

func Test_loadMNISTData(t *testing.T) {
	tests := []struct {
		name       string
		wantImages int
		wantLabels int
	}{
		{"mustbe60k", 60000, 60000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImages, gotLabels := loadMNISTData()
			if !reflect.DeepEqual(len(gotImages), tt.wantImages) {
				t.Errorf("loadMNISTData() gotImages = %v, want %v", len(gotImages), tt.wantImages)
			}
			if !reflect.DeepEqual(len(gotLabels), tt.wantLabels) {
				t.Errorf("loadMNISTData() gotLabels = %v, want %v", len(gotLabels), tt.wantLabels)
			}
		})
	}
}

// quick and easy test of converting map to slice
// actual file uses f.AnomalyScores which in and of itself is a map
// doesn't really test to see if the actual version receives a forest
func Test_makeScoreSliceTestable(t *testing.T) {
	tests := []struct {
		name       string
		args       map[int]float64
		wantScores [][]string
	}{
		{"smalltest1", map[int]float64{0: 1.1, 1: 2.2, 2: 3.3}, [][]string{{"id", "scores"}, {"0", "1.100000"}, {"1", "2.200000"}, {"2", "3.300000"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScores := makeScoreSliceTestable(tt.args); !reflect.DeepEqual(gotScores, tt.wantScores) {
				t.Errorf("makeScoreSlice() = %v, want %v", gotScores, tt.wantScores)
			}
		})
	}
}
