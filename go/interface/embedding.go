package main

import (
	"fmt"
)

type Notify interface {
	SendMessage(message string) error
}

type Subscribe interface {
	Add() error
}

type Broadcast interface {
	Notify
	Subscribe
}

type Email struct {
	user string
	address string
}

func (e Email) Add() error {
	_, err := fmt.Printf("%s, your email %s successfully register with us.\n", e.user, e.address)
	return err
}

func (e Email) SendMessage(message string) error {
	_, err := fmt.Printf("Successfully sending email to %s(%s) with content %s.\n", e.user, e.address, message)
	return err
}

func main() {
	e := Email{"John", "john@doe.com"}

	var b Broadcast = e
	var n Notify = e
	var s Subscribe = e
	
	fmt.Printf("Dynamic type and value of interface n type Notify is %T and %v\n", n, n)
	fmt.Printf("Dynamic type and value of interface s type Notify is %T and %v\n", s, s)
	fmt.Printf("Dynamic type and value of interface b type Notify is %T and %v\n", b, b)
}