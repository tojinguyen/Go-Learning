package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	results := make(chan string, 3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			select {
			case <-ctx.Done():
				return
			case results <- callAPI(i, ctx):
			}
		}(i)
	}

	select {
	case res := <-results:
		println("First result:", res)
	case <-ctx.Done():
		println("Context timed out")
		return
	}
}

func callAPI(id int, ctx context.Context) string {
	delay := time.Duration(rand.Float32()*(3)+1) * time.Second
	println("Calling API with delay:", delay)

	select {
	case <-time.After(delay):
		println("API call finished")
		return fmt.Sprintf("Result from API %d", id)
	case <-ctx.Done():
		println("API call canceled")
		return fmt.Sprintf("API %d canceled", id)
	}
}
