package main

import (
	"time"

	pokeapi "github.com/sutradev/pokedexcli/internal/pokeapi"
	"github.com/sutradev/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextPokemonLocationsURL *string
	prevPokemonLocationsURL *string
	pokeCache               pokecache.Cache
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(8 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}
	startcli(cfg)
}
