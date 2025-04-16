package main

import "fmt"

// Generic types
type List[T any] struct {
	next *List[T]
	val  T
}

// Index returns the index of x in s, or -1 if not found.
// comparable is a useful constraint that makes it possible
// to use the == and != operators on values of the type
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}

	return -1
}

func main() {
	ll := &List[int]{val: 2}
	l := List[int]{val: 1, next: ll}
	fmt.Println(l)
	fmt.Println(l.next.val)

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
