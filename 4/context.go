package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Println("failed to print greeting:", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewell(ctx); err != nil {
			fmt.Println("failed to print farewell:", err)
		}
	}()
	wg.Wait()
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Println(farewell)
	return nil
}

func genFarewell(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Hour)
	defer cancel()

	switch local, err := locale(ctx); {
	case err != nil:
		return "", err
	case local == "EN":
		return "heelo", nil

	}
	return "", fmt.Errorf("unsupported")
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Println(greeting)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Hour)
	defer cancel()
	switch local, err := locale(ctx); {
	case err != nil:
		return "", err
	case local == "EN":
		return "goddbye", nil

	}
	return "", fmt.Errorf("unsupported")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):

	}
	return "EN", nil
}
