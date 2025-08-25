package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Producer gửi:", i)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Consumer nhận:", v)
	}

	fmt.Println("Thoát vòng lặp")
}
