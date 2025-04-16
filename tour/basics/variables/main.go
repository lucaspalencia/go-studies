package main

import "fmt"

// var defined in package level
var c, python, java bool

// var with initializers
var i2, j int = 1, 2

// short declaration cannot be used in package level
// e.g: k:= 3

func main() {
	// var defined in function level
	var i int
	fmt.Println(i, c, python, java)

	fmt.Println(i2, j)

	// short declaration
	k := 3
	fmt.Println(k)
}
