package main

import "fmt"

//function with value takes value types and pointer takes pointer types.
//methods can take value / pointer types irrespective of value/pointer arguments specified.
//all methods on a given type should have either pointer or value receiver, not both.

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) { //works with both pointer and value types. Go automatically converts value to ptr.
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) //(&v).Scale(10)
	fmt.Printf("%v", v)

	v2 := Vertex{3, 4}
	p := &v2
	p.Scale(5)
	fmt.Printf("%v", v2)

	// ScaleFunc(v,3)               //can not pass value types
}
