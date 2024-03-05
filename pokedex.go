package main

import (
	"sync"
)

type Pokedex struct {
	entries map[string]Pokemon
	mu      *sync.Mutex
}

func (c *Pokedex) Add(key string, pokemon Pokemon) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.entries[key]

	if ok {
		return
	}

	c.entries[key] = pokemon
}

func (c *Pokedex) Get(key string) (Pokemon, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.entries[key]
	if !ok {
		return Pokemon{}, false
	}

	return v, true
}

func (c *Pokedex) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	count := 0
	for range c.entries {
		count++
	}

	return count
}

func constructPokedex() Pokedex {
	pokedex := Pokedex{
		entries: make(map[string]Pokemon),
		mu:      &sync.Mutex{},
	}

	return pokedex
}
