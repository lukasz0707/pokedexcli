package main

import (
	"errors"
	"fmt"

	"github.com/lukasz0707/pokedexcli/internal/utility"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	t, ok := resp.(utility.LocationAreaResp)
	if !ok {
		return fmt.Errorf("wrong type of response %T", resp)
	}
	fmt.Println("Location areas:")
	for _, area := range t.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = t.Next
	cfg.prevLocationAreaUrl = t.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}
	t, ok := resp.(utility.LocationAreaResp)
	if !ok {
		return fmt.Errorf("wrong type of response %T", resp)
	}
	fmt.Println("Location areas:")
	for _, area := range t.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = t.Next
	cfg.prevLocationAreaUrl = t.Previous
	return nil
}
