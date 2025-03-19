package main

import "fmt"

func main() {
	fmt.Println("Enter your age - ")
	var age int
	fmt.Scanln(&age)

	if adult := 18; age < adult { //variable can be initialized in if and are accessible inside if-else blocks
		fmt.Printf("Not eligible to vote since age less than %d.", adult)
	} else {
		fmt.Printf("Eligible to vote since age greater than %d.", adult)
	}
}
