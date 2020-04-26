package playground

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type Server struct {
	RDB	*redis.Client
}

func (server *Server) Init()  {
	server.RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password set
		DB: 0, // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	_, err := server.RDB.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("we are connected")
	}
	
}