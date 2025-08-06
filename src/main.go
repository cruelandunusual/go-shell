package main

import (
	"bufio"
	"fmt"
	"os"
)

// set global variable to hold prompt string
var globalPrompt = ""
var myHistory = "cmdHistory"
var cmdHistory = createHistory(myHistory)

func main() {
	// create an instance of bufio.NewReader to capture input
	reader := bufio.NewReader(os.Stdin)
	globalPrompt = createPrompt("")

	run(*reader) // pass the bufio.Reader to run()
}

func run(reader bufio.Reader) error {
	// use an infinite loop to capture input until we Ctrl-C or enter `exit` or `quit`
	for {
		fmt.Print(globalPrompt)

		// Read the keyboard input until newline reached
		input, err := reader.ReadString('\n')
		cmdHistory.addHistoryItem(input)
		if err != nil {
			// fmt.Fprintln allows us to specify an output device, in this case Stderr
			// we could also import "log" and use log.Fatal(err)
			fmt.Fprintln(os.Stderr, err)
		}

		// handle input execution
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
