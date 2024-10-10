package main

import (
	"errors"
	"fmt"

	"github.com/lukasz0707/pokedexcli/internal/utility"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	t, ok := resp.(utility.LocationArea)
	if !ok {
		return fmt.Errorf("wrong type of response %T", t)
	}

	fmt.Printf("Pokemon in %s:\n", t.Name)
	for _, pokemon := range t.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
