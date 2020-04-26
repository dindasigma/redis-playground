package playground

import (
	"fmt"
)

func (server *Server) Basic() {
	var err error
	// we can call set with a key and a value
	err = server.RDB.Set("key", "value", 0).Err()

	// if there has been an error setting the value
	// handle the error
	if err != nil {
		panic(err)
	}

	val, err := server.RDB.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}