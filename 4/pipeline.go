package main

import "fmt"

func main() {
	generator := func(done <-chan interface{},integers ...int) <-chan int {
		intStream := make(chan int,len(integers))
		go func() {
			defer close(intStream)
			for _, integer := range integers {
				select {
				case <-done:
					return
				case intStream<- integer:

				}

			}
		}()

		return intStream
	}

	multiply := func(done <-chan interface{},intStream <-chan int,multiply int) <- chan int{
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream<- i*multiply:

				}

			}
		}()

		return multipliedStream
	}

	add := func(done <- chan interface{},intStream <-chan int,additive int) <-chan int {
		addedStream :=make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream{
				select {
				case <- done:
					return
				case addedStream<- i+additive:

				}
			}
		}()

		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done,1,2,3,4)
	pipeline := multiply(done,add(done,multiply(done,intStream,2),1),2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
