package main

import (
	"fmt"
	"time"
)

func main() {
	n := 42
	go alert()
	result := fib(n)
	fmt.Printf("Fib(%d) is %d\n", n, result)

}

func alert() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("It's not over yet, wait!\n")
	}
}

func fib(x int) int {
	if x < 2 {
		return 2
	}
	return fib(x-1) + fib(x-2)
}
