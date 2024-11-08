package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/ayahiro1729/go_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Print("Feel free to type in commands")
  repl.Start(os.Stdin, os.Stdout)
}
