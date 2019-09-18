package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

var SomeError = errors.New("error: the err")
var EOFError = errors.New("EOF")

type CommandError struct {
	err string
}

func (e CommandError) Error() string {
	return e.err
}

func Avoid() error {
	return &CommandError{"avoid command"}
}

func main() {
	if EOFError == io.EOF {
		log.Fatal("Should not be true")
	}

	fmt.Println(reflect.TypeOf(SomeError))

	switch SomeError.(type) {
	case error:
		fmt.Println("Yup it is an error")
	}
	switch SomeError {
	case SomeError:
		fmt.Println("Yup it is an error")
	}

	err := Avoid()
	if err, ok := err.(*CommandError); ok {
		fmt.Println(err)
	}
}
