package main

import (
	"fmt"
	"sync"
	"time"

)

func main() {
	c := sync.NewCond(&sync.Mutex{})

	quene := make([]interface{}, 0, 10)

	removeFromQunue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		quene = quene[1:]
		fmt.Println("Removed from quuene")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()

		for len(quene) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		quene = append(quene, struct{}{})
		go removeFromQunue(1 * time.Second)
		c.L.Unlock()
	}

}
