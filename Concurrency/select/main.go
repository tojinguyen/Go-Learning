package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numChan = make(chan int)
	var quitChan = make(chan bool)

	go func() {
		for i := 1; i <= 5; i++ {
			numChan <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond)
		}

		quitChan <- true
	}()

	for {
		select {
		case num := <-numChan:
			println("Nhận số:", num)
		case <-quitChan:
			println("Nhận tín hiệu thoát")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
