package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-
// Channels are thread safe
//By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	foo := []int{12, 123, 32, 1, -13, 1}
	c := make(chan int)
	go sum(foo[:len(foo)/2], c)
	go sum(foo[len(foo)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
