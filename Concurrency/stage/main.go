package main

import (
	"fmt"
	"sync"
)

func main() {
	stg1 := make(chan int, 10)
	stg2 := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range stg1 {
			stg2 <- val + 1
		}
		close(stg2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range stg2 {
			fmt.Println("Final value:", val)
		}
	}()

	for i := 0; i < 100; i++ {
		stg1 <- i * 2
	}

	close(stg1)

	wg.Wait()
	println("All stages completed")
}
