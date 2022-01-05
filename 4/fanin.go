package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	funIn := func(done <-chan interface{},channels ...<-chan interface{}) <-chan interface{}{
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <-i:

				}
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	rand := func() interface{} {
		return rand.Intn(50000000)
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
	done := make(chan interface{})
	defer close(done)
	randIntStream := toInt(done,repeatFn(done,rand))

	numFinders := runtime.NumCPU()
	fmt.Println(numFinders)
	finders := make([]<- chan interface{},numFinders)
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

	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done,randIntStream)
	}
}
