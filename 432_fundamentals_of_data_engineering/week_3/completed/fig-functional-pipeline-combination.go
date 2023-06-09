package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do_mult_int(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func do_add_int(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

func do_mult_float(values []float32, multiplier float32) []float32 {
	multipliedValues := make([]float32, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func do_add_float(values []float32, additive float32) []float32 {
	addedValues := make([]float32, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

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
	rand.Seed(time.Now().UnixNano())
	count := 1000000
	newfloats := make([]float32, count)
	newints := make([]int, count)

	//chose float 32 to save on some memory
	//big = 1M
	//med = 100K
	//small = 10k
	//doing the floats
	floats_slice := randFloats(0, 10000000, count)

	fmt.Println("Floats")
	fmt.Println(floats_slice[0:5])
	start := time.Now()
	for i, v := range do_add_float(do_mult_float(floats_slice, 2), 1) {
		newfloats[i] = v
	}
	fmt.Println(newfloats[0:5])
	duration := time.Since(start)
	fmt.Println("Float time: ", duration)
	fmt.Println("----------------------")

	//doing the ints
	int_slice := randInt(0, 10000000, count)

	fmt.Println("----------------------")
	fmt.Println("Int")
	fmt.Println(int_slice[0:5])
	start = time.Now()
	for i, v := range do_add_int(do_mult_int(int_slice, 2), 1) {
		newints[i] = v
	}
	fmt.Println(newints[0:5])
	duration = time.Since(start)

	fmt.Println("Int time: ", duration)
	fmt.Println("----------------------")

}
