package main

import (
	"errors"
	"fmt"
	"math"
)

// if params are of same type, we can define type once
// in go functions can return multiple values
func divide(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator can not be zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

//naked return (not recommended for larger functions)

// func Divide2(numerator, denominator int) (quotient int, remainder int, er error) {
// 	if denominator == 0 {
// 		quotient = 0
// 		remainder = 0
// 		er = errors.New("can not divide by zero")
// 		return
// 	}
// 	er = nil
// 	quotient = numerator / denominator
// 	remainder = numerator % denominator
// 	return
// }

func sum(a ...int) int { //variadic functions
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println(divide(5, 0))
	fmt.Println(sum(11, 2, 45, 3))

	//functions can be used as parameters and passed around
	hypot := func(x, y float64) float64 { //anonymous function/ function literal assigned to hypot
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("hypot type: %T, value: %f \n", hypot, hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
