.PHONY: protos

protos:
	protoc -I protos/ protos/chat.proto --go_out=plugins=grpc:protos/chat