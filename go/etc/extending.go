package main

import (
	"fmt"
)

type User struct {
	id int
	name string
}

type UserAndCompany struct {
	User
	company string
	addres string
}

func main() {
	person := User {id: 1, name: "John Doe"}
	personCompany := &UserAndCompany{User: person, company: "PT. Company"}
	personCompany.addres = "Jl. A"

	fmt.Println(person, personCompany.company, personCompany.addres)
	// output: {1 John Doe} PT. Company Jl. A
}