package main

import (
	"fmt"
)

func commandMap(cfg *config, _ ...string) error {
	if cfg.mapStarted && cfg.url.next == nil {
		fmt.Println("You're already on the last page")
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.url.next)
	if err != nil {
		return err
	}
	cfg.mapStarted = true
	cfg.url.next = locations.Next
	cfg.url.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.url.previous == nil {
		fmt.Println("You're already on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.url.previous)
	if err != nil {
		return err
	}
	cfg.url.next = locations.Next
	cfg.url.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
