package main

import (
	"fmt"
	"log"

	"example/greetings"
)


func main() {
	//Set properties of the predefined Logger including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and the line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	//Get a greeting message and print it
	message, err  := greetings.Hello("")
	//If an error was returned, print it to the console and
	//exit the program
	if err != nil {
		log.Fatal(err)
	}
	//If no error was returned, print the returned message
	// to the consile
	fmt.Println(message)
}
