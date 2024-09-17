package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter any year: ")
	switch fmt.Scan(&year); {
	case year%4 == 1:
		fmt.Println("Not a leap year.")
	case year%4 == 0 && year%100 != 0:
		fmt.Println("It's a leap year.")
	case year%400 == 0:
		fmt.Println("Leap year.")
	default:
		fmt.Println("Not a leap year.")
	}
}
