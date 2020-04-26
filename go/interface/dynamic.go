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

type Telegram struct {
	user string
	userid string
}

func (e Email) Register() error {
	_, err := fmt.Printf("%s, your email %s successfully register with us.\n", e.user, e.address)
	return err
}

func (e Email) SendMessage(message string) error {
	_, err := fmt.Printf("Successfully sending email to %s(%s) with content %s.\n", e.user, e.address, message)
	return err
}

func (t Telegram) Register() error {
	_, err := fmt.Printf("%s, your telegram account %s successfully register with us.\n", t.user, t.userid)
	return err
}

func (t Telegram) SendMessage(message string) error {
	_, err := fmt.Printf("Successfully sending email to %s(%s) with content %s.\n", t.user, t.userid, message)
	return err
}

func main() {
	var n Notify = Email{"John", "john@doe.com"}
	// first n is main.Email
	fmt.Printf("Type of n is %T\n", n)
	fmt.Printf("Value of n is %v\n", n)
	n.Register()

	n = Telegram{"John", "johndoe"}
	// n is main.Telegram
	fmt.Printf("Type of n is %T\n", n)
	fmt.Printf("Value of n is %v\n", n)
	n.Register()
}