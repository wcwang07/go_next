package main

import (
	"fmt"
	"unsafe"
)

type MySlice struct {
	elems unsafe.Pointer
	len   int
	cap   int
}

const e = 2

func main() {
	var mapping = make(map[int]int)
	mapping[e] = e

	value, ok := mapping[e]
	fmt.Println(value, ok)

	var array2 = [...]int{2: 1, 3: 2, 4: 3}
	fmt.Println(array2)

	var x = array2[2:3]
	fmt.Println(x)
	array2[2] = 9

	fmt.Println(x)

	fmt.Println(len(array2), cap(array2))
	fmt.Println(len(mapping))
	delete(mapping, e)
	fmt.Println(len(mapping))

}
