package main

import "fmt"

func main() {
	fmt.Println("Enter your age.")
	var age int8
	fmt.Scanln(&age)

	if adult := 18; age < int8(adult) {
		fmt.Println("Not eligible to vote.")
	} else {
		fmt.Println("Eligible to vote.")
	}
}
