package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sutradev/pokedexcli/internal/pokeapi"
)

func listCommands(cfg *config) error {
	commandNames := makeCommands()
	fmt.Println("Welcome to Pokidex CLI!")
	fmt.Println("Here are a list of commands:")
	for _, command := range commandNames {
		fmt.Printf(" - %s: %s \n", command.name, command.description)
	}
	return nil
}

func exitPokidex(cfg *config) error {
	os.Exit(0)
	return nil
}

func next20(cfg *config) error {
	if cfg.nextPokemonLocationsURL != nil {
		cacheLocation, ok := cfg.pokeCache.Get(*cfg.nextPokemonLocationsURL)
		if ok {
			locations := pokeapi.PokemonLocations{}
			err := json.Unmarshal(cacheLocation, &locations)
			if err != nil {
				return err
			}
			fmt.Println("Here are some cached locations:")

			for _, location := range locations.Results {
				fmt.Println(" - " + location.Name)
			}

			cfg.pokeCache.Add(*cfg.nextPokemonLocationsURL, cacheLocation)
			cfg.nextPokemonLocationsURL = locations.Next
			cfg.prevPokemonLocationsURL = locations.Previous
			return nil
		}
		if !ok {
			locations, err := cfg.pokeapiClient.LocationCalls(cfg.nextPokemonLocationsURL)
			if err != nil {
				return err
			}
			fmt.Println("Here are some locations:")

			for _, location := range locations.Results {
				fmt.Println(" - " + location.Name)
			}

			byteLocations, err := json.Marshal(locations)
			if err != nil {
				return err
			}

			if cfg.nextPokemonLocationsURL != nil {
				cfg.pokeCache.Add(*cfg.nextPokemonLocationsURL, byteLocations)
			}
			cfg.nextPokemonLocationsURL = locations.Next
			cfg.prevPokemonLocationsURL = locations.Previous
			return nil
		}
	} else {
		locations, err := cfg.pokeapiClient.LocationCalls(cfg.nextPokemonLocationsURL)
		if err != nil {
			return err
		}
		fmt.Println("Here are some locations:")

		for _, location := range locations.Results {
			fmt.Println(" - " + location.Name)
		}

		byteLocations, err := json.Marshal(locations)
		if err != nil {
			return err
		}

		if cfg.nextPokemonLocationsURL != nil {
			cfg.pokeCache.Add(*cfg.nextPokemonLocationsURL, byteLocations)
		}
		cfg.nextPokemonLocationsURL = locations.Next
		cfg.prevPokemonLocationsURL = locations.Previous
		return nil
	}
	return nil
}

func prev20(cfg *config) error {
	if cfg.prevPokemonLocationsURL == nil {
		return errors.New("Can't go back any further!")
	}

	cacheLocation, ok := cfg.pokeCache.Get(*cfg.prevPokemonLocationsURL)
	if ok {
		locations := pokeapi.PokemonLocations{}
		err := json.Unmarshal(cacheLocation, &locations)
		if err != nil {
			return err
		}
		fmt.Println("Here are the previous cached locations:")

		for _, location := range locations.Results {
			fmt.Println(" - " + location.Name)
		}
		cfg.pokeCache.Add(*cfg.prevPokemonLocationsURL, cacheLocation)
		cfg.nextPokemonLocationsURL = locations.Next
		cfg.prevPokemonLocationsURL = locations.Previous
		return nil
	} else {

		locations, err := cfg.pokeapiClient.LocationCalls(cfg.prevPokemonLocationsURL)
		if err != nil {
			return err
		}

		fmt.Println("Here are the previous locations:")

		for _, location := range locations.Results {
			fmt.Println(" - " + location.Name)
		}
		byteLocations, err := json.Marshal(locations)
		if err != nil {
			return err
		}

		cfg.pokeCache.Add(*cfg.prevPokemonLocationsURL, byteLocations)
		cfg.nextPokemonLocationsURL = locations.Next
		cfg.prevPokemonLocationsURL = locations.Previous
		return nil
	}
}
