package main

import "sync"

// Counter a simple implementation of a counter
type Counter struct {
	mu    sync.Mutex // Mutex is mutal exclusion lock (0 = unlock)
	value int
}

// NewCounter constructor function
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter
func (c *Counter) Inc() {
	// locking and unlocking the Counter to avoid
	// multiple mutations at the same time
	// caused by the goroutines
	// each goroutine will acquire the lock on Counter if they are first
	// all the other goroutines will have to wait for the unlock
	c.mu.Lock()
	defer c.mu.Unlock() // this will be run after `c.value++`
	c.value++
}

// Value returns the value of the counter
func (c *Counter) Value() int {
	return c.value
}
