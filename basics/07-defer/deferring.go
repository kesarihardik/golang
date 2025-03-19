package main

import "fmt"

func deff() {
	fmt.Println("counting")

	//deferred function calls are pushed on stack. when a func call returns def func calls are executed in LIFO manner
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
