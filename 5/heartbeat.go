package main

import (
	"fmt"
	"time"
)

func main() {
	dowork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartBeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			defer close(heartBeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartBeat <- struct {
				}{}:
				default:

				}
			}
			sendResult := func(r time.Time) {
				for true {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for true {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)

				}
			}
		}()
		return heartBeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second
	heartbeat, results := dowork(done, timeout/2)
	for true {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results :%v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}
