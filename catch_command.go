package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := cfg.pokeapiClient.PokemonInfo(args[0])
	if err != nil {
		return err
	}

	random := rand.Intn(pokemon.BaseExperience)

	if random <= 50 {
		fmt.Printf("%s was caught!\n", args[0])
		fmt.Println("You may now inspect it with the inspect command")
		cfg.pokedex[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}

	return nil
}
