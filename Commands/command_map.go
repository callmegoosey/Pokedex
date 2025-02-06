package commands

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"
)

func (c *commandManager) commandMap_forward(ptr_config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.next != nil {
		url = *ptr_config.next
	}

	result, success := c.cache.Get(url)

	if success {
		fmt.Println("==========================")
		fmt.Println("used cache")
		fmt.Println("==========================")
	} else {
		return_result, err := pokeapi.CommandMap(url)
		if err != nil {
			return err
		}
		result = return_result
		fmt.Println("==========================")
		fmt.Println("created cache")
		fmt.Println("==========================")
	}

	poke_location := pokeapi.PokeLocationAreas{}

	if err := json.Unmarshal(result, &poke_location); err != nil {
		return err
	}

	if !success {
		c.cache.Add(url, result)
	}

	ptr_config.next = &poke_location.Next
	ptr_config.previous = &poke_location.Previous

	//print all 20 location
	for _, location := range poke_location.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func (c *commandManager) commandMap_backward(ptr_config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.previous != nil {
		url = *ptr_config.previous
	}

	result, success := c.cache.Get(url)

	if success {
		fmt.Println("==========================")
		fmt.Println("used cache")
		fmt.Println("==========================")
	} else {
		return_result, err := pokeapi.CommandMap(url)
		if err != nil {
			return err
		}
		result = return_result
		fmt.Println("==========================")
		fmt.Println("created cache")
		fmt.Println("==========================")
	}

	poke_location := pokeapi.PokeLocationAreas{}

	if err := json.Unmarshal(result, &poke_location); err != nil {
		return err
	}

	if !success {
		c.cache.Add(url, result)
	}

	ptr_config.next = &poke_location.Next
	ptr_config.previous = &poke_location.Previous

	//print all 20 location
	for _, location := range poke_location.Results {
		fmt.Println(location.Name)
	}

	return nil
}
