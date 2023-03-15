package main

import (
	"fmt"
	"math/rand"
)

func generateRandomArray(array []int64, min int64, max int64) {
	for i := 0; i < len(array); i++ {
		array[i] = rand.Int63n(max-min) + min
	}
}

func sum_array(array []int64) int64 {
	var sum int64 = 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

func main() {
	rand.Seed(42)

	var v1 [20]int64
	var v2 [20]int64

	generateRandomArray(v1[:], 0, 100)
	generateRandomArray(v2[:], 0, 100)

	var sum1 = sum_array(v1[:])
	var sum2 = sum_array(v2[:])

	fmt.Println("v1 = ", v1)
	fmt.Println("v2 = ", v2)
	fmt.Println("Suma v1 = ", sum1)
	fmt.Println("Suma v2 = ", sum2)
}
