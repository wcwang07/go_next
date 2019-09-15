package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type intPtr *int
type chanType chan bool

//type rightChan chanType<-
type rightChan chan<- bool
type leftChan <-chan bool

type uPointer unsafe.Pointer

// type byte = uint8
// type rune = int32
// type []byte = []uint8

func main() {
	n := 1
	var pn intPtr = &n
	var up = uPointer(&pn)

	fmt.Println("->" + strconv.FormatBool(reflect.TypeOf(pn).Comparable())) // no
	fmt.Println(pn)
	fmt.Println(intPtr(up) == pn)
	fmt.Println(reflect.TypeOf(pn)) //intPtr
	fmt.Println(reflect.TypeOf(up))

	fmt.Println(reflect.TypeOf(nil)) // <nil>

	// Underlying types
	type IntSlice []int             // []int
	type CustomSlice IntSlice       // IntSlice -> []int
	type AnotherSlice = CustomSlice //alias
	type AnotherCustomSlice CustomSlice

	fmt.Println(reflect.TypeOf([]CustomSlice{}).Comparable())
	fmt.Println(reflect.TypeOf([]AnotherSlice{}))

	var is = IntSlice{}
	var cs = CustomSlice{}
	var as = AnotherSlice{}
	var acs = AnotherCustomSlice{}
	cs = as
	as = cs
	//cs = is
	//cs = acs not allowed
	cs = CustomSlice(is)
	acs = AnotherCustomSlice(cs)

	fmt.Println(acs)
}
