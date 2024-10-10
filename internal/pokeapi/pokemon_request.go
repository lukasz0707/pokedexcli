package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lukasz0707/pokedexcli/internal/utility"
)

func (c *Client) GetPokemon(pokemonName string) (any, error) {
	endpoint := "pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	cache, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		return cache, nil
	}
	fmt.Println("cache miss!")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return utility.Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return utility.Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return utility.Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data := utility.Pokemon{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return utility.Pokemon{}, err
	}
	c.cache.Add(fullURL, data)
	return data, nil
}
