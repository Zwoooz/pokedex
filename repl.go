package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := cfg.commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}
		args := input[1:]
		if len(args) != command.arguments {
			fmt.Println("Usage: ", command.usage)
			fmt.Println()
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	strArr := strings.Fields(lower)

	return strArr
}
