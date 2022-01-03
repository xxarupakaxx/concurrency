package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	syaHello := func ()  {
		defer wg.Done()
		fmt.Println("hello")
		
	}

	wg.Add(1)
	go syaHello()
	wg.Wait()

	
}