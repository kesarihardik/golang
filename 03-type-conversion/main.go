package main

import "fmt"

func main() {
	//unlike C type conversion is explicit in go.
	//go also forbids operations between different types.

	//int to float32
	var x int = 65
	var f float32 = float32(x)
	fmt.Println(f)

	//typed and untyped constants
	const c1 = 12
	const c2 int = 12

	var y float32 = c1
	//var z float32 = c2  // implicit conversion not possible in case of typed constant
	fmt.Println(c1, c2, y)

	// bool to integer & vice-versa is not available
	var b bool = false
	var bitConvert int8
	if b {
		bitConvert = 1
	}
	fmt.Println(bitConvert)

	//int to bool
	var boolConvert bool = bitConvert != 0
	fmt.Println(boolConvert)

}
