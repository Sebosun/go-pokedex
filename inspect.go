package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, input []string) error {
	if len(input) >= 3 {
		return errors.New("Too many arguments!")
	}

	val, ok := config.pokedex.Get(input[1])

	if !ok {
		errors.New("You haven't caught this pokemon or pokemon doesn't exist!")
	}

	fmt.Println("You have ", val.Name)
	return nil
}
