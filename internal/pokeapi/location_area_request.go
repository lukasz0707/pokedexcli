package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lukasz0707/pokedexcli/internal/utility"
)

func (c *Client) ListLocationAreas(pageUrl *string) (utility.LocationAreaResp, error) {
	endpoint := "location-area/"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	cache, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		return cache, nil
	}
	fmt.Println("cache miss!")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return utility.LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return utility.LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return utility.LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data := utility.LocationAreaResp{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return utility.LocationAreaResp{}, err
	}
	c.cache.Add(fullURL, data)
	return data, nil
}
