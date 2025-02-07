package commands

import (
	"fmt"

	pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"
)

func (c *commandManager) commandMap_catch(ptr_config *config, other_inputs *[]string) error {
	if len(*other_inputs) == 0 {
		return fmt.Errorf("missing other arguments")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + (*other_inputs)[0] + "/"

	cached_body, err := c.get(url)

	if err != nil {
		return err
	}
	caught, p_data, err := pokeapi.Catch(&cached_body, other_inputs)

	if err != nil {
		return err
	}

	if caught {
		if _, ok := c.pokedex[p_data.Name]; !ok {
			fmt.Printf("%s has been added to pokedex\n", p_data.Name)
			c.pokedex[p_data.Name] = p_data
		}
	}

	return nil
}
