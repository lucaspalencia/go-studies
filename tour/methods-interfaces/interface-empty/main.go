package main

import "fmt"

// The interface type that specifies zero methods is known as the empty interface.
// An empty interface may hold values of any type
// For recent versions, you can use any instead of interface{}
// any is an alias for interface{}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
