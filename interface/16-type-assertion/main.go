package main

import "fmt"

func main() {
	var i interface{} = 42

	a, ok := i.(int)
	fmt.Printf("\n %v %v", a, ok)

	b, ok := i.(string)
	fmt.Printf("\n %v %v", b, ok)
}
