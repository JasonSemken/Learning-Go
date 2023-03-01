package main

import (
	"fmt"
)

func main() {

	var firstName string
	var lastName string

	// Requests user input of names
	fmt.Println("What is your firstname?")
	fmt.Scan(&firstName)
	fmt.Println("What is your lastname?")
	fmt.Scan(&lastName)

	// prints users name like they are James Bond
	fmt.Printf("Your name is %v, %v %v.\n", lastName, firstName, lastName)

}
