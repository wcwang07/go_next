APP=crossapp

all: clean fmt test build build-win

.PHONY: run
run: build
	./$(APP)

.PHONY: doc
doc:
	godoc -http=:6060 -index

.PHONY: build
build:
	go build -o $(APP) 

.PHONY: build-win
build-win:
	GOOS=windows GOARCH=386 go build -o $(APP).exe
	 
.PHONY: test
test:
	go test -timeout 20s -race 

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