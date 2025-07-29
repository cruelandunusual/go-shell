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
var promptMsg = "Enter command"
var promptStub = ": > "
var promptSpacer = " - "
var promptComplete = ""

func buildPrompt() {
	promptComplete = promptHost + promptSpacer + promptMsg + promptStub
}

func main() {

	// create an instance of bufio.NewReader to capture input
	// technically the reader variable is a pointer to a bufio.Reader struct
	reader := bufio.NewReader(os.Stdin)
	// use an infinite loop to capture input until we Ctrl-C
	for {
		// create a prompt
		buildPrompt()
		fmt.Print(promptComplete)

		// Read the keyboard input
		// a newline stops input reading
		// and the input variable stores the input since the last newline
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

func execInput(input string) error {

	buildPrompt()
	// strip the newline character from the input string
	input = strings.TrimSuffix(input, "\n")

	// create an arry of any arguments passed
	// args[0] will be the command itself, args[1] will be the first argument, etc
	args := strings.Split(input, " ")

	// check if args[0] is a shell builtin, e.g. cd
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir(home)
		}
		return os.Chdir(args[1])

	case "setPrompt":
		return setPromptMessage(args[1:]...)

	case "exit":
		os.Exit(0)
	}

	// prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...) // neat little syntax for executing all the args, given that we don't know how many will be passed

	// set appropriate outputs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run the command, returning its results and exit status
	return cmd.Run()
}

// Variadic function: takes a variable number of input strings
func setPromptMessage(input ...string) (err error) {
	promptString := ""
	// the underscore below signifies we're ignoring the returned index from range and only using the value at that index (stored in i)
	for _, i := range input {
		promptString += i + " "
	}
	if promptString == "" {
		err = fmt.Errorf("Error: empty prompt string")
		return
	}
	// set the global promptMsg var to the promptString we've built
	promptMsg = strings.TrimSuffix(promptString, " ")
	return
}
