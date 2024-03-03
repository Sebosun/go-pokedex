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
