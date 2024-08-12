package main

type clicommands struct {
	name        string
	description string
	callback    func() error
}

func makeCommands() map[string]clicommands {
	pokemonUrls := &PokemonUrls{
		Next: new(string),
		Prev: new(string),
	}
	return map[string]clicommands{
		"help": {
			name:        "help",
			description: "Gives user list of commands",
			callback:    listCommands,
		},
		"exit": {
			name:        "exit",
			description: "Exits program",
			callback:    exitPokidex,
		},
		"map": {
			name:        "map",
			description: "Shows next 20 pokemon locations",
			callback:    pokemonUrls.next20,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back to the previous 20 pokemon locations",
			callback:    pokemonUrls.prev20,
		},
	}
}
