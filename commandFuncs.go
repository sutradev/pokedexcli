package main

import (
	"errors"
	"fmt"
	"os"
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
	locations, err := cfg.pokeapiClient.LocationCalls(cfg.nextPokemonLocationsURL)
	if err != nil {
		return err
	}

	fmt.Println("Here are the Locations:")

	for _, location := range locations.Results {
		fmt.Println(" - " + location.Name)
	}

	cfg.nextPokemonLocationsURL = locations.Next
	cfg.prevPokemonLocationsURL = locations.Previous
	return nil
}

func prev20(cfg *config) error {
	if cfg.prevPokemonLocationsURL == nil {
		return errors.New("Can't go back any further!")
	}

	locations, err := cfg.pokeapiClient.LocationCalls(cfg.prevPokemonLocationsURL)
	if err != nil {
		return err
	}

	fmt.Println("Here are the previous locations:")

	for _, location := range locations.Results {
		fmt.Println(" - " + location.Name)
	}
	cfg.nextPokemonLocationsURL = locations.Next
	cfg.prevPokemonLocationsURL = locations.Previous
	return nil
}
