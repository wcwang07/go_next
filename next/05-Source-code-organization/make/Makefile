APP=app

all: fmt test build build-win32

.PHONY: run
run: build
	./$(APP)

.PHONY: build
build:
	go build -o $(APP) -v ./...

.PHONY: build-win32
build-win32: 
	GOOS=windows GOARCH=386 go build -o $(APP).exe

.PHONY: test
test:
	go test -timeout 20s -race -v ./...

.PHONY: debs
debs:
	go get ./...

.PHONY: fmt
fmt:
	gofmt -l -s .

.PHONY: clean
clean:
	go clean
	rm -f $(APP)