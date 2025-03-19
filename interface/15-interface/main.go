package main

import (
	"fmt"
	"math"
)

//interface - set of method signature

type Abser interface {
	Abs() float64
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 { //implicit implementation of interface by Vertex
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	a = &v //explicit implementation
	fmt.Print(a.Abs())
	describe(a)
}

func describe(i Abser) {
	fmt.Printf("\n%T , %v", i, i)
}
