package main

import (
	"fmt"
)

func main() {
	safeWrapper(func() {
		riskyFunc(10, 10)
	})
	fmt.Println("execution of main() continues anyways")
}

func riskyFunc(a int, b int) int {
	c := a / b
	fmt.Printf("%v/%v=%v\n", a, b, c)
	return c
}

// wrapper to safely execute any function
func safeWrapper(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
	fn()
}

// TODO
// func safeWrapper[T interface{}](fn func(...T) T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered: ", r)
// 		}
// 	}()
// 	fn()
// }
