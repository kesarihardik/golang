package main

import "fmt"

func main() {
	defer fmt.Println("Closed connection.")
	defer fmt.Println("Closing connection!") //defer is executed in LIFO order

	fmt.Println("Opening connection!")

	deff()
}
