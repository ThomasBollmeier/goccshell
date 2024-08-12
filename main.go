package main

import (
	"errors"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"os/exec"
	"strings"
)

var curDir string

func main() {

	curDir, _ = os.Getwd()

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

	parts := strings.Fields(input)
	command := parts[0]
	args := parts[1:]

	if isBuiltIn(command) {
		result, err := handleBuiltin(command, args)
		if err != nil {
			panic(err)
		}
		return result
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}

	return cont
}

func handleBuiltin(command string, args []string) (handleResult, error) {
	switch command {
	case "exit":
		return exit, nil
	case "cd":
		var newDir string
		if len(args) > 0 {
			newDir = args[0]
		} else {
			newDir, _ = os.UserHomeDir()
		}
		err := os.Chdir(newDir)
		if err == nil {
			curDir, _ = os.Getwd()
		}
		return cont, nil
	case "pwd":
		fmt.Println(curDir)
		return cont, nil
	default:
		return -1, errors.New("unknown command: " + command)
	}

}

func isBuiltIn(command string) bool {
	switch command {
	case "exit", "cd", "pwd":
		return true
	default:
		return false
	}
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
