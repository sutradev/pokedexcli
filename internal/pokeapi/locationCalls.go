package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) LocationCalls(pageUrl *string) (PokemonLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, ok := c.pokeCache.Get(url)
	if ok {
		pLocations := PokemonLocations{}
		err := json.Unmarshal(data, &pLocations)
		if err != nil {
			return PokemonLocations{}, err
		}
		fmt.Println("GOT FROM CACHE")

		return pLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocations{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonLocations{}, err
	}

	c.pokeCache.Add(url, data)

	pLocations := PokemonLocations{}
	err = json.Unmarshal(data, &pLocations)
	if err != nil {
		return PokemonLocations{}, err
	}

	return pLocations, nil
}
