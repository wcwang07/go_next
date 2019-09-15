
APP=cgo

all: clean  build-c++ build-go

.PHONY: clean
clean:
	go clean
	rm -f $(APP)

.PHONY: build-c++
build-c++:
	gcc -c greeter.c
	
.PHONY: build-go
build-go:
	go build -o $(APP)


