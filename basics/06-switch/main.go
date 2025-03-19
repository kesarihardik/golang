package main

import (
	"fmt"
	"time"
)

//Go's switch work with runtime expression and float, string etc.
// break not needed in golang switch

func main() {
	// var year int
	// fmt.Print("Enter any year: ")
	// switch fmt.Scanln(&year); {
	// case year%4 == 1:
	// 	fmt.Println("Not a leap year.")
	// case year%4 == 0 && year%100 != 0:
	// 	fmt.Println("It's a leap year.")
	// case year%400 == 0:
	// 	fmt.Println("Leap year.")
	// default:
	// 	fmt.Println("Not a leap year.")
	// }

	//use switch rather than long if else
	t := time.Now()
	switch {
	case t.Hour() < 12: // Go's switch cases need not be constants, and the values involved need not be integers.
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 21:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good night.")
	}
}
