package main

import "fmt"

func main() {
	defer fmt.Println("program will not crash")

	defer func() {
		fmt.Println( recover() ) // 3
	}()

	defer fmt.Println("now, panic 3 suppresses panic 2")
	defer panic(3)
	defer fmt.Println("now, panic 2 suppresses panic 1")
	defer panic(2)
	panic(1)
}
