package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonCall(pokemon string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + pokemon
	data, ok := c.pokeCache.Get(url)
	if ok {
		pokemonData := PokemonData{}
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return PokemonData{}, err
		}

		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return PokemonData{}, fmt.Errorf(
			"received status code: %s. Try a different pokemon!",
			res.Status,
		)
	}

	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	c.pokeCache.Add(url, data)

	pokemonData := PokemonData{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonData, nil
}
