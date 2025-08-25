package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func HandlePayment(ctx context.Context, amount int) error {
	processTime := time.Duration(rand.Intn(4)+1) * time.Second
	fmt.Printf("Processing payment %d... (will take %v)\n", amount, processTime)

	select {
	case <-time.After(processTime):
		fmt.Printf("Payment %d done!\n", amount)
		return nil
	case <-ctx.Done():
		fmt.Printf("Payment %d canceled: %v\n", amount, ctx.Err())
		return ctx.Err()
	}
}
