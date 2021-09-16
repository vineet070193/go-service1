To start the server, run:
```
$ go run main.go
```

Install grpcurl to interact with the local grpc server:
```
brew install grpcurl
```

To test the server, run:
```
$ grpcurl --plaintext -d '{"FirstName": "Vineet", "LastName": "Kumar"}' localhost:9092 Chat.Greet
```