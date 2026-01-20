package main

import "fmt"

func main() {
	// var firstName string = "John"
	// var lastName string
	// lastName = "Wick"
	// lastName := "Wick"

	// var firstName, lastName string = "John", "Wick"
	firstName, lastName := "John", "Wick"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)
}
