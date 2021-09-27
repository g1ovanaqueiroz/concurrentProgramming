package main

import (
	"fmt"
	"math/rand"
)

func producer(ch chan int) {
	rand.Seed(42)
	for i := 0; i < 100; i++ {
		v := rand.Intn(10)
		ch <- v
	}
	close(ch)
}

func consumer(ch chan int, join_ch chan int) {
	for {
		x, ok := <-ch
		if !ok {
			join_ch <- 1
			break
		}
		if x%2 == 0 {
			fmt.Printf("Even number (%d)\n", x)
		}
	}
}

func main() {
	items := make(chan int)
	join := make(chan int)

	go producer(items)
	go consumer(items, join)

	<-join
}
