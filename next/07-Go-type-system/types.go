package main

func main() {
	type IntSlice []int
	type CustomSlice IntSlice
	type AnotherSlice = CustomSlice
	type YetAnotherCustomSlice CustomSlice

	var is = IntSlice{}
	var cs = CustomSlice{}
	var as = AnotherSlice{}
	//	var yeas = YetAnotherCustomSlice{}

	cs = as
	as = cs

	cs = CustomSlice(is)

}
