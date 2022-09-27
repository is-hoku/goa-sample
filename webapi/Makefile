GOOS ?= linux
GOARCH ?= amd64
MAIN_PATH ?= ./cmd/sample

build:
	mkdir -p dist
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o dist/sample $(MAIN_PATH)
exec: build
	./dist/sample
test:
	go test ./...
	go vet ./...
clean:
	rm -rf dist
goagen:
	mv gen/student/error.go .
	goa gen github.com/is-hoku/goa-sample/webapi/design
	rm -rf gen/http/cli/ gen/http/*/client/ gen/*/client*
	mv ./error.go gen/student
goaexample:
	goa example github.com/is-hoku/goa-sample/webapi/design
	rm -rf cmd/*-cli
inspectdb: datastore/table/students.sql
	atlas schema inspect -u ${ATLAS_DSN} > datastore/schema.hcl
.PHONY: build exec test clean goagen goaexample inspectdb
