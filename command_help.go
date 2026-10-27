package main

import "fmt"

func commandHelp(cfg *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cfg.commands {
		fmt.Printf("%s: %v\n", command.name, command.description)
		if command.usage != "" {
			fmt.Println("	example usage: ", command.usage)
		}
		fmt.Println()
	}
	return nil
}
