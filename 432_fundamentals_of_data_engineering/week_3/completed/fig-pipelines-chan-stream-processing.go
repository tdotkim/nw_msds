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
	generator := func(done <-chan interface{}, integers []int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	floatGenerator := func(done <-chan interface{}, floats []float32) <-chan float32 {
		floatStream := make(chan float32)
		go func() {
			defer close(floatStream)
			for _, i := range floats {
				select {
				case <-done:
					return
				case floatStream <- i:
				}
			}
		}()
		return floatStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	floatmultiply := func(
		done <-chan interface{},
		floatStream <-chan float32,
		multiplier float32,
	) <-chan float32 {
		multipliedStream := make(chan float32)
		go func() {
			defer close(multipliedStream)
			for i := range floatStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	floatadd := func(
		done <-chan interface{},
		floatStream <-chan float32,
		additive float32,
	) <-chan float32 {
		addedStream := make(chan float32)
		go func() {
			defer close(addedStream)
			for i := range floatStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	//1000000
	//100000
	//10000
	floats_slice := randFloats(0, 10000000, 10000)
	int_slice := randInt(0, 10000000, 10000)
	fmt.Println("Int Slice size: ", len(int_slice))
	fmt.Println("Float Slice size: ", len(floats_slice))
	fmt.Println("\n")

	intStream := generator(done, int_slice)
	floatStream := floatGenerator(done, floats_slice)

	intstart := time.Now()
	multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	intduration := time.Since(intstart)
	fmt.Println("Int pipeline time: ", intduration)
	fmt.Println("\n")

	floatstart := time.Now()
	floatmultiply(done, floatadd(done, floatmultiply(done, floatStream, 2), 1), 2)
	floatduration := time.Since(floatstart)
	fmt.Println("Float pipeline time: ", floatduration)

	//for v := range pipeline {
	//fmt.Println(v)
	//}
}
