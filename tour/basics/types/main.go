package main

import (
	"fmt"
	"math/cmplx"
)

/*
go basic types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte - alias for uint8

rune - alias for int32
  - represents a Unicode code point

float32 float64

complex64 complex128
*/

// varaible declarations can be factored into blocks
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversions
	// T(v) converts the value v to the type T
	x := 42
	y := float64(x)
	u := uint(y)

	fmt.Println(x, y, u)

	// type inference
	// variable's type is inferred form the value if
	// not specifying an explicit type

	var l int
	j := l // j is an int
	fmt.Println(j)

	v := "dsada" // change here and see the type change
	fmt.Printf("v is of type %T\n", v)
}
