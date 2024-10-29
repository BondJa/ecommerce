build:
	@go build -o bin/ecommerce cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/ecommerce