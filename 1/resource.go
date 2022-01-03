package main

import (
	"fmt"
	"sync"
	"time"
)


func main()  {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greadyWorker :=func ()  {
		defer wg.Done()

		var count int
		for begin := time.Now();time.Since(begin) <= runtime;{
			sharedLock.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("cready worker was able to execute %v work loops\n",count)
	}

	politeWorker := func ()  {
		defer wg.Done()

		var count int
		for begin := time.Now();time.Since(begin) <= runtime;{
			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()
			
			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("poloite worker was able to execute %v worke loops \n",count)
	}

	wg.Add(2)
	go greadyWorker()
	go politeWorker()

	wg.Wait()
}