package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List some location areas",
			callback:    callbackMapb,
		},

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

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit(cfg *config) error {
	fmt.Println("Existing Pokedex...")
	time.Sleep(time.Second)
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("No Map Back")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
