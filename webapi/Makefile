GOOS ?= linux
GOARCH ?= amd64
MAIN_PATH ?= ./cmd/sample

build:
	mkdir -p dist
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o dist/sample $(MAIN_PATH)
	touch built
exec: build
	./dist/sample
test:
	go test ./...
	go vet ./...
clean:
	rm -rf dist
goagen:
	goa gen github.com/is-hoku/goa-sample/webapi/design
	rm -rf gen/http/cli/ gen/http/*/client/ gen/*/client*
	touch goagen
goaexample:
	goa example github.com/is-hoku/goa-sample/webapi/design
	rm -rf cmd/*-cli
	touch goaexample
sqlc-generate: datastore/query datastore/table
	sqlc generate --file sqlc.yaml
	touch sqlc-generate
plantuml: built
	goplantuml -recursive -show-aggregations -show-aliases -show-compositions -show-connection-labels -show-implementations ./ > ./diagram.puml
	touch plantuml
inspect:
	atlas schema inspect -u ${ATLAS_DSN} > datastore/schema.hcl
migrate: datastore/schema.hcl
	atlas schema apply -u ${ATLAS_DSN} -f datastore/schema.hcl
	touch migrate
sqlgen:
	go run ./bin/table_def_sql.go
db: migrate sqlgen sqlc-generate

.PHONY: build exec test clean goagen goaexample inspect sqlgen db
