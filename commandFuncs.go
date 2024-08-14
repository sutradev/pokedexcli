package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var unknown = errors.New("Command Not Found, Use 'help' For Command List")

func listCommands(cfg *config, s []string) error {
	if len(s) > 1 {
		return unknown
	}
	commandNames := makeCommands()
	fmt.Println("Here are a list of commands:")
	for _, command := range commandNames {
		fmt.Printf(" - %s: %s \n", command.name, command.description)
	}
	return nil
}

func exitPokidex(cfg *config, s []string) error {
	if len(s) > 1 {
		return unknown
	}
	os.Exit(0)
	return nil
}

func next20(cfg *config, s []string) error {
	if len(s) > 1 {
		return unknown
	}
	locations, err := cfg.pokeapiClient.LocationCalls(cfg.nextPokemonLocationsURL)
	if err != nil {
		return err
	}
	fmt.Println("Here are some locations:")

	for _, location := range locations.Results {
		fmt.Println(" - " + location.Name)
	}

	cfg.nextPokemonLocationsURL = locations.Next
	cfg.prevPokemonLocationsURL = locations.Previous
	return nil
}

func prev20(cfg *config, s []string) error {
	if len(s) > 1 {
		return unknown
	}
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

func explore(cfg *config, location []string) error {
	if len(location) <= 1 {
		return errors.New("Please give a location to explore!")
	}
	if len(location) > 2 {
		return errors.New("Please give one location at a time!")
	}

	encounters, err := cfg.pokeapiClient.EncounterCalls(&location[1])
	if err != nil {
		return err
	}
	pokemonEncounters := encounters.PokemonEncounters
	fmt.Printf("Welcome to %v!\n", encounters.Location.Name)
	fmt.Println("Pokemon in the area:")
	for _, pokemon := range pokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
		cfg.pokemonSeen[pokemon.Pokemon.Name] = struct{}{}
	}
	return nil
}

func catch(cfg *config, pokemon []string) error {
	if len(pokemon) <= 1 {
		return errors.New("Please name a pokemon you are trying to catch!")
	}
	if len(pokemon) > 2 {
		return errors.New("You can only catch one pokemon at a time!")
	}
	_, ok := cfg.pokemonSeen[pokemon[1]]
	if !ok {
		return errors.New("You haven't seen this pokemon yet, explore more areas!")
	}

	pokemonData, err := cfg.pokeapiClient.PokemonCall(pokemon[1])
	if err != nil {
		return err
	}
	cfg.pokemon[pokemonData.Name] = &Pokemon{
		name:     pokemonData.Name,
		isCaught: false,
		data:     pokemonData,
	}

	xpOnCatch := pokemonData.BaseExperience
	caughtStatus := catchResults(xpOnCatch)
	if caughtStatus {
		fmt.Println("You caught the pokemon!")
		cfg.pokemon[pokemonData.Name].isCaught = true
	} else {
		fmt.Println("Aw so close! Try again!")
	}
	return nil
}

func catchResults(difficulty int) bool {
	// Seed the random number generator to ensure different results each time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Normalize difficulty to a range, e.g., 1 to 1000, with a higher difficulty meaning a harder catch
	normalizedDifficulty := difficulty
	if normalizedDifficulty < 1 {
		normalizedDifficulty = 1 // Ensure minimum difficulty
	}

	// Calculate the base probability (lower difficulty means higher probability)
	// For example, a lower difficulty results in a higher chance to catch
	baseProbability := 1000 - normalizedDifficulty

	// Introduce randomness by generating a random number within a certain range
	randomFactor := r.Intn(1000) // Range 0 to 999

	// Determine the threshold for catching the Pokémon
	// Here, if the random factor is less than or equal to the base probability, the catch succeeds
	if randomFactor <= baseProbability {
		return true
	}

	// If the random factor exceeds the base probability, the catch fails
	return false
}

func (cfg *config) inspect(pokemon []string) error {
	if len(pokemon) <= 1 {
		return errors.New("Please name a pokemon you are trying to catch!")
	}
	if len(pokemon) > 2 {
		return errors.New("You can only inspect one pokemon at a time!")
	}
	caughtPokemon, ok := cfg.pokemon[pokemon[1]]
	if !ok {
		return errors.New("You haven't seen this pokemon yet, explore more areas!")
	}

	if !caughtPokemon.isCaught {
		return errors.New("You haven't caught this pokemon yet! Try now!")
	}
}
