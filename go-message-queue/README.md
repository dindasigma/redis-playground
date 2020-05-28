# Message Queue with Go

## Why we need message queue
In microservices, no single service go alone without cooperation from other services, it also getting more complex as the requirements increase and the system evolves. In message queue, the producer can send data to pool, then active consumer will get the data. Also, we need data persistance to ensure the data not gone until received by the consumer. What we expect from message queue:
- If some nodes go down, we can still use the messaging service
- A mechanism with retry (and delay retry upon failure)
- Pub-sub (publish-subscribe) pattern

## Candidate
### RabbitMQ
RabbitMQ use standard protocol AMQP and has some exchange type between producer (P) and consumer (C) which are direct, pub-sub/broadcast, route and RPC.

![rabbitmq](https://blog.dinda.id/images/uploads/rabbitmq.png)

### NSQ
NSQ has one concise message model which simple and direct. All you need is define the producer and the consumer. If you need pub-sub or broadcast, you can add multiple channels for the topic.

![nsq](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

NSQD daemon is the core part of NSQ. It’s a standalone binary that listens for incoming messages on a single port.

## Pros and Cons
### Pros of RabbitMQ
- Based on the open standard protocol AMQP
- Mature and stable
- Introduce exchange between producer and consumer and fledged with lots of patterns (Direct/Worker/Pub-Sub/Route/RPC…)

### Cons of RabbitMQ
- A bit of steep learning curve
- Not easy to implement retry with failure (using [dead letter exchange](https://www.rabbitmq.com/dlx.html))

### Pros of NSQ
- Written in Go
- Message model is simple and direct
- Writing into NSQD can be as simple as performing an HTTP POST request to the endpoint
- Supports retry with delay naturally. When message handler indicates failure, that information is sent to nsqd in the form of a REQ (re-queue) command. Also, nsqd will automatically time out (and re-queue) a message if it hasn't been responded to in a configurable time window. We also can call `requeue()` with parameter of delay.

### Cons of NSQ
- There’s no concept of routing based upon on key

## Implementation

### RabbitMQ
- [Direct](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/basic)
- [Direct with queues](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/queues)
- [Pub-sub/Broadcast](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/pubsub)
- [Routing based upon on key](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/routing)
- [Routing based upon some pattern](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/topics)
- [RPC](https://github.com/dindasigma/my-playground/tree/master/go-message-queue/go-rabbitmq/exercise/rpc)

### NSQ
- Direct (run [producer.go](https://github.com/dindasigma/my-playground/blob/master/go-message-queue/go-nsq/producer.go) and [consumer1.go](https://github.com/dindasigma/my-playground/blob/master/go-message-queue/go-nsq/consumer1.go))
- Subscribe (run [producer.go](https://github.com/dindasigma/my-playground/blob/master/go-message-queue/go-nsq/producer.go), [consumer1.go](https://github.com/dindasigma/my-playground/blob/master/go-message-queue/go-nsq/consumer1.go), and [consumer2.go](https://github.com/dindasigma/my-playground/blob/master/go-message-queue/go-nsq/consumer2.go))

## References
- [https://www.rabbitmq.com/getstarted.html](https://www.rabbitmq.com/getstarted.html)
- [https://godoc.org/github.com/nsqio/go-nsq](https://godoc.org/github.com/nsqio/go-nsq)
- [https://nsq.io/overview/features_and_guarantees.html](https://nsq.io/overview/features_and_guarantees.html)
- [https://wiredcraft.com/blog/building-microservices-nsq-rabbitmq/](https://wiredcraft.com/blog/building-microservices-nsq-rabbitmq/)
- [https://segment.com/blog/scaling-nsq/](https://segment.com/blog/scaling-nsq/)
