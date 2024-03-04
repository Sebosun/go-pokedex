package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	directionNext = "next"
	directionPrev = "prev"
)

func updateMap(config *Config, pokeMap PokeApiMap) {
	config.UpdateNext(pokeMap.Next)

	switch v := pokeMap.Previous.(type) {
	case string:
		config.UpdatePrev(v)
	default:
		config.UpdatePrev("")
	}
}

func commandMapFowrards(config *Config, input []string) error {
	pokeMap, err := handleMapRequest(*config, directionNext)

	if err != nil {
		return err
	}

	printMaps(pokeMap)
	updateMap(config, pokeMap)

	return nil
}

func commandMapBackwards(config *Config, input []string) error {
	pokeMap, err := handleMapRequest(*config, directionPrev)

	if err != nil {
		return err
	}

	printMaps(pokeMap)
	updateMap(config, pokeMap)
	return nil
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

	val, getOk := config.cache.Get(url)
	if getOk {
		data := PokeApiMap{}
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokeApiMap{}, errors.New("Unmarshaling failed")
		}
		return data, nil
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

	config.cache.Add(url, body)
	return data, nil
}
