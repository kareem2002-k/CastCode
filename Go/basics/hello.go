package main

import (
	"fmt"
	"log"
	"tut/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
