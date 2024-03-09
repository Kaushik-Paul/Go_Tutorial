package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	//// Declaring structs
	//// #1
	//alex := person{"Alex", "Anderson"}
	//// #2
	//peter := person{firstName: "Peter", lastName: "England"}
	//// #3
	//var arnold person
	//arnold.firstName = "Arnold"
	//arnold.lastName = "Schwarzenegger"
	//
	//fmt.Println(alex)
	//fmt.Println(peter)
	//fmt.Println(arnold)
	//fmt.Printf("%+v", arnold)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 895678,
		},
	}

	fmt.Printf("%+v\n", jim)
	jim.print()

	// Use pointer to update struct
	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()

	// Another way
	jim.updateFirstName("Jimmy22")
	jim.print()
}
