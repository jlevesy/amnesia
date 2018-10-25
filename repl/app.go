package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	prompt = "amnesia> "

	cmdExit = ".exit"

	errMsgUnknownCommand = "unknown command"
)

var (
	// ErrExit is raised when the REPL handles a .exit command
	// It represents a correct exit of the REPL loop.
	ErrExit = errors.New("exit")
)

// App is the representation of an REPL app
type App interface {
	Run() error
}

type app struct {
	out io.Writer
	in  *bufio.Reader
}

// New returns an instance of an App
func New(in io.Reader, out io.Writer) App {
	return &app{
		out: out,
		in:  bufio.NewReader(in),
	}
}

func (a *app) Run() error {
	for {
		if err := a.renderPrompt(); err != nil {
			return err
		}

		cmd, err := a.readLine()

		if err != nil {
			return err
		}

		if err := a.handleCommand(cmd); err != nil {
			return err
		}
	}

}

func (a *app) renderPrompt() error {
	_, err := a.out.Write([]byte(prompt))
	return err
}

func (a *app) readLine() (string, error) {
	raw, _, err := a.in.ReadLine()
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

func (a *app) handleCommand(cmd string) error {
	switch cmd {
	case cmdExit:
		return ErrExit
	default:
		fmt.Fprintln(os.Stdout, errMsgUnknownCommand)
	}

	return nil

}
