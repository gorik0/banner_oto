.PHONY: proto
proto:
	protoc --go_out=./gen/ --go-grpc_out=./gen/ proto/*/*.proto