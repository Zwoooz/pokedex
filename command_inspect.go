package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Printf("%s is not present in your pokedex, catch one first!\n", args[0])
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
	return nil
}
