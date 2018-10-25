package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	prompt = "amnesia> "

	cmdExit = ".exit"

	errMsgUnknownCommand = "unknown command"
)

func renderPrompt(out io.Writer) error {
	_, err := out.Write([]byte(prompt))

	return err
}

func readLine(in *bufio.Reader) (string, error) {
	raw, _, err := in.ReadLine()
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func main() {
	bufStdin := bufio.NewReader(os.Stdin)

	for {
		if err := renderPrompt(os.Stdout); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cmd, err := readLine(bufStdin)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch cmd {
		case cmdExit:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stdout, errMsgUnknownCommand)
		}
	}
}
