package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of methods signatures
// A type implements an interface by implementing its methods
// There is no explicit declaration of intent, no "implements" keyword.

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I
func (t T) M() {
	fmt.Println(t.S)
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())
}
