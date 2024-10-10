package main

import (
	"time"

	"github.com/lukasz0707/pokedexcli/internal/pokeapi"
	"github.com/lukasz0707/pokedexcli/internal/utility"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	caughtPokemon       map[string]utility.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute),
		caughtPokemon: make(map[string]utility.Pokemon),
	}
	startRepl(&cfg)
}
