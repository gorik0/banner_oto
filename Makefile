.PHONY: proto
proto:
	protoc --go_out=./gen/ proto/*/*.proto