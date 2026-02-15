BINARY=avro
BUILD_DIR=bin
MAIN=./cmd/avro

.PHONY: build run test vet lint clean install

build:
	go build -o $(BUILD_DIR)/$(BINARY) $(MAIN)

run:
	go run $(MAIN)

test:
	go test ./...

vet:
	go vet ./...

lint: vet
	@echo "Lint passed (go vet)"

clean:
	rm -rf $(BUILD_DIR)

install: build
	cp $(BUILD_DIR)/$(BINARY) $(GOPATH)/bin/$(BINARY) 2>/dev/null || \
	cp $(BUILD_DIR)/$(BINARY) $(HOME)/go/bin/$(BINARY)

# Cross-platform builds
build-all:
	GOOS=darwin  GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-darwin-amd64 $(MAIN)
	GOOS=darwin  GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY)-darwin-arm64 $(MAIN)
	GOOS=linux   GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe $(MAIN)
