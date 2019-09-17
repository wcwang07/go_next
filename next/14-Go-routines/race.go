package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	//time.Sleep(1 * time.Millisecond)
	return i
}
