package main

import "fmt"

/*
	The select statement lets a goroutine wait
	on multiple communication operations.

	A select blocks until one of its cases can run,
	then it executes that case
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("===== c <- x =====")
			x, y = y, x+y
		case <-quit:
			fmt.Println("===== quit ====")
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	fmt.Println("====== main ======")
	go func() {
		for range 10 {
			fmt.Println("====== range =====")
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fmt.Println("===== calling fibonacci =====")
	fibonacci(c, quit)
}
