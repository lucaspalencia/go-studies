package main

import "fmt"

// range iterates over a slice or map

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// index, value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// if you only want the index
	for i := range pow {
		fmt.Println(i)
	}

	// can skip index or value with _
	for _, value := range pow {
		fmt.Println(value)
	}
}
