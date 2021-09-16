### To install dependencies:
Keep the package definition in proto relative to the protos folder (as makefile protos command assumes all the protos will be created in protos directory).

Push all the code on github.com assuming the correct import path to the github repo.
```
$ go mod init github.com/vineet070193/go-service1
$ go mod tidy
```

### To start the server, run:
```
$ go run main.go
```

### Install grpcurl to interact with the local grpc server:
```
$ brew install grpcurl
```

### To test the server, run:
```
$ grpcurl --plaintext -d '{"FirstName": "Vineet", "LastName": "Kumar"}' localhost:9092 Chat.Greet
```
You can also run list command to list down the RPC details. For more details, read grpcurl documentation.