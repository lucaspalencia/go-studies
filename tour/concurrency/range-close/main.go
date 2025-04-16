package main

import "fmt"

/*
	A sender can close a channel to indicate that no more values will be sent
	Only the sender should close a channel
	Sending value to a closed channel will cause a panic

	Receivers can test wheter a channel has been closed:
	v , ok := <- ch // ok = false if the channel is closed
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for range n {
		c <- x
		x, y = y, x+y
	}

	// closing channel
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// the loop receives values from the channel
	// repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}
}
