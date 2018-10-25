package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
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

	readErr chan error
	readCmd chan string
	sigs    chan os.Signal
}

// New returns an instance of an App
func New(in io.Reader, out io.Writer) App {
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	return &app{
		out: out,
		in:  bufio.NewReader(in),

		readErr: make(chan error),
		readCmd: make(chan string),
		sigs:    sigChan,
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
	go func() {
		raw, _, err := a.in.ReadLine()
		if err != nil {
			a.readErr <- err
		}
		a.readCmd <- string(raw)
	}()

	select {
	case cmd := <-a.readCmd:
		return cmd, nil
	case err := <-a.readErr:
		return "", err
	case sig := <-a.sigs:
		return "", fmt.Errorf("received a signal %v, exiting", sig)
	}
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
