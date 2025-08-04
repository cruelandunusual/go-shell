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
		// if the builtin has been handled then err is expected to be false.
		// However, we need to return err to satisfy execInput's own return value
		return err
	}

	// If the command entered is not a builtin then we proceed as though it's
	// a UNIX command, and prepare to pass it to the OS to be handled
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
		globalPrompt = createPrompt(setPromptMessage(args[1:]...))
		return true, nil
	case "exit", "quit":
		os.Exit(0)
	}
	return false, nil
}
