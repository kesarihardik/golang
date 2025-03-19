package main

import "fmt"

//Empty interface -- specifies no methods

func main() {
	var i interface{} //empty-interface
	fmt.Printf("\n %T , %v", i, i)

	i = 42 //empty interface can hold any value and type
	fmt.Printf("\n %T , %v", i, i)

	i = "av"
	fmt.Printf("\n %T , %v", i, i)
}
