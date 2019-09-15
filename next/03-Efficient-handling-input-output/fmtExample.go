package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("%v (%d €)", p.Name, p.Price)
}

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	// This is wa
	flag.Parse()
}

func main() {
	// Fmt relies on the Stringer interface to print custom types
	s := fmt.Sprint(Product{"Quirky Pants", 100})
	// build-in function
	glog.V(2).Info(s)
}
