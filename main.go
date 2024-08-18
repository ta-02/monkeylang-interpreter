package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ta-02/c-like-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello %s! This is the Monkey Programming Language\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
