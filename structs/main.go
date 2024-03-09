package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Declaring structs
	// #1
	alex := person{"Alex", "Anderson"}
	// #2
	peter := person{firstName: "Peter", lastName: "England"}
	// #3
	var arnold person
	arnold.firstName = "Arnold"
	arnold.lastName = "Schwarzenegger"

	fmt.Println(alex)
	fmt.Println(peter)
	fmt.Println(arnold)
	fmt.Printf("%+v", arnold)
}
