package main

import (
	"os"
	"fmt"
)

var supportedCommands map[string]cliCommand

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
    fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range supportedCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func initCommands() {
	supportedCommands = map[string]cliCommand {
    	"exit": {
        	name:        "exit",
        	description: "Exit the Pokedex",
        	callback:    commandExit,
    	},
		"help": {
        	name:        "help",
        	description: "Displays a help message",
        	callback:    commandHelp,
    	},
	}
}
