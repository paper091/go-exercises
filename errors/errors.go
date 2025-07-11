package main

import (
	"errors"
	"fmt"
)

func main() {
	const a = "error message"
	b := errors.New(a)
	c := errors.New(a)
	if b == c {
		fmt.Println("This won;t print")
	} else {
		fmt.Println("Hello")
	}
}
