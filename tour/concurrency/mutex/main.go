package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	If we just want to make sure only one goroutine can access
	a variable at a time to avoid conflicts?

	This concept is called mutual exclusion, and the conventional name
	for the data structure that provides it is mutex.

	sync.Mutex provides mutual exclusions and two methods Lock and Unlock
*/

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter fo the given key
func (c *SafeCounter) Inc(key string) {
	// Lock, so only one goroutine at a time can access the map c.v
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for range 1000 {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
