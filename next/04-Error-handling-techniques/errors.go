package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	debug.PrintStack()
	defer debug.PrintStack()

	done := make(chan bool)
	go func(done chan bool) {
		debug.PrintStack()
		done <- true
		close(done)
	}(done)

	<-done
	stackTrace := debug.Stack()
	fmt.Printf("%v \n", string(stackTrace))
}
