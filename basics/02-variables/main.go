package main

import "fmt"

//global variables
var a int16 = 343

const (
	x1 = 12 //untype constants - leads to type flexibility. high precision arithmetic, avoids overflow error,
	y1 = 23
)

// b:=23               //short assignment works only inside function and with non-constants
func main() {
	fmt.Println(x1, y1)

	var a2 bool = true
	b := false //short assignment

	// multiple variables can be initialized with diff types
	// var x,y = 1, true
	x, y := 1, true //type inference - type is inferred from value. in case of numbers type depends on precision.
	fmt.Println(a, a2, b, x, y)

	//short assignment not available for const
	const n = 123
	fmt.Printf("%v %T", n, n)
}
