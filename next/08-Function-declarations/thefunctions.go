package main

import "fmt"

type Handler func() *int
type VarHandler func(...int)
type IntHandler func(int) int

func compose(a IntHandler, b IntHandler) IntHandler {
	return func(c int) int {
		return a(b(c))
	}
}

func init() {
	fmt.Print("call-1")
}
func init() {
	fmt.Print("call-12")
}
func main() {
	var f Handler = func() *int {
		i := 1
		return &i
	}
	fmt.Println(&f)

	var vh VarHandler = func(i ...int) {
		fmt.Println(i)
	}
	vh([]int{1, 2, 3}...)

	add2 := compose(func(i int) int {
		return i + 1
	}, func(i int) int {
		return i + 2
	})

	fmt.Println(add2(0))
}
