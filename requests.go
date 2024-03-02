package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	directionNext = "next"
	directionPrev = "prev"
)

type PokeApiMap struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func printMaps(pokeMaps PokeApiMap) {
	for _, key := range pokeMaps.Results {
		fmt.Println(key.Name)
	}
}

func mapForward(config *Config) error {
	conf := *config
	pokeMap, err := handleMapRequest(conf, directionNext)

	if err != nil {
		return err
	}

	printMaps(pokeMap)
	conf.UpdateNext(pokeMap.Next)

	switch v := pokeMap.Previous.(type) {
	case string:
		conf.UpdatePrev(v)
	default:
		conf.UpdatePrev("")
	}

	*config = conf

	return nil
}

func mapBackwards(config *Config) error {
	conf := *config
	pokeMap, err := handleMapRequest(*config, directionPrev)

	if err != nil {
		return err
	}

	printMaps(pokeMap)
	conf.UpdateNext(pokeMap.Next)
	switch v := pokeMap.Previous.(type) {
	case string:
		conf.UpdatePrev(v)
	default:
		conf.UpdatePrev("")
	}

	*config = conf

	return err
}

func handleMapRequest(config Config, requestType string) (PokeApiMap, error) {
	var url string
	if requestType == directionNext {
		url = config.next
	} else {
		url = config.prev
	}

	if url == "" {
		return PokeApiMap{}, errors.New("Invalid URL, Prev or Next missing")
	}

	res, ok := http.Get(url)
	if ok != nil {
		return PokeApiMap{}, errors.New("Invalid request")
	}
	defer res.Body.Close()

	body, ok := io.ReadAll(res.Body)
	if ok != nil {
		return PokeApiMap{}, errors.New("Body parsing failed")
	}

	data := PokeApiMap{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		return PokeApiMap{}, errors.New("Unmarshaling failed")
	}

	return data, nil
}
