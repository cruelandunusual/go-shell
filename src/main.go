package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"log"
	"syscall"

	"golang.org/x/term"
)

// set global variable to hold prompt string
var globalPrompt = ""

// create command history list
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

		// read arrow keys
		captureInputBytes()
		// Read the keyboard input until newline reached
		input, err := reader.ReadString('\n')
		if err != nil {
			// fmt.Fprintln allows us to specify an output device, in this case Stderr
			// we could also import "log" and use log.Fatal(err)
			fmt.Fprintln(os.Stderr, err)
		}

		// add the command to the history list
		cmdHistory.addHistoryItem(&input)

		// handle input execution
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
			// remove the command from the history list if it returned an error
			cmdHistory.removeHistoryItem(&input) // TODO: this probably won't need &input passed to it; just remove the current item
		}
		// if err = cmdHistory.showAllHistory(); err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// }
	}
}



/********************************************************************************/

/* chatGPT code below this point */

/********************************************************************************/



func captureInputBytes() {
	// Set terminal to raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// Handle interrupt signals to restore terminal state
	// NOTE this is a goroutine
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	fmt.Println("Press arrow keys (up, down) or 'q' to quit:")

	for {
		var buf [1]byte
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Check for escape sequence
		if buf[0] == 27 { // Escape character (\e, or \033 in octal)
			_, err := os.Stdin.Read(buf[:])
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			if buf[0] == 91 { // '[' character
				_, err := os.Stdin.Read(buf[:])
				if err != nil {
					fmt.Println("Error reading input:", err)
					return
				}
				switch buf[0] {
				case 'A':
					fmt.Println("Up arrow pressed")
					// Handle showing the last command here
				case 'B':
					fmt.Println("Down arrow pressed")
					// Handle showing the next command here
				}
			}
		} else if buf[0] == 'q' {
			break // Exit on 'q'
		}
	}
}


