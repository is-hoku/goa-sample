GOOS ?= linux
GOARCH ?= amd64
MAIN_PATH ?= ./cmd/student

build: dist/student
dist/student:
	mkdir -p dist
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ $(MAIN_PATH)
exec: build
	./dist/student
test:
	go test ./...
	go vet ./...
clean:
	rm -rf dist
.PHONY: build test clean run
