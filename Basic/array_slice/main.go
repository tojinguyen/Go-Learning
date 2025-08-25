package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	batches := BatchSplit(arr, 3)
	for i, batch := range batches {
		fmt.Printf("Batch %d: %v with length %d and capacity %d\n", i+1, batch, len(batch), cap(batch))
	}
}

func BatchSplit[T any](arr []T, size int) [][]T {
	var batches [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		batches = append(batches, arr[i:end])
	}
	return batches
}
