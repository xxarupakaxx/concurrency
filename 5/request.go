package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	doWork := func(done <-chan interface{},id int,wg *sync.WaitGroup,result chan<- int) {
		started := time.Now()
		defer wg.Done()

		simulatedLoadTime := time.Duration(1+rand.Intn(5))*time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):

		}
		select {
		case <-done:
		case result <- id:

		}

		took := time.Since(started)
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}

		fmt.Println(id,"took ",took)


	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go doWork(done,i,&wg,result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()
	fmt.Println(firstReturned)
}
