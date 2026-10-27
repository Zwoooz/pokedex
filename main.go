package main

import (
	"time"

	"github.com/zwoooz/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		commands:      initCommands(),
		pokeapiClient: *pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
