package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func withStruct(jsonData []byte) (*Person, error) {
	var person Person
	err := json.Unmarshal(jsonData, &person)

	return &person, err
}

func withInterface(jsonData []byte) (interface{}, error) {
	var person interface{}
	err := json.Unmarshal(jsonData, &person)

	return person, err
}

func main() {
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	// with struct
	person, err := withStruct(jsonData)
	fmt.Println(person.Name, err)

	// with interface
	person2, err := withInterface(jsonData)
	data := person2.(map[string]interface{})
	fmt.Println(data["Name"], err)

	// both can be encoded to json
	json.NewEncoder(os.Stdout).Encode(person)
	// output: {"Name":"Eve","Age":6}

	json.NewEncoder(os.Stdout).Encode(data)
	// output: {"Age":6,"Name":"Eve","Parents":["Alice","Bob"]}

	// another example
	result := map[string]interface{}{
		"Age":  6,
		"Name": "Eve",
	}
	json.NewEncoder(os.Stdout).Encode(result)
	// output: {"Age":6,"Name":"Eve"}
}
