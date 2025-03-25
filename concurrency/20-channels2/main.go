package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for range n {
		c <- x
		x, y = y, x+y
	}
	close(c) //close indicates to receivers to signal receiver that no more channels are there.
	//only senders should close channels since sending on close channels cause panic.

}

func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c { //sender must close c to indicate that it is done executing else a deadlock will occur.
		fmt.Println(i)
	}
}
