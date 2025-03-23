package main

import "fmt"

//Unlike functions, methods can take value / pointer types irrespective of value/pointer receiver specified.
//all methods on a given type should have either pointer or value receiver, not both.

type Vertex struct {
	X, Y float64
}

// pointer receiver methods are recommended to save memory and modify original values.
func (v *Vertex) Scale(f float64) { //works with both pointer and value types. Go automatically converts value to ptr.
	v.X *= f
	v.Y *= f
}

// func (v Vertex) Scale2(f float64) {
// 	v.X *= f
// 	v.Y *= f
// }

func ScaleFunc(v Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) //(&v).Scale(10)
	fmt.Printf("%v", v)

	v2 := Vertex{3, 4}
	ScaleFunc(v, 3)
	fmt.Printf("\n after calling function v - %v", v2)

	p := &v2
	p.Scale(5)
	fmt.Printf("\n after calling method scale, v - %v", v2)

	// p.Scale2(3)
	// fmt.Printf("\n after calling method scale2, v - %v", v2)

}
