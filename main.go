package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var cliName = "Pokedex"

type Config struct {
	cache mapCache
	prev  string
	next  string
}

type ModifConfig interface {
	UpdatePrev(string)
	UpdateNext(string)
}

func (c *Config) UpdateNext(next string) {
	c.next = next
}

func (c *Config) UpdatePrev(prev string) {
	c.prev = prev
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	config := Config{
		cache: constructCache(time.Second * 10),
		prev:  "",
		next:  "https://pokeapi.co/api/v2/location-area",
	}

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

		err := val.callback(&config)
		if err != nil {
			fmt.Println(err)
		}
	}

}
