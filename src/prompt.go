package main

import (
	"fmt"
	"os"
	"strings"
)

// set a global variable to the user's home directory
var home, err = os.UserHomeDir()

func createPrompt(message string) string {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var promptSpacer = ": "
	var promptMsg = ""
	if message == "" {
		promptMsg = getDefaultPrompt()
	} else {
		promptMsg = message
	}
	var promptStub = ": > "
	var promptComplete = ""
	promptComplete = hostName + promptSpacer + promptMsg + promptStub
	return promptComplete
}

// Variadic function: takes a variable number of input strings
func setPromptMessage(input ...string) string {
	promptString := ""
	// ignore the returned index from range, we only want the value at that index
	for _, i := range input {
		promptString += i + " "
	}
	if promptString == "" {
		return getDefaultPrompt()
	}
	promptString = strings.TrimSuffix(promptString, " ")
	return promptString
}

func getHomeDir() string {
	return home
}

func getDefaultPrompt() string {
	return "go-shell"
}
