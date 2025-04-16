package main

import (
	"fmt"
	"math"
)

// Go does not have classes. You can defined methods on types.
// a method is just a function with a receiver argument

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receivers
// Methods with pointer receivers can modify the value to which the receiver points
// Since methos often need to modify their receiver, pointer are more common.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs()) // 50
}
