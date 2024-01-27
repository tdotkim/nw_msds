package main

import (
	"reflect"
	"testing"

	"github.com/montanaflynn/stats"
)

func Test_roundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{val: 1.23456789, precision: 2}, 1.23},
		{"test2", args{val: 1.23456789, precision: 3}, 1.235},
		{"test3", args{val: 1.23456789, precision: 4}, 1.2346},
		{"test4", args{val: 1.23456789, precision: 5}, 1.23457},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundFloat(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("roundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sums(t *testing.T) {
	type args struct {
		s stats.Series
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		want1   int
		wantErr bool
	}{
		{"test1", args{s: quartet1}, []float64{99, 82.51000000000002, 1001, 797.6, 660.1726999999998}, 11, false},
		{"test2", args{s: quartet2}, []float64{99, 82.51, 1001, 797.59, 660.1763}, 11, false},
		{"test3", args{s: quartet3}, []float64{99, 82.50000000000001, 1001, 797.4699999999999, 659.9762000000001}, 11, false},
		{"test4", args{s: quartet4}, []float64{99, 82.50999999999999, 1001, 797.58, 660.1324999999998}, 11, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := sums(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("sums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sums() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("sums() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_slope_intercept(t *testing.T) {
	type args struct {
		sum []float64
		i   int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		want1   float64
		wantErr bool
	}{
		{"test1", args{sum: []float64{99, 82.51000000000002, 1001, 797.6, 660.1726999999998}, i: 11}, 0.5001, 3.0001, false},
		{"test2", args{sum: []float64{99, 82.51, 1001, 797.59, 660.1763}, i: 11}, 0.5, 3.0009, false},
		{"test3", args{sum: []float64{99, 82.50000000000001, 1001, 797.4699999999999, 659.9762000000001}, i: 11}, 0.4997, 3.0025, false},
		{"test4", args{sum: []float64{99, 82.50999999999999, 1001, 797.58, 660.1324999999998}, i: 11}, 0.4999, 3.0017, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := slope_intercept(tt.args.sum, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("slope_intercept() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("slope_intercept() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("slope_intercept() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_get_rsquared(t *testing.T) {

	q1rs, _ := stats.LinearRegression(quartet1)
	q2rs, _ := stats.LinearRegression(quartet2)
	q3rs, _ := stats.LinearRegression(quartet3)
	q4rs, _ := stats.LinearRegression(quartet4)

	type args struct {
		regressions stats.Series
		quartet     stats.Series
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"test1", args{regressions: q1rs, quartet: quartet1}, .6665, false},
		{"test2", args{regressions: q2rs, quartet: quartet2}, .6662, false},
		{"test3", args{regressions: q3rs, quartet: quartet3}, .6663, false},
		{"test4", args{regressions: q4rs, quartet: quartet4}, .6667, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := get_rsquared(tt.args.regressions, tt.args.quartet)
			if (err != nil) != tt.wantErr {
				t.Errorf("get_rsquared() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("get_rsquared() = %v, want %v", got, tt.want)
			}
		})
	}
}
