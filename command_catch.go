package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/lukasz0707/pokedexcli/internal/utility"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	t, ok := resp.(utility.Pokemon)
	if !ok {
		return fmt.Errorf("wrong type of response %T", t)
	}

	const threshold = 50
	randNum := rand.IntN(t.BaseExperience)
	fmt.Println(t.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = t
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
