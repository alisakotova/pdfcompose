.PHONY: proto

proto:
	protoc proto/*.proto --go-grpc_out=pkg --go_out=pkg