package main

import (
	"fmt"
)

func main() {

	// declaring variable
	var n int
	// taking input
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	// printing patterns
	fmt.Println("Pattern 1:")
	rightTriangle(n)

	fmt.Println("\nPattern 2:")
	rightNumericTriangle(n)

	fmt.Print("\nPattern 3:\n")
	equilateralTriangle(n)

}

// functions in Go are first class citizens
func rightTriangle(n int) {
	arr := "*"
	for i := 1; i < n; i++ {
		arr += arr
	}
	for i := 1; i <= n; i++ {
		fmt.Println(arr[0:i])
	}
}

// Go Genercis
// [T any] => type parameter
func fill[T any](slice []T, value T) {
	for index := range slice {
		slice[index] = value
	}
}

func rightNumericTriangle(n int) {
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, i+1)
		fill(nums[i], (i+1))
		fmt.Println(nums[i])
	}
}

func equilateralTriangle(n int) {
	// allocates memory
	nums := make([][]string, n)
	for i := 0; i < n; i++ {
		for j := n - i - 1; j >= 0; j-- {
			nums[i] = append(nums[i], " ")
		}
		for j := 0; j < 2*i+1; j++ {
			nums[i] = append(nums[i], "*")
		}
	}

	for _, val := range nums {
		fmt.Println(val)
	}
}
