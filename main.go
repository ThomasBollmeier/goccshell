package main

import (
	"github.com/chzyer/readline"
	"os"
	"os/exec"
	"strings"
)

func main() {
	input, err := prompt()
	if err != nil {
		panic(err)
	}
	handleInput(input)
}

func handleInput(input string) bool {
	parts := strings.Split(input, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err == nil
}

func prompt() (string, error) {
	command, err := readline.Line("ccsh> ")
	if err != nil {
		return "", err
	}
	command = strings.TrimRight(command, " \n")
	return command, nil
}
