package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct{
		Clicked *sync.Cond
	}
	button := Button{
		sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func (c *sync.Cond,fn func())  {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegister sync.WaitGroup
	clickRegister.Add(3)
	subscribe(button.Clicked,func ()  {
		fmt.Println("maximizing window")
		clickRegister.Done()
	})
	subscribe(button.Clicked,func() {
		fmt.Println("Dispalyigannoying dialog box")
		clickRegister.Done()
	})

	subscribe(button.Clicked,func ()  {
		fmt.Println("Mouse clisedk")
		clickRegister.Done()
	})

	button.Clicked.Broadcast()
	clickRegister.Wait()
}