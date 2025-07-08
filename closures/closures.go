package main

import (
	"fmt"
)

func fibonacci() func() []int {
	sequence := []int{}
	return func() []int {
		if len(sequence) == 0 {
			sequence = append(sequence, 0)
		} else if len(sequence) < 2 {
			sequence = append(sequence, 1)
		} else {
			sequence = append(sequence, (sequence[len(sequence)-1] + sequence[len(sequence)-2]))
		}
		return sequence
	}
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	fibonacciGenerator := fibonacci()
	var fibSequence []int
	for i := 0; i < n; i++ {
		fibSequence = fibonacciGenerator()
	}
	fmt.Println(fibSequence)
}
