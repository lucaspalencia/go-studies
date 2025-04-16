package main

import (
	"fmt"
	"strings"
)

// an array has a fixed sice
// a slice is a dynamically sized, much more
// common than arrays
// []T is a slice with elements of type T

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(s)

	// slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literals
	q := []int{2, 3, 5, 7}
	fmt.Println(q)

	r := []bool{true, false, true, false}
	fmt.Println(r)

	bb := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}

	fmt.Println(bb)

	//slice bounds
	ss := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(ss)

	ss = ss[1:4]
	fmt.Println(ss) // [3, 5, 7]

	ss = ss[:2]
	fmt.Println(ss) // [3, 5]

	ss = ss[1:]
	fmt.Println(ss) // [5]

	// nil slices
	var sss []int
	fmt.Println(sss, len(sss), cap(sss))

	if sss == nil {
		fmt.Println("nil!")
	}

	// creating a slice with make
	m := make([]int, 5)
	println("m", m)

	n := make([]int, 0, 5)
	printSlice("n", n)

	c := m[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	//slice of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append to a slice
	var sa []int
	fmt.Println(sa)

	// append works on nil slices.
	sa = append(sa, 0)
	fmt.Println(sa)

	// The slice grows as needed.
	sa = append(sa, 1)
	fmt.Println(sa)

	// We can add more than one element at a time.
	sa = append(sa, 2, 3, 4)
	fmt.Println(sa)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
