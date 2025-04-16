package main

import "fmt"

// The example code sums the numbers in a slice, distributing the work
// between two goroutines. Once both goroutines have completed their
// computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// like maps and slices, channels must be created before use
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// buffered channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 - fatal error deadlock as channel was created with 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
