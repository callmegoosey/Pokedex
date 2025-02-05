package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/callmegoosey/Pokedex/Commands"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	current_config := commands.Config{}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			fmt.Println("Please put in an input")
			fmt.Println("[command] [arguments]")
		} else {
			if c, ok := commands.GetCommands()[cleanedInput[0]]; ok {
				c.Callback(&current_config)
			} else {
				fmt.Printf("Unknown command: %s \n", cleanedInput[0])
			}
		}

		fmt.Print("Pokedex > ")
	}
}
