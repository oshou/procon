.PHONY: run
run:
	go vet ./...
	golangci-lint run
	go run $1
