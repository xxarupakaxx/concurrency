package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sort"
	"sync"
)

func Open() *APIConnection {

	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1),1),
	}
}

type APIConnection struct {
	rateLimiter *rate.Limiter

}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	defer log.Printf("Done.")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannnot ReadFile:%v",err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resplveAdrree :%v",err)
			}
			log.Println("Resolve")
		}()
	}

	wg.Wait()
}

type Ratelimiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...Ratelimiter) *multiLimiter {
	byLimit := func(i,j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters,byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []Ratelimiter
}

func (m *multiLimiter) Wait(ctx context.Context) error {
	for _, limiter := range m.limiters {
		if err := limiter.Wait(ctx); err!=nil{
			return err
		}
	}

	return nil
}



func (m *multiLimiter) Limit() rate.Limit {
	return m.limiters[0].Limit()
}
