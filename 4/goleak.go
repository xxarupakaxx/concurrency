package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	newRandStream := func (done <- chan interface{}) <- chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exitieed")
			defer close(randStream)
			for{
				select {
				case randStream <- rand.Int():
				case <- done:
					return
				}
			
			}

		}()

		return randStream
	}
	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Print("aaa")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d : %d \n",i,<- randStream)
	}
	close(done)
	time.Sleep(1 *time.Second)
}