package main

import (
	"fmt"
)

type Notify interface {
	Register() error
	SendMessage(message string) error
}

type Email struct {
	user string
	address string
}

func (e Email) Register() error {
	_, err := fmt.Printf("%s, your email %s successfully register with us.\n", e.user, e.address)
	return err
}

func (e Email) SendMessage(message string) error {
	_, err := fmt.Printf("Successfully sending email to %s(%s) with content %s.\n", e.user, e.address, message)
	return err
}

func main() {
	var n Notify
	// first n is nil
	fmt.Printf("Type of s is %T\n", n)

	n = Email{"John", "john@doe.com"}
	// n is main.Email
	fmt.Printf("Type of n is %T\n", n)
	fmt.Printf("Value of s is %v\n", n)
	
	n.Register()
}