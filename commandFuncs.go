package main

import (
	"fmt"
	"os"
)

func listCommands() error {
	commandNames := makeCommands()
	fmt.Println("Welcome to Pokidex CLI!")
	fmt.Println("Here are a list of commands:")
	for _, command := range commandNames {
		fmt.Printf(" - %s: %s \n", command.name, command.description)
	}
	return nil
}

func exitPokidex() error {
	os.Exit(0)
	return nil
}
