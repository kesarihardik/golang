package main

import (
	"fmt"
	"math"
)

//interface - set of method signature

type Abser interface {
	Abs() float64
	Describe()
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 { //implicit implementation of interface by Vertex
	if v == nil {
		return 0
	}
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Describe() {
	if v == nil { //handle nil underlying value
		fmt.Printf("\n%v %T %s", v, v, "Nil type")
	}
	fmt.Printf("\n%v %T %s", v, v, "Not nil type")
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	a = &v //explicit implementation

	describe(a)
	fmt.Printf("\n %f", a.Abs())
	v.Describe()

	var v2 Vertex
	describe(&v2)
	fmt.Printf("\n\n %f", v2.Abs())

}

func describe(i Abser) {
	fmt.Printf("\n%T , %v", i, i) //interface can be thought of as  a tuple (value, type)
}
