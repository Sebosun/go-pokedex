package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandPrintList,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Go forward in a map",
			callback:    commandMapForwards,
		},
		"mapb": {
			name:        "mapb",
			description: "Go backwards in a map",
			callback:    commandMapBackwards,
		},
		"explore": {
			name:        "explore",
			description: "Explore a map",
			callback:    commandExplore,
		},
	}
}

func commandExit(_ *Config, _ []string) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
