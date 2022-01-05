package main

import (
	"fmt"
	"time"
)

func main() {
	
	data := make([]int,4)

	loopData := func (handleData chan<- int)  {
		defer close(handleData)
		for i := range data{
			handleData <- data[i]
			time.Sleep(1*time.Second)
		}
		
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Print(num)
	}
}