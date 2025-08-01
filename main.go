package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// set a global variable to the user's home directory
var home, errHome = os.UserHomeDir()
var hostName, errHost = os.Hostname()
var promptHost = hostName
var promptMsg = "go-shell"
var promptStub = ": > "
var promptSpacer = ": "
var promptComplete = ""

func createPrompt() {
	promptComplete = promptHost + promptSpacer + promptMsg + promptStub
	//promptComplete = promptMsg + promptStub
}

func main() {
	// create an instance of bufio.NewReader to capture input
	reader := bufio.NewReader(os.Stdin)
	// use an infinite loop to capture input until we Ctrl-C or enter `exit` or `quit`
	for {
		createPrompt()
		fmt.Print(promptComplete)

		// Read the keyboard input until newline reached
		input, err := reader.ReadString('\n')

		// fmt.Fprintln allows us to specify an output device, in this case Stderr
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle input execution
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

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
		if len(args) < 2 {
			return true, os.Chdir(home)
		}
		return true, os.Chdir(args[1])
	case "setPrompt":
		return true, setPromptMessage(args[1:]...) // variadic argument expansion
	case "exit", "quit":
		os.Exit(0)
	}
	return false, nil
}

// Variadic function: takes a variable number of input strings
func setPromptMessage(input ...string) (err error) {
	promptString := ""
	// ignore the returned index from range, we only want the value at that index
	for _, i := range input {
		promptString += i + " "
	}
	if promptString == "" {
		promptMsg = "go-shell"
		return
	}
	// set the global promptMsg var to the promptString we've built
	promptMsg = strings.TrimSuffix(promptString, " ")
	return
}
