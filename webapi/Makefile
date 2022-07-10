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
goagen:
	goa gen github.com/is-hoku/goa-template/webapi/design
	rm -rf gen/http/cli gen/http/students/client
.PHONY: build dist/student exec test clean goagen
