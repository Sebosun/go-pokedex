package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var exploreURL = "/location-area"

func printExplore(data ExploreReturnType) {
	fmt.Println("Exploring ", data.Name, "...")

	fmt.Println("Found Pokemon: ")
	for _, key := range data.PokemonEncounters {
		fmt.Println(" - ", key.Pokemon.Name)
	}

}

func commandExplore(config *Config, input []string) error {
	if len(input) > 2 {
		return errors.New("Too many arguments!")
	}

	url := config.baseURL + exploreURL + "/" + input[1]

	res, ok := http.Get(url)
	if ok != nil {
		return errors.New("Invalid request")
	}
	defer res.Body.Close()

	val, getOk := config.cache.Get(url)
	if getOk {
		data := ExploreReturnType{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return errors.New("Unmarshaling failed")
		}
		printExplore(data)
		return nil
	}

	body, ok := io.ReadAll(res.Body)
	if ok != nil {
		return errors.New("Body parsing failed")
	}

	data := ExploreReturnType{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		return errors.New("Unmarshaling failed")
	}

	config.cache.Add(url, body)
	printExplore(data)
	return nil
}
