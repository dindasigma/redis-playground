# RabbitMQ-playground

## Running with Docker locally
```
$ docker pull rabbitmq:3-management
$ docker run -d -p 15672:15672 -p 5672:5672 --name go-rabbitmq rabbitmq:3-management
```
`15672` is the default port for RabbitMQ GUI, `5672` for RabbitMQ message broker.

```
$ go run basic.go
```

```
$ go run sensor.go
```