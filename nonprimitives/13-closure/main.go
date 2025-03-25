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
	counter := 0

	func() {
		counter += 10 //counter is accessed by anonymous function
	}()
	fmt.Print(counter, "\n")

	pos := adder()
	fmt.Printf("%T\n", pos)
	for i := range 10 {
		fmt.Println(pos(i))
	}
}
