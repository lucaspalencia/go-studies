package main

import "fmt"

// type assertion provides access to an interface value's
// t := i.(T)

func main() {
	var i any = "hello"

	s := i.(string)
	fmt.Println(s)

	// type assertion can return two values: the underlying value,
	// and a boolean tha reports whether the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// if i does not hold a T, it will trigger a panic
	f = i.(float64) // panic
	fmt.Println(f)
}
