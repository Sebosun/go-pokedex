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
			callback:    commandMapFowrards,
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
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show caught pokemons",
			callback:    commandPokedex,
		},
	}
}

func commandExit(_ *Config, _ []string) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func printPokedex(pokedex Pokedex) {
	fmt.Println("Your Pokedex: ")
	for k := range pokedex.entries {
		pokemon, ok := pokedex.entries[k]
		if ok {
			fmt.Println(" -", pokemon.Name)
		}
	}
}

func commandPokedex(config *Config, _ []string) error {
	if config.pokedex.Len() > 0 {
		printPokedex(config.pokedex)
		return nil

	}
	fmt.Println("Your pokedex is empty")
	return nil
}
