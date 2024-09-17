package main

import (
	"fmt"
)

// bool
// int,int8, int16,int32,int64 , unint, float32, float64, complex64, complex128
// string

// untyped variables and constants take their type as per their context
// alias: byte = uint8, rune = int32
// uintptr - unsafe pointer

// int, unint, uintptr - default size 32 bit on 32 bit system
//zero values:  0 for int, "" for string & false for boolean

func main() {
	var isEven, isOdd bool = true, false //variable initialization
	fmt.Println(isEven, isOdd)

	const x int = 123 //short assignment not available for const
	fmt.Println(x)

	var (
		f float32   = 231.1
		c complex64 = 3 + 4i //	c complex64 = complex(3, 4)
	)
	fmt.Println(f, c)

	//go doesn't have char
	//rune are alias for int32. they are untyped constants so their type can change.
	//they represent unicode codepoints instead of ASCII

	// sp := byte('♛')                   // a byte (uint8) can not represent such characters due to overflow.
	sp := '♛'                             //sp := rune('♛')
	fmt.Printf("%c,  %d, %T", sp, sp, sp) //%T is used for type inference
}
