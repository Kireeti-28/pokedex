package main

import "github.com/kireeti-28/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.GetClient(),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}


	startRepl(&cfg)
}
