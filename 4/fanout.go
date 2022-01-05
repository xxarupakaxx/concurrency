package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():

				}
			}
		}()
		return valueStream
	}
	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)

		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):

				}
			}
		}()
		return intStream
	}
	rand := func() interface{} {
		return rand.Intn(50000000)
	}
	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:

				}
			}
		}()

		return takeStream
	}

	primeFinder := func(done <-chan interface{},valueStream <-chan int) <-chan interface{}{
		primeStream := make(chan interface{})

		go func() {
			defer close(primeStream)
			for i := range valueStream {
				select {
				case <-done:
					return
				case primeStream<- i:

				}
			}
		}()
		return primeStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Prime:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)

	}
	fmt.Printf("Search took:%v ", time.Since(start))
}
