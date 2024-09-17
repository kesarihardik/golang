package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
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
	var wg sync.WaitGroup
	c := Counter{count: 0}

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go c.increment(&wg)
	}

	wg.Wait()
	fmt.Print(c.fetchCount())
}
