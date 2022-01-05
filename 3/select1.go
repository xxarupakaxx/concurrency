package main

import (
	"fmt"
	"time"
)

func main() {
	
	done := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(done)
	}()

	woker :=0
	loop:
	for{
		select {
		case <-done:
			break loop
		default:	
		}

		woker++
		time.Sleep(1*time.Second)
		
	}

	fmt.Printf("Achived %v cycles of work before signalled to sstpo ",woker)
}