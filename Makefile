format:
	gofumpt -l -w .

test:
	go test -v ./...

lint:
	golangci-lint run