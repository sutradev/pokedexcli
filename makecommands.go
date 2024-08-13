package main

type clicommands struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
