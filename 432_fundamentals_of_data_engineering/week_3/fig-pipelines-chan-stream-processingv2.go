package main

import (
	"fmt"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Println(msg, "Runtime : ", time.Since(start), "")
}

func generator(done <-chan interface{}, integers ...int) <-chan int {
	defer duration(track("Int Generator Time: "))
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

func floatGenerator(done <-chan interface{}, floats ...float32) <-chan float32 {
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

func main() {

	floatGenerator := func(done <-chan interface{}, floats ...float32) <-chan float32 {
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

	intStream := generator(done, 1, 2, 3, 4)
	floatStream := floatGenerator(done, 1, 2, 3, 4)

	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	floatPipeline := floatmultiply(done, floatadd(done, floatmultiply(done, floatStream, 2), 1), 2)

	//for v := range pipeline {
	//fmt.Println(v)
	//}
}
