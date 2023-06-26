package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
			description: "List some location areas backward",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in location areas",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View information about caught pokemon",
			callback:    callbackInspect,
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

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("Existing Pokedex...")
	time.Sleep(time.Second)
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config, args ...string) error {
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

func callbackMapb(cfg *config, args ...string) error {
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

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No/Incorrect location area provided")
	}

	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s \n", locationAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No/Incorrect pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randInt := rand.Intn(pokemon.BaseExperience)
	if randInt > threshold {
		return fmt.Errorf("failed to caught %s\n", pokemonName)
	}

	cfg.pokemonCaught[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No/Incorrect pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokemonCaught[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}
