build:
	@go build -o bin/img src/main.go
	

test:
	@go test -v ./...

run: build
	@./bin/img