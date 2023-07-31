package main

import (
	"fmt"
	"os"
	"os/user"
	"duskterpreter/repl"

)

func main () {
	user, err := user.Current()

	if (err != nil) {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Dusk language!\n", user.Username)
	fmt.Printf("Feel free to test out the commands :)\n")

	repl.Start(os.Stdin, os.Stdout)
}