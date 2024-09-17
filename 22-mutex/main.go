package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock() // without lock the counter will not behave consistently
	c.count++
	c.mu.Unlock()
}

func (c *Counter) fetchCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	c := Counter{count: 0}

	for i := 1; i <= 1000; i++ {
		go c.increment()
	}

	// time.Sleep(time.Second) //sleep is important so that main waits for execution of all go routines.
	fmt.Print(c.fetchCount())
}
