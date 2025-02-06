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

	command_manager := commands.NewCommandManager()

	//fmt.Println(command_manager.GetConfig())
	// command_manager.SetConfig("test", "b")
	// fmt.Println(command_manager.GetConfig())

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			fmt.Println("Please put in an input")
			fmt.Println("[command] [arguments]")
		} else {
			if c, ok := command_manager.GetCommands()[cleanedInput[0]]; ok {
				c.Callback(command_manager.GetConfig())
			} else {
				fmt.Printf("Unknown command: %s \n", cleanedInput[0])
			}
		}

		fmt.Print("Pokedex > ")
	}
}
