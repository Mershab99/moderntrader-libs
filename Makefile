export PATH := $$PATH:$(shell go env GOPATH)/bin

generate-proto-go:
	rm -rf ./go
	mkdir go
	protoc --go_out=./go --go_opt=paths=source_relative ./protobuf/kafka.proto

generate-api-go:
	docker pull openapitools/openapi-generator-cli
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api/api-spec.yaml -g go -o /local/go/api --global-property models,supportingFiles

