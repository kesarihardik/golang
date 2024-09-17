package main

import "fmt"

func main() {
	//buffered channels are blocked if the buffer is full.
	c := make(chan int, 1) // causes deadlock
	c <- 1
	c <- 2 //blocked forever since buffer is full
	fmt.Println(c)
	fmt.Println(c)
}
