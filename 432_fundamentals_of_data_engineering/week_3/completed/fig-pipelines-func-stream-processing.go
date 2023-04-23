package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randFloats(min, max float32, n int) []float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = min + rand.Float32()*(max-min)
	}
	return res
}

func randInt(min, max int, n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(max-min) + min
	}
	return res
}

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := randInt(0, 10000000, 10000)

	start := time.Now()
	for _, v := range ints {
		(multiply(add(multiply(v, 2), 1), 2))
	}
	duration := time.Since(start)
	fmt.Println("Int 10K time: ", duration)

	float_multiply := func(value, multiplier float32) float32 {
		return value * multiplier
	}

	float_add := func(value, additive float32) float32 {
		return value + additive
	}

	floats := randFloats(0, 10000000, 10000)

	start = time.Now()
	for _, v := range floats {
		(float_multiply(float_add(float_multiply(v, 2), 1), 2))
	}
	duration = time.Since(start)
	fmt.Println("Float time: ", duration)
}
