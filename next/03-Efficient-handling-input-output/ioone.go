package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	fmt.Println(os.Args)

	_, err := os.Stat("i-dont-exist")
	fmt.Printf("%v\n", os.IsNotExist(err))

	currDir, err := os.Getwd()
	fmt.Printf("Current Dir: %s\n", currDir)

	fmt.Printf("Current $PATH: %s\n", os.Getenv("PATH"))

	currUser, err := user.Current()
	if err != nil {
		panic("Cannot find user")
	}
	fmt.Printf("Current USer: %s\n", currUser)

	cmd := exec.Command("ls", "-ltr")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s \n", stdout)

	defer fmt.Println("Call to Exit..")
	os.Exit(1)
	//log.Fatal("exit")
}
