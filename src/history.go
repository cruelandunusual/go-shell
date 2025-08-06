package main

import "fmt"

type historyItem struct {
	command string
	next    *historyItem
	prev    *historyItem
}

type history struct {
	name       string
	head       *historyItem
	currentCmd *historyItem
}

func createHistory(name string) *history {
	return &history{
		name: name,
	}
}

func (h *history) addHistoryItem(command string) error {
	c := &historyItem{
		command: command,
	}
	if h.head == nil {
		h.head = c
	} else {
		currentNode := h.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = c
	}
	return nil
}

func (h *history) showAllHistory() error {
	currentNode := h.head
	if currentNode == nil {
		fmt.Println("History is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (h *history) nextItem() *historyItem {
	h.currentCmd = h.currentCmd.next
	return h.currentCmd
}

func (h *history) prevItem() *historyItem {
	h.currentCmd = h.currentCmd.prev
	return h.currentCmd
}
