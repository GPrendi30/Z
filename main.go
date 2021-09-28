package main

import (
	"Z/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Z programming language!\n",
		user.Username)
	fmt.Println("Feel free to type commands.")
	repl.Start(os.Stdin, os.Stdout)
}
