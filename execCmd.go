package main

import (
	"os"
	"os/exec"
	"strings"
)

// Handles standard shell commands (ls, cat, etc).
func execInput(input string) error {
	// strip the newline character from the input string
	input = strings.TrimSuffix(input, "\n")

	// create an array of any arguments passed
	args := strings.Split(input, " ")

	// check if args[0] is a shell builtin
	if handled, err := execBuiltin(args); handled {
		return err
	}

	// prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	// set appropriate outputs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run the command, returning its results and exit status
	return cmd.Run()
}

// Handles shell builtins. Returns true if a builtin was executed.
func execBuiltin(args []string) (bool, error) {
	switch args[0] {
	case "cd", "chdir", "set-location", "Set-Location", "loc":
		if len(args) == 1 {
			return true, os.Chdir(getHomeDir())
		}
		return true, os.Chdir(args[1])
	case "setPrompt":
		globalPrompt, err = createPrompt(setPromptMessage(args[1:]...))
		return true, err
	case "exit", "quit":
		os.Exit(0)
	}
	return false, nil
}
