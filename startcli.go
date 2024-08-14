package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startcli(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commandNames := makeCommands()
	fmt.Println("pokedex > Hi!")
	fmt.Println("pokedex > Welcome to your Pokedex!")
	fmt.Println("pokedex > Here you can explore the land of pokemon!")
	fmt.Println("pokedex > You can try and catch pokemon you find while exploring!")
	fmt.Println("pokedex > Then use me to learn about the pokemon you caught!")
	fmt.Println("pokedex > Just like a real Pokedex!")
	fmt.Println("pokedex > To get started type 'help' to get a list of commands!")
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		if input == "" {
			continue
		}

		slicedInput := stringToSlice(input)

		mainCommand := slicedInput[0]

		command, ok := commandNames[mainCommand]
		if !ok {
			fmt.Println("Command Not Found, Use 'help' For Command List")
			continue
		}
		switch command.name {
		case "explore":
			if err := command.callback(cfg, slicedInput); err != nil {
				fmt.Printf("Error when executing command: '%v' Error: %v \n", input, err)
			}
		default:
			if err := command.callback(cfg, slicedInput); err != nil {
				fmt.Printf("Error when executing command: '%v' Error: %v \n", input, err)
			}
		}
	}
}

func stringToSlice(s string) []string {
	slicedStrings := strings.Fields(s)
	return slicedStrings
}
