package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	buff := make([]byte, 1024)
	f, err := os.Open("/etc/passwd")
	check(err)
	defer f.Close()
	for {
		n, e := f.Read(buff)
		if n == 0 {
			break
		}
		check(e)
		os.Stdout.Write(buff[:n])
	}
}
