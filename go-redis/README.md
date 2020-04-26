# Redis-playground

## Running with Docker locally
```
$ docker pull redis
$ docker run --name redis-playground -p 6379:6379 -d redis

$ go run main.go
```