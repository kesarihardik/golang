package main

import "fmt"

//closure   - func that references variable outside its body
// used to encapsulate state within functions.
// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		t := x
		x = y
		y += t
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
