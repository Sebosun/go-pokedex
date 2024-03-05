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
		return errors.New("You haven't caught this pokemon or pokemon doesn't exist!")
	}

	fmt.Println("Name:", val.Name)
	fmt.Println("Height:", val.Height)
	fmt.Println("Height:", val.Weight)

	for _, stat := range val.Stats {
		fmt.Println("  -"+stat.Stat.Name+": ", stat.BaseStat)
	}

	return nil
}
