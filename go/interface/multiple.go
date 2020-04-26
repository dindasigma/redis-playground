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

type Email struct {
	user string
	address string
}

func (e Email) SendMessage(message string) error {
	_, err := fmt.Printf("Successfully sending email to %s(%s) with content %s.\n", e.user, e.address, message)
	return err
}

func (e Email) Add() error {
	_, err := fmt.Printf("%s, your email %s successfully register with us.\n", e.user, e.address)
	return err
}

func main() {
	e := Email{"John", "john@doe.com"}
	var n Notify = e
	var s Subscribe = e

	n.SendMessage("test")
	s.Add()

	// with assertion
	var notif Notify = Email{"John", "john@doe.com"}
	send := notif.(Email)
	send.SendMessage("test")
	send.Add()
}