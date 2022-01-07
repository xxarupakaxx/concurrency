package main

import (
	"log"
	"os"
	"time"
)

type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:

				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}

	newSteward := func(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
		return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDonwe chan interface{}
				var wardHeatBeat <-chan interface{}
				startWard := func() {
					wardDonwe = make(chan interface{})
					wardHeatBeat = startGoroutine(or(wardDonwe, done), timeout/2)

				}
				startWard()
				pules := time.Tick(pulseInterval)

			monitoreLoop:
				for true {
					timeoutSignal := time.After(timeout)

					for true {
						select {
						case <-pules:
							select {
							case heartbeat <- struct {
							}{}:
							default:

							}
						case <-wardHeatBeat:
							continue monitoreLoop
						case <-timeoutSignal:
							log.Println("steward: ward unhealthy;restarting")
							close(wardDonwe)
							startWard()
							continue monitoreLoop
						case <-done:
							return
						}

					}
				}
			}()

			return heartbeat
		}
	}
		log.SetOutput(os.Stdout)
		log.SetFlags(log.Ltime| log.LUTC)

		doWork := func(done <-chan interface{},_ time.Duration) <-chan interface{}{
			log.Println("ward: Hello I am irresponsible")
			go func() {
				<-done
				log.Println("ward: I am halting")
			}()

			return nil
		}

		doWorkWithSteward := newSteward(4*time.Second,doWork)

		done :=make(chan interface{})

		time.AfterFunc(9*time.Second, func() {
			log.Println("main: halting steward and ward")
			close(done)
		})

	for range doWorkWithSteward(done,4*time.Second){

	}
	log.Println("done")

}
