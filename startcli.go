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
	for {
		fmt.Print("pokidex > ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		if input == "" {
			continue
		}

		command, ok := commandNames[input]
		if !ok {
			fmt.Println("Command Not Found, Use 'help' For Command List")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Printf("Error when executing command: '%v' Error: %v \n", input, err)
		}
	}
}
