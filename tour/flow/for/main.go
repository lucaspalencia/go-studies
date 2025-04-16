package main

import "fmt"

// for is the only loop construct in Go

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for is Go's "while"
	// you can drop the semicolons
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// forever loop
	// for {
	// 	//infinite loop
	// }
}
