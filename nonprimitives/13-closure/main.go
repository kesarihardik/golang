package main

import (
	"fmt"
)

// closure - function that capture variables from its surrounding scope even after scope is done executing.
// useful to maintain state(counter), memoization, concurrency etc.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x //sum is retained by inner function even after outer adder() is executed.
		return sum
	}
}

func main() {
	pos := adder()
	fmt.Printf("%T\n", pos)
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
