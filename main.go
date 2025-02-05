package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	next     *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(ptr_config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(ptr_config *config) error {
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints all command and description",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints the next 20 locations",
			callback:    commandMap_forward,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous 20 locations",
			callback:    commandMap_backward,
		},
	}
}

func main() {
	current_config := config{}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			fmt.Println("Please put in an input")
			fmt.Println("[command] [arguments]")
		} else {
			if command, ok := getCommands()[cleanedInput[0]]; ok {
				command.callback(&current_config)
			} else {
				fmt.Printf("Unknown command: %s \n", cleanedInput[0])
			}
		}

		fmt.Print("Pokedex > ")
	}
}
