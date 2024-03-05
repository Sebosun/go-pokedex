package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var pokemonURL string = "/pokemon/"

func catchPokemon(baseExp int) bool {
	chanceBase := int32(baseExp / 2)
	randomNum := rand.Int31n(int32(float64(baseExp) * 0.9))

	if randomNum > chanceBase {
		return true
	}
	return false

}

func printPokemon(config *Config, pokemon Pokemon) {
	fmt.Println("Attempting to catch...", pokemon.Name)
	isCaught := catchPokemon(pokemon.BaseExperience)
	if isCaught {
		fmt.Println("Congrats caught", pokemon.Name)
		config.pokedex.Add(pokemon.Name, pokemon)
		return
	}
	fmt.Println("Oops... the pokemon escaped")
	return
}

func commandCatch(config *Config, input []string) error {
	if len(input) >= 3 {
		return errors.New("Too many arguments!")
	}

	url := config.baseURL + pokemonURL + input[1]

	res, ok := http.Get(url)
	if ok != nil {
		return errors.New("Invalid request")
	}
	defer res.Body.Close()

	val, getOk := config.cache.Get(url)
	if getOk {
		data := Pokemon{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return errors.New("Unharshalin from catche GET failed")
		}
		printPokemon(config, data)
		return nil
	}

	body, ok := io.ReadAll(res.Body)
	if ok != nil {
		return errors.New("Body parsing failed")
	}

	data := Pokemon{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		return errors.New("Unmarshaling failed")
	}

	config.cache.Add(url, body)
	printPokemon(config, data)
	return nil
}
