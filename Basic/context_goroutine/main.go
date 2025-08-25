package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		amount := i * 100
		go func(amount int) {
			defer wg.Done()
			_ = HandlePayment(ctx, amount)
		}(amount)
	}

	// Chờ tất cả goroutines xong
	wg.Wait()
	fmt.Println("🎉 All payments processed (or canceled)")
}
