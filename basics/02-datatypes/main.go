package main

import (
	"fmt"
	"math"
)

/*basic types

bool
string
int,int8, int16, rune(int32), int64 ,uint, unint8(byte)....

float32, float64
complex64, complex128

uintptr - unsafe pointer
*/

func main() {

	var isEven, isOdd bool = true, false //variable initialization
	fmt.Println(isEven, isOdd)

	const x int = 123 //short assignment not available for const
	fmt.Printf("\nvalue = %v, type = %T ,int value = %d, binary = %b", x, x, x, x)

	//multi variable initialization
	var (
		f float32   = 231.1
		c complex64 = 3 + 4i //	c complex64 = complex(3, 4)
	)
	fmt.Println("\n", f, c)

	//zero values
	var (
		z1 int
		z2 bool
		z3 string
	)
	fmt.Printf("\nzero values : -%T - %v , %T - %v, %T - %v", z1, z1, z2, z2, z3, z3)

	//go doesn't have char
	// go has byte to represent ASCII and rune to represent unicode points.

	// var aa byte ='₹'          //byte can not represent a unicode character(>1 byte)
	var rupee rune = '₹' //rune are alias for int32. they are untyped constants.
	fmt.Printf("\nrupee = %c, type = %T, Unicode = %U, int value = %d", rupee, rupee, rupee, rupee)

	//Range of datatypes
	fmt.Println("\n\nInteger Ranges in Go:")
	fmt.Printf("int8  : [%d, %d]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("uint8 : [%d, %d]\n", 0, math.MaxUint8)
	fmt.Printf("int : [%d, %d]\n", math.MinInt, math.MaxInt) //int range depends upon system (64 bits on 64 bits system)

	fmt.Println("\n\nFloating-Point Ranges:")
	fmt.Printf("float32: [%e, %e]\n", -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("float64: [%e, %e]\n", -math.MaxFloat64, math.MaxFloat64)
}
