package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CommandExecutor interface {
	next20() error
	prev20() error
}

type PokemonUrls struct {
	Next *string
	Prev *string
}

type PokemonLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func listCommands() error {
	commandNames := makeCommands()
	fmt.Println("Welcome to Pokidex CLI!")
	fmt.Println("Here are a list of commands:")
	for _, command := range commandNames {
		fmt.Printf(" - %s: %s \n", command.name, command.description)
	}
	return nil
}

func exitPokidex() error {
	os.Exit(0)
	return nil
}

func (p PokemonUrls) next20() error {
	if *p.Next == "" {
		*p.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(*p.Next)
	if err != nil {
		return errors.New("Error getting pokemon locations")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error reading response body")
	}

	var locations PokemonLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return errors.New("Error unmarshalling JSON response")
	}

	*p.Next = locations.Next
	if prev, ok := locations.Previous.(string); ok {
		*p.Prev = prev
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func (p PokemonUrls) prev20() error {
	if *p.Prev == "" {
		return errors.New("Can't go back any further!")
	}
	res, err := http.Get(*p.Prev)
	if err != nil {
		return errors.New("Error getting pokemon locations")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error reading response body")
	}

	var locations PokemonLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return errors.New("Error unmarshalling JSON response")
	}

	if prev, ok := locations.Previous.(string); ok {
		*p.Prev = prev
	}
	*p.Next = locations.Next
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
