package commands

import pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"

func commandMap_forward(ptr_config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.next != nil {
		url = *ptr_config.next
	}
	next, prev, err := pokeapi.CommandMap(url)

	ptr_config.next = &next
	ptr_config.previous = &prev

	return err
}

func commandMap_backward(ptr_config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if ptr_config.previous != nil {
		url = *ptr_config.previous
	}

	next, prev, err := pokeapi.CommandMap(url)

	ptr_config.next = &next
	ptr_config.previous = &prev

	return err
}
