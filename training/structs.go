package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "Giri", lastName: "Naga",
		contactInfo: contactInfo{
			email:   "naga@gmail.com",
			zipCode: 89898,
		},
	}
	alex.print()
	fmt.Println()
	alex.updateName("swetha")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
