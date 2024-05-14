export PATH := $$PATH:$(shell go env GOPATH)/bin

generate-proto-go:
	rm -rf ./go
	mkdir go
	protoc --go_out=./go --go_opt=paths=source_relative ./protobuf/kafka.proto

