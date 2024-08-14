package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) EncounterCalls(location *string) (LocationEncounters, error) {
	if location == nil {
		return LocationEncounters{}, errors.New("No Location URL found")
	}

	url := baseURL + "/location-area/" + *location
	data, ok := c.pokeCache.Get(url)
	if ok {
		selectedLocation := LocationEncounters{}
		err := json.Unmarshal(data, &selectedLocation)
		if err != nil {
			return LocationEncounters{}, err
		}

		return selectedLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationEncounters{}, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return LocationEncounters{}, fmt.Errorf(
			"received status code: %s. Try a different location!",
			res.Status,
		)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationEncounters{}, err
	}

	c.pokeCache.Add(url, data)

	selectedLocation := LocationEncounters{}
	err = json.Unmarshal(data, &selectedLocation)
	if err != nil {
		return LocationEncounters{}, err
	}

	return selectedLocation, nil
}
