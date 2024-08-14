package main

type clicommands struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func makeCommands() map[string]clicommands {
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
			description: "Next page of 20 pokemon locations",
			callback:    next20,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous 20 pokemon locations",
			callback:    prev20,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore a specific location to view Pokemon in that area!",
			callback:    explore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Throw a pokeball and see if you can catch a pokemon!",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a pokemon you have caught to learn more!",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the pokemon you have caught!",
			callback:    pokedex,
		},
	}
}
