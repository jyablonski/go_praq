package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// c is a pointer of type Counter
// Since c is a pointer, it directly modifies the original Counter instance.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
