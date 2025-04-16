package main

import "fmt"

// struct is a collection of fields
type Vertex struct {
	X, Y int
}

// struct literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// acessing struct fields
	v := Vertex{1, 2}
	fmt.Println(v.X) // 1
	v.X = 4
	fmt.Println(v.X) // 4

	fmt.Println(v1, v2, v3, p)
}
