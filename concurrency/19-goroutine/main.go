package main

import (
	"fmt"
	"time"
)

//goroutine is light weight thread managed by go runtime.
// It runs in same memory space, so access to shared memory must be synchronized.

func print(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go print("World") //execution order is non-deterministic. Go scheduler decides when goroutine is run.
	print("Hello")
}
