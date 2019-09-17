APP=crossapp

all: clean  build build-win

.PHONY: run
run: build
	./$(APP)

.PHONY: build
build:
	go build -o $(APP) 

.PHONY: build-win
build-win:
	GOOS=windows GOARCH=386 go build -o $(APP).exe
	 
.PHONY: debs
debs:
	go get ./...

.PHONY: clean
clean:
	go clean
	rm -f $(APP)