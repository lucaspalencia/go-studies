package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Lucas", "Talita", "Noah", "Zion"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	message, err2 := greetings.Hello("Lucas")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(message)
}
