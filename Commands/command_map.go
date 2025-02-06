package commands

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"
)

func (c *commandManager) commandMap_forward(ptr_config *config, other_inputs *[]string) error {
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
		return_result, err := pokeapi.Get(url)
		if err != nil {
			return err
		}
		result = return_result
		fmt.Println("==========================")
		fmt.Println("created cache")
		fmt.Println("==========================")
	}

	poke_location_areas := pokeapi.PokeLocationAreas{}

	if err := json.Unmarshal(result, &poke_location_areas); err != nil {
		return err
	}

	if !success {
		c.cache.Add(url, result)
	}

	ptr_config.next = &poke_location_areas.Next
	ptr_config.previous = &poke_location_areas.Previous

	//print all 20 location
	for _, location := range poke_location_areas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func (c *commandManager) commandMap_backward(ptr_config *config, other_inputs *[]string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.previous != nil {
		url = *ptr_config.previous
	}

	result, success := c.cache.Get(url)

	if !success {
		return_result, err := pokeapi.Get(url)
		if err != nil {
			return err
		}
		result = return_result
	}
	// 	fmt.Println("==========================")
	// 	fmt.Println("created cache")
	// 	fmt.Println("==========================")
	// } else {

	// 	fmt.Println("==========================")
	// 	fmt.Println("used cache")
	// 	fmt.Println("==========================")
	// }

	poke_location_areas := pokeapi.PokeLocationAreas{}

	if err := json.Unmarshal(result, &poke_location_areas); err != nil {
		return err
	}

	if !success {
		c.cache.Add(url, result)
	}

	ptr_config.next = &poke_location_areas.Next
	ptr_config.previous = &poke_location_areas.Previous

	//print all 20 location
	for _, location := range poke_location_areas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func (c *commandManager) commandMap_explore(ptr_config *config, other_inputs *[]string) error {
	if len(*other_inputs) == 0 {
		return fmt.Errorf("missing other arguments")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + (*other_inputs)[0] + "/"

	return_result, err := pokeapi.Get(url)
	if err != nil {
		return err
	}

	poke_location_areas := pokeapi.PokeLocation{}

	if err := json.Unmarshal(return_result, &poke_location_areas); err != nil {
		return err
	}

	//fmt.Printf("%s", return_result)

	fmt.Printf("Exploring %s...\n", (*other_inputs)[0])
	fmt.Printf("Found Pokemon:\n")

	for _, pokemon := range poke_location_areas.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
