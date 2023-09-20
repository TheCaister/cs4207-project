package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := generateRandomArray(100, 100)
	fmt.Println("Unsorted array: ", arr)
}

func generateRandomArray(size, upperValue int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(upperValue)
	}
	return arr
}
