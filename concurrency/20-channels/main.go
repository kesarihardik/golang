package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-
// Channels are thread safe. Values are passed through channels in FIFO order.
//By default, sends and receives block until the other side is ready.
//This allows goroutines to synchronize without explicit locks or condition variables.

func sum(arr []int, c chan int) {
	fmt.Print("\n summing these values: ", arr)
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
	x, y := <-c, <-c //values are read, write in FIFO order in channels.
	fmt.Println("\n", x, y, x+y)
}
