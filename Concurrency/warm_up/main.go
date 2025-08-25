package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("A: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for ch := 'a'; ch <= 'j'; ch++ {
			fmt.Printf("B: %c\n", ch)
			time.Sleep(350 * time.Millisecond)
		}
	}()

	wg.Wait()
	print("All goroutines complete.")
}
