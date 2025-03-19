package main

import (
	"fmt"
	"math/rand"
)

// package in go is way to group functions.
// only one main method in one directory allowed.
func main() {
	//only methods starting with upper case letter are public like Println, Printf
	fmt.Printf("Random number: %d", rand.Int())
}
