package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
	}
}

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit() error {
	fmt.Println("Existing Pokedex...")
	os.Exit(0)
	return nil
}
