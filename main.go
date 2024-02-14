package main

import (
	"fmt"
	"mylang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!, Please type in commands\n", user)

	repl.Start(os.Stdin, os.Stdout)
}
