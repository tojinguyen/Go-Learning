package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
