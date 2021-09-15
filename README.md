To start the server, run:
```
$ go run main.go
```


To test the server, run:
```
$ grpcurl --plaintext -d '{"FirstName": "Vineet", "LastName": "Kumar"}' localhost:9092 Chat.Greet
```