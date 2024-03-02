package main

import (
	"bufio"
	"fmt"
	"os"
)

var cliName = "cli"

func printPrompt() {
	fmt.Print(cliName, "> ")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func commandHelp() error {
	fmt.Println("List of commands: ")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		reader.Scan()

		text := reader.Text()

		commands := getCommands()
		val, ok := commands[text]
		if !ok {
			fmt.Println("Command not found! Use help to list available commands")
			continue
		}

		val.callback()
	}

}
