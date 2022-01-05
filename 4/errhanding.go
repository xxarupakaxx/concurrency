package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error error
	Response *http.Response
}

func main() {
	checkStatu := func(done <-chan interface{},urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp ,err := http.Get(url)
				result = Result{Error: err ,Response: resp}
				select {
				case <-done:
					return
				case results <- result:

				}
			}

		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	errCount := 0
	urls := []string{"a","https://www.google.com","https://badhost","b","c"}
	for r:= range checkStatu(done, urls...){
		if  r.Error != nil{
			fmt.Println(r.Error)
			errCount++
			continue
		}
		fmt.Printf("Response: %v\n", r.Response.Status)
	}
}


