# Message Queue with Go

## Why we need message queue
In microservices, no single service can go it alone without cooperation from other services as the requirements increase in volume and complexity and the system evolves. In message queue, the producer can send data to pool, then active consumer can get the data. Also, we need data persistance to ensure the data not gone until received by the consumer. If some nodes go down, you can still use the messaging service. What we expect from message queue:
- A mechanism with retry (and delay retry upon failure)
- Pub-sub (publish-subscribe) pattern

## Candidate
### RabbitMQ
RabbitMQ use standard protocol AMQP and has some exchange type between producer (P) and consumer (C) which are Direct, Worker, Pub-Sub, Route and RPC.

![rabbitmq](https://blog.dinda.id/images/uploads/rabbitmq.png)

### NSQ
NSQ has one concise message model which simple and direct. All you need is define the producer and the consumer. If you need pubsub or broadcast, you can add multiple channels for the topic.

![nsq](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

NSQD daemon is the core part of NSQ. It’s a standalone binary that listens for incoming messages on a single port.

## Pros and Cons
### Pros of RabbitMQ
- Based on the open standard protocol AMQP
- Mature and stable
- Introduce exchange between producer and consumer and fledged with lots of patterns (Direct/Worker/Pub-Sub/Route/RPC…)

### Cons of RabbitMQ
- A bit of steep learning curve

### Pros of NSQ
- Written in Go
- Message model is simple and direct
- Writing into NSQD can be as simple as performing an HTTP POST request to the /put endpoint
- Supports retry with delay naturally. When message handler indicates failure, that information is sent to nsqd in the form of a REQ (re-queue) command. Also, nsqd will automatically time out (and re-queue) a message if it hasn't been responded to in a configurable time window.

### Cons of NSQ
- There’s no concept of routing based upon on key

## Implementation

### RabbitMQ
- Direct
- Direct with queues
- Subscribe
- Routing based upon on key
- Routing based upon some pattern
- RPC

### NSQ
- Direct
- Subscribe