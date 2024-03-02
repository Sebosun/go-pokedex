package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Go forward in a map",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Go backwards in a map",
			callback:    mapBackwards,
		},
	}
}

func commandHelp(config *Config) error {
	fmt.Println("List of commands: ")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandExit(config *Config) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
