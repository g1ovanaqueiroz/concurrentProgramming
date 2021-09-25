package main

import (
	"fmt"
)

func main() {
	n := 42
	result := fib(n)
	fmt.Printf("fib(%d = %d\n", n, result)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
