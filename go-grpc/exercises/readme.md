# gRPC Playground

gRPC using HTTP/2.0 as the communication protocol and protobuf as the communication contract between server and client.

Protocol buffers (protobuf) are a data serialization that is extensible machanism for serializing structured data. Once serialized, this protocol produces a binary payload that is smaller in size than a normal JSON payload. This serialized binary payload then travel over the binary transport protocol called HTTP/2.0.

Generate Protobuf:
```
make generate-proto
```

HTTP/2.0 support multiplexing, it allows a single user to have multiple responses, and vice versa.

This repo implements 4 types communication pattern of gRPC: 
- Unary RPC: single request and response. \
  See implementation of `rpc create` and `rpc retrieve`.
  ```
  rpc create(Order) returns (google.protobuf.StringValue);
  rpc retrieve(google.protobuf.StringValue) returns (Order);
  ```
- Server-side Streaming RPC: single request from client with several responses from server.\
   See implementation of `rpc search`.
   ```
   rpc search(google.protobuf.StringValue) returns (stream Order);
   ```
- Client-side Streaming RPC: client have multiple requests and server only return single response.\
  See implementation of `rpc update`.
  ```
  rpc update(stream Order) returns (google.protobuf.StringValue);
  ```
- Bi-directional Streaming RPC: multiple requests and responses within a single connection.\
  See implementation of `rpc process`.
  ```
  rpc process(stream google.protobuf.StringValue) returns (stream CombinedShipment);
  ```
  

## Create private key and the certificate
```
$ openssl genrsa -out certs/server.key 2048
 
$ openssl req -nodes -new -x509 -sha256 -days 1825 -config certs/cert.conf -extensions 'req_ext' -key certs/server.key -out certs/server.crt
```

## Run
Server
```
make run-server
```
Client
```
make run-client
```
## References
[gRPC Up & Running](https://github.com/grpc-up-and-running/samples)

[Using TLS/SSL certificates for gRPC](http://www.inanzzz.com/index.php/post/jo4y/using-tls-ssl-certificates-for-grpc-client-and-server-communications-in-golang-updated)