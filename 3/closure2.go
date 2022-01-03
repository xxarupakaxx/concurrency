package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	for _, v := range []string{"hello","welocome","おはよう"} {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			fmt.Println(v)
		}(v)
	}
	wg.Wait()
}