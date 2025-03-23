package main

import "fmt"

// func modifyValue(i int) {  //copy of parameter is modified here
// 	i = 34
// }

func modifyOriginal(i *int) { //a copy of pointer is received with same memory address
	fmt.Printf("\nPointer to i address: %v", &i)
	fmt.Printf("\nReceived value of pointer: %v", i)
	*i = 34
}

//go doesn't allow pointer arithmetic.
func main() {
	var ptr2 *int
	fmt.Printf("\n zero value of pointer: %v", ptr2)
	// fmt.Print(*ptr2)              //runtime error: nil pointer dereferencing.

	var ptr *int = new(int) //new first allocates memory based on type and then returns a pointer to that location.
	fmt.Printf("\nvalue of ptr : %v", *ptr)

	//go supports pass by value. use pointer to modify original field.
	var i int
	// modifyValue(i)
	// fmt.Printf("\nvalue of i: %v", i)
	p := &i
	fmt.Printf("\nPointer to i address: %v", &p)
	fmt.Printf("\n value of i: %v, location of i: %v", i, p)
	modifyOriginal(p)
	fmt.Printf("\nvalue of i: %v", i)
}
