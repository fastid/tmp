BINARY_NAME=fastid

.PHONY: help
help:
	@echo ""
	@echo "make build     - Building a project"
	@echo "make run       - Project run"
	@echo "make clean     - Cleans up temporary files"
	@echo "make br        - Building and run a project"
	@echo "make test      - Runs tests"
	@echo "make report    - Creates a test coverage reporter"
	@echo ""

.PHONY: build
build:
	GOARCH=amd64 go build -o ${BINARY_NAME} cmd/api/main.go

.PHONY: run
run:
	./${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f coverage.out
	rm -f coverage.html

.PHONY: br
br: build run

.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: migrate_up
migrate_up:
	go run cmd/fastid
