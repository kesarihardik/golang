package main

import "fmt"

//nil interface - no concrete value and no concrete type.
//If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

type I interface {
	M()
}

func main() {
	var i I //nil interface
	fmt.Printf("%T, %v", i, i)
	//i.M()                // segmentation fault - go doesn't know which M() to call
}
