package main

import (
	"fmt"
)

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func printMaps(pokeMaps PokeApiMap) {
	for _, key := range pokeMaps.Results {
		fmt.Println(key.Name)
	}
}

func commandPrintList(_ *Config, _ []string) error {
	fmt.Println("List of commands: ")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
