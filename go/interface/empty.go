package main

import "fmt"

type MyString string

type Email struct {
	user string
	address string
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
	ms := MyString("Hello World!")
	r := Email{"John", "john@doe.com"}
	
	// value given to explain function is of type 'main.MyString' with value Hello World!
	explain(ms)

	// value given to explain function is of type 'main.Email' with value {John john@doe.com}
	explain(r)

	// so empty interface is dynamic value which point to whatever value we have passed to the function as the argument

}
