.PHONY: protogen start

protogen:
	@cd protodef && buf generate && cd -

start:
	@go run cmd/main.go

tidy:
	@gofumpt -l -w **/*.go
	@go mod tidy
