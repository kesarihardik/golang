package main

import "fmt"

var a int16 = 343

const (
	x1 = 12
	y1 = 23
)

// b:=23               //short assignment works only inside function and with non-constants
func main() {
	fmt.Println(x1, y1)

	var a2 bool = true //var [name] type
	b := false         //short assignment. valid only inside function

	// multiple variables can be initialized with diff types
	x, y := 1, true
	fmt.Println(a, a2, b, x, y)

	//short assignment not available for const
	const n = 123
	fmt.Printf("%v %T", n, n)
}
