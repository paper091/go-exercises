package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	// initiated generator with 0!
	factRes, factorialGenerator := factorial()
	for i := 0; i < n; i++ {
		factNum := factorialGenerator(&factRes)
		fmt.Printf("%v!: %v\n", factNum, factRes)
	}
}

func factorial() (int, func(fact *int) int) {
	var n int = 0
	var fact int = 1
	return fact, func(fact *int) int {
		n++
		*fact = *fact * n
		return n
	}
}

/*
NOTE: pre/post increment/decrement does NOT exist in GO.
the default decrement in GO cannot be chained with return statements or other expressions
*/
