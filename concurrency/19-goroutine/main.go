package main

import (
	"fmt"
	"time"
)

//goroutine is light weight thread managed by go runtime.
// It runs in same memory space, so access to shared memory must be synchronized.

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go print("World")
	print("Hello")
}
