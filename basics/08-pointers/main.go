package main

import "fmt"

//go doesn't allow pointer arithmetic.
func main() {
	// var ptr2 *int
	// fmt.Print(*ptr2)              //runtime error: nil pointer dereferencing.

	var ptr *int32 = new(int32) //new first allocates memory based on type and then returns a pointer to that location.
	fmt.Printf("value of ptr : %v", *ptr)

	var i int32
	fmt.Printf("\nvalue of i: %v", i)

	ptr = &i
	*ptr = 10
	fmt.Printf("\nvalue of i: %v ;; type of ptr: %T", i, ptr)
}
