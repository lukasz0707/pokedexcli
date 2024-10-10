package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokemons := cfg.caughtPokemon
	if len(pokemons) == 0 {
		fmt.Println("No pokemons catched yet")
		return nil
	}
	fmt.Println("Pokemons in pokedex:")
	for pokemonName := range pokemons {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
