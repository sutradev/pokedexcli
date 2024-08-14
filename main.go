package main

import (
	"time"

	pokeapi "github.com/sutradev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextPokemonLocationsURL *string
	prevPokemonLocationsURL *string
	pokemon                 map[string]*Pokemon
	pokemonSeen             map[string]struct{}
}

type Pokemon struct {
	name     string
	isCaught bool
	data     pokeapi.PokemonData
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemonSeen:   make(map[string]struct{}),
		pokemon:       make(map[string]*Pokemon),
	}
	startcli(cfg)
}
