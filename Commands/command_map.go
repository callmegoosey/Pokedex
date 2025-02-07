package commands

import (
	"errors"
	"fmt"

	pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"
)

// ===========================================
//
//	Check the pokecache
//	if exist, reuse
//	else creates a new cache
//
// ===========================================

func (c *commandManager) get(url string) (body []byte, err error) {
	cached_body, success := c.cache.Get(url)

	if !success {
		return_result, err := pokeapi.Get(url)
		if err != nil {
			return nil, err
		}

		if string(return_result) == "Not Found" {
			return nil, errors.New("not found")
		}

		cached_body = return_result
		c.cache.Add(url, cached_body)
		fmt.Println("==========================")
		fmt.Println("created cache")
		fmt.Println("==========================")
	} else {

		fmt.Println("==========================")
		fmt.Println("used cache")
		fmt.Println("==========================")
	}
	//fmt.Printf("%s\n", cached_body)
	return cached_body, nil
}

// ===========================================
//
//	Base code for commandMap to reduce duplicate code
//
// ===========================================
func (c *commandManager) commandMap_base(forward bool, ptr_config *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if forward {
		if ptr_config.next != nil {
			url = *ptr_config.next
		}
	} else {
		if ptr_config.previous != nil {
			url = *ptr_config.previous
		}
	}

	cached_body, err := c.get(url)

	if err != nil {
		return err
	}

	next, prev, err := pokeapi.Get_map_names(&cached_body)

	if err != nil {
		return err
	}

	ptr_config.next = &next
	ptr_config.previous = &prev

	return nil
}

// ===========================================
//
//	Returns next 20 map name
//
// ===========================================
func (c *commandManager) commandMap_forward(ptr_config *config, other_inputs *[]string) error {
	return c.commandMap_base(true, ptr_config)
}

// ===========================================
//
//	Returns previous 20 map name
//
// ===========================================
func (c *commandManager) commandMap_backward(ptr_config *config, other_inputs *[]string) error {
	return c.commandMap_base(false, ptr_config)
}

func (c *commandManager) commandMap_explore(ptr_config *config, other_inputs *[]string) error {
	if len(*other_inputs) == 0 {
		return fmt.Errorf("missing other arguments")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + (*other_inputs)[0] + "/"

	cached_body, err := c.get(url)

	if err != nil {
		return err
	}

	return pokeapi.Explore_map(&cached_body, other_inputs)
}
