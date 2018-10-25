package main

import (
	"fmt"
	"os"

	"github.com/jlevesy/amnesia/repl"
)

const (
	msgExit = "Bye."
)

func main() {
	app := repl.New(os.Stdin, os.Stdout)

	err := app.Run()

	if err == repl.ErrExit {
		fmt.Println(msgExit)
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
