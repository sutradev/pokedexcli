package main

import (
	"time"

	pokeapi "github.com/sutradev/pokidexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextPokemonLocationsURL *string
	prevPokemonLocationsURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startcli(cfg)
}
