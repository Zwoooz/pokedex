package main

import "github.com/zwoooz/bootdev/pokedex/internal/pokeapi"

type url struct {
	next     *string
	previous *string
}
type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	url           url
	mapStarted    bool
	pokedex       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	arguments   int
	usage       string
	callback    func(cfg *config, args ...string) error
}

func initCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			usage:       "'help'",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			usage:       "'exit'",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists next page of 20 pokemon areas",
			usage:       "'map'",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of 20 pokemon areas",
			usage:       "'mapb'",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon in provided area, takes area as argument",
			arguments:   1,
			usage:       "'explore <location>'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the provided pokemon, takes pokemon as argument",
			arguments:   1,
			usage:       "'catch <pokemon>'",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex, tales pokemon as argument",
			arguments:   1,
			usage:       "'inspect <pokemon>'",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught pokemon",
			usage:       "'pokedex'",
			callback:    commandPokedex,
		},
	}
}
