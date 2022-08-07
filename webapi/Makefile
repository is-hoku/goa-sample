GOOS ?= linux
GOARCH ?= amd64
MAIN_PATH ?= ./cmd/sample

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
	goa gen github.com/is-hoku/goa-sample/webapi/design
	rm -rf gen/http/cli/ gen/http/*/client/ gen/*/client*
goaexample:
	goa example github.com/is-hoku/goa-sample/webapi/design
	rm -rf cmd/*-cli
create_ddl:
	go run cmd/ddl/main.go -d mysql
.PHONY: build dist/student exec test clean goagen goaexample create_ddl
