package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	location, err := cfg.pokeapiClient.LocationInfo(args[0])
	if err != nil {
		return err
	}

	for _, encounters := range location.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil
}
