package main

import (
	"fmt"
	"time"
)

// --------- Function: Array passed by VALUE (copy) ----------
func modifyByValue(arr [3]int) {
	arr[0] = 100
	fmt.Println("Inside modifyByValue:", arr)
}

// --------- Function: Array passed by POINTER (no copy) ----------
func modifyByPointer(arr *[3]int) {
	arr[0] = 200
	fmt.Println("Inside modifyByPointer:", *arr)
}

// --------- Manual Traversal ----------
func manualTraversal(arr [3]int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Manual traversal sum:", sum)
}

// --------- Copy Cost Demo ----------
func arrayCopyBenchmark() {
	var big [100000]int

	start := time.Now()
	copyArray := big // FULL COPY happens here
	elapsed := time.Since(start)

	fmt.Println("Array copy time:", elapsed)
	_ = copyArray
}

func main() {

	// 1. Fixed-size array
	arr := [3]int{10, 20, 30}
	fmt.Println("Original array:", arr)

	// 2. Indexing
	fmt.Println("First element:", arr[0])

	// 3. Iteration using range
	fmt.Println("Iterating array:")
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// 4. Pass array to function (value copy)
	modifyByValue(arr)
	fmt.Println("After modifyByValue:", arr)

	// 5. Pass array using pointer
	modifyByPointer(&arr)
	fmt.Println("After modifyByPointer:", arr)

	// 6. Manual traversal
	manualTraversal(arr)

	// 7. Array copy cost demo
	arrayCopyBenchmark()
}
