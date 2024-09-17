package main

import "fmt"

func main() {
	defer fmt.Println("Take care !")
	defer fmt.Println("Bye") //defer is executed in LIFO order
	fmt.Println("Hi there!")

	deff()
}
