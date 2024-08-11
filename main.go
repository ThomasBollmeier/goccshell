package main

import (
	"errors"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"os/exec"
	"strings"
)

func main() {

	for {
		input, err := prompt()
		if err != nil {
			continue
		}
		result := handleInput(input)
		if result == exit {
			break
		}
	}

	os.Exit(0)
}

type handleResult int

const (
	exit handleResult = iota
	cont
)

func handleInput(input string) handleResult {

	if input == "exit" {
		return exit
	}

	parts := strings.Split(input, " ")
	command := parts[0]
	args := parts[1:]
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}

	return cont
}

func prompt() (string, error) {
	command, err := readline.Line("ccsh> ")
	if err != nil {
		return "", err
	}
	command = strings.TrimSpace(command)
	if command != "" {
		return command, nil
	} else {
		return "", errors.New("empty command")
	}

}
