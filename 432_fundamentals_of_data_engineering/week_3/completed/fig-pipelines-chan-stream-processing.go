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
	floats_slice := randFloats(0, 10000000, count)
	int_slice := randInt(0, 10000000, count)
	fmt.Println("Int Slice size: ", len(int_slice))
	fmt.Println("Float Slice size: ", len(floats_slice))
	fmt.Println("\n")

	intStream := generator(done, int_slice)
	floatStream := floatGenerator(done, floats_slice)

	intstart := time.Now()
	intpipe := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	fmt.Println(int_slice[0:5])
	intduration := time.Since(intstart)
	i := 0
	for v := range intpipe {
		newints[i] = v
		i++
		//fmt.Println(i)
	}
	fmt.Println(newints[0:5])
	fmt.Println("Int pipeline time: ", intduration)
	fmt.Println("\n")

	fmt.Println(floats_slice[0:5])
	floatstart := time.Now()
	floatpipe := floatmultiply(done, floatadd(done, floatmultiply(done, floatStream, 2), 1), 2)
	floatduration := time.Since(floatstart)
	j := 0
	for v := range floatpipe {
		newfloats[j] = v
		j++
		//fmt.Println(j)
	}
	fmt.Println(newfloats[0:5])
	fmt.Println("Float pipeline time: ", floatduration)

	//for v := range pipeline {
	//fmt.Println(v)
	//}
}
