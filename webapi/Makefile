build: dist/student
dist/student:
	mkdir -p dist
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ $(MAIN_PATH)
test:
	go test ./...
	go vet ./...
clean:
	rm -rf dist
.PHONY: build test clean
