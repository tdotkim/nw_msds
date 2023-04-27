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
	count := 1000000
	newfloats := make([]float32, count)
	newints := make([]int, count)
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := randInt(0, 10000000, count)
	fmt.Println(ints[0:5])
	start := time.Now()
	for i, v := range ints {
		newints[i] = (multiply(add(multiply(v, 2), 1), 2))
	}
	fmt.Println(newints[0:5])
	duration := time.Since(start)
	fmt.Println("Int time:", duration)

	float_multiply := func(value, multiplier float32) float32 {
		return value * multiplier
	}

	float_add := func(value, additive float32) float32 {
		return value + additive
	}

	floats := randFloats(0, 10000000, count)
	fmt.Println(floats[0:5])
	start = time.Now()
	for i, v := range floats {
		newfloats[i] = (float_multiply(float_add(float_multiply(v, 2), 1), 2))
	}
	duration = time.Since(start)
	fmt.Println(newfloats[0:5])
	fmt.Println("Float time: ", duration)
}
