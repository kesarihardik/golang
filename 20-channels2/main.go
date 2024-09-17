package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //close indicates to receivers to signal receiver that no more channels are there. not mandatory though.
}

func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
