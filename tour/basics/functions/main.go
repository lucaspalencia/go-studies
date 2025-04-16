package main

import "fmt"

// or add(x int, y int)
func add(x, y int) int {
	return x + y
}

// multiple returns
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
// they are treated as variables defined at the
// top of the function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// naked return
	return
}

func main() {
	fmt.Println(add(10, 20))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
