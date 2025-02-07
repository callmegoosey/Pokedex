package commands

import (
	"fmt"
)

func (c *commandManager) commandMap_pokedex(ptr_config *config, other_inputs *[]string) error {
	if len(*other_inputs) == 0 {
		if len(c.pokedex) == 0 {
			fmt.Println("You have not caught any pokemon")
		} else {
			fmt.Println("Current pokedex: ")

			for _, pokemon := range c.pokedex {
				fmt.Printf(" - %s\n", pokemon.Name)
			}
		}
		return nil
	} else {
		result, success := c.pokedex[(*other_inputs)[0]]
		if !success {
			fmt.Printf("%s not in pokedex\n", (*other_inputs)[0])
		} else {
			fmt.Printf("Name: %s\n", result.Name)
			fmt.Printf("Height: %d\n", result.Height)
			fmt.Printf("Weight: %d\n", result.Weight)

			fmt.Println("Stats:")
			for index, stat := range result.Stats {
				fmt.Printf("	- %s: %d\n", index, stat)
			}

			fmt.Println("Types:")
			for _, t_type := range result.Types {
				fmt.Printf("	- %s\n", t_type)
			}
		}
	}
	return nil
}
