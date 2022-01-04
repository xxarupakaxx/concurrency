package main

import (
	"fmt"
	"math"
	"sync"
	"text/tabwriter"
	"time"
	"os"
)

func main() {
	producer := func (wg *sync.WaitGroup, l sync.Locker)  {
		defer wg.Done()
		for i := 0; i <= 5; i++ {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}

	}
	observer := func (wg *sync.WaitGroup, l sync.Locker)  {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()

	}

	test := func (count int,mutex,rwMutex sync.Locker) time.Duration  {
		var wg sync.WaitGroup
		wg.Add(count+1)
		begintestTime := time.Now()
		go producer(&wg,mutex)

		for i := 0; i < count; i++ {
			go observer(&wg,rwMutex)
		}
		wg.Wait()

		return time.Since(begintestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout,0,1,2,' ',0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw,"Reader \t RWMutext \t Mutex \n")
	for i := 0; i < 29; i++ {
		count := int(math.Pow(2,float64(i)))
		fmt.Fprintf(tw,"%d\t %v \t %v \n",count,test(count,&m,m.RLocker()),test(count,&m,&m))
	}
}