package main

//select lets goroutine wait on multiple communication operations.
// must define default case. A select blocks until one of its cases can run.
import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { //select will cause the goroutine to wait.
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Printf("\n Nothing in channel currently. Blocking goroutine till channel is ready.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for range 10 {
			fmt.Println("\n", <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
