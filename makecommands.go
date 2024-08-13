package main

type clicommands struct {
	name        string
	description string
	callback    func() error
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
	}
}
