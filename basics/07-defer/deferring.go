package main

import "fmt"

func deff() {
	fmt.Println("Counting starting.")
	defer fmt.Println("Closing counter.")
	//deferred function calls are pushed on stack. when a func call returns def func calls are executed in LIFO manner
	for i := range 10 {
		defer fmt.Println(i)
	}
}
