package main

import "fmt"

// go function may be closures
// closures are functions values that reference variables from outside their body
// the function may access and assign to the referenced variables; in this sense the function is "bound" to the variables

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
