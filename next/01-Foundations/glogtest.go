package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}
func init() {
	flag.Usage = usage
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	flag.Parse()
}

func main() {
	number_of_lines := 10
	for i := 0; i < number_of_lines; i++ {
		glog.V(2).Infof("LINE: %d", i)
		message := fmt.Sprintf("TEST LINE: %d", i)
		glog.Warning(message)
	}
	glog.Flush()
}
