package main

import (
	"fmt"
	"reflect"
	"time"
)

type UserType reflect.Type
type User struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User: %v, %v", u.FirstName, u.LastName)
}

func main() {

	alex := User{}
	userType := reflect.TypeOf(alex)
	fmt.Println(userType.NumField())
	fmt.Println(userType.Comparable())
	fmt.Println(userType.Kind())
	fmt.Println(userType.NumMethod())
	fmt.Println(userType.MethodByName("String"))
	//	g, _ := userType.MethodByName("String")

	intSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 0)
	fmt.Println(intSlice)
	intSlice = reflect.Append(intSlice, reflect.ValueOf(1))
	fmt.Println(intSlice)

}
