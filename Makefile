.PHONY: protos

protos:
	protoc -I protos/ protos/chat.proto --go_out=plugins=grpc:protos/
	protoc -I protos/ protos/feed.proto --go_out=plugins=grpc:protos/