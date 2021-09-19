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


### To test feed server, run:
```
$ grpcurl --plaintext localhost:9092 describe FeedService
FeedService is a service:
service FeedService {
  rpc AddFeed ( .AddFeedRequest ) returns ( .AddFeedResponse );
  rpc ListFeed ( .ListFeedRequest ) returns ( .ListFeedResponse );
}
$ grpcurl --plaintext -d '{"FeedID": "Priyanka"}' localhost:9092 FeedService.ListFeed
{
  "FeedList": [
    {
      "FeedID": "FeedID1",
      "Snippet": "Snippet1",
      "Content": "Content",
      "Timestamp": "123456"
    }
  ],
  "Status": {
    "Status": true
  }
}
$ grpcurl --plaintext -d '{"Feed": {"FeedID": "feed_id", "Snippet": "snippet1", "Content": "feed content", "Timestamp": 12345}}' localhost:9092 FeedService.AddFeed
{
  "Status": {
    "Status": true
  }
}
```

## Error Troubleshooting
When getting:
```
[ERROR] Error creating datastoreClient: : EXTRA_VALUE_AT_END="dialing: google: could not find default credentials. See https://developers.google.com/accounts/docs/application-default-credentials for more information."
```

Try running the below command to refresh the credentials on cli and use default service account for the project:
```
$ gcloud auth application-default login
```