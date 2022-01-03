package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup


	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("first gorutine sleeping ...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping ")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Print(":ALL goroutines complete\n")


}
