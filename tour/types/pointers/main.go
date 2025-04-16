package main

import "fmt"

// type *T is a pointer to a T value. Its zero value is nil

func main() {
	i, j := 42, 2701

	p := &i         // i pointer 0x14000104020
	fmt.Println(*p) //read i though the pointer 42

	*p = 21        //set i through the pointer
	fmt.Println(i) // 21

	p = &j         // point to j 0x14000104020
	*p = *p / 37   // divide j thorugh the pointer
	fmt.Println(j) // new value of j 2701 / 37 = 73
}
