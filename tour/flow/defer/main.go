package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defers
	// when a function returns, its deferred calls
	// are executed in last-in-first-out order
	fmt.Println("counting")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
