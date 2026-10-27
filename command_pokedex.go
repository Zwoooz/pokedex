package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet so the pokedex is empty\nGet our there!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	fmt.Printf("You have caught a total of %v pokemon", len(cfg.pokedex))
	return nil
}
