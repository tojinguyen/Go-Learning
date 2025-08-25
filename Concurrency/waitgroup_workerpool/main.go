package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Microsecond * 200)
	}
}

func main() {
	var wg sync.WaitGroup
	var numWorkers int = 3
	var numJobs int = 10

	jobs := make(chan int, numJobs)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()

	fmt.Println("All jobs processed")
}
