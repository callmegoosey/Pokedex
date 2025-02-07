package commands

import (
	"time"

	pokeapi "github.com/callmegoosey/Pokedex/Internal/Pokeapi"
	pokecache "github.com/callmegoosey/Pokedex/Internal/Pokecache"
)

type config struct {
	next     *string
	previous *string
}

type CliCommand struct {
	name        string
	description string
	Callback    func(*config, *[]string) error
}

type commandManager struct {
	config  config
	cache   pokecache.Cache
	pokedex map[string]pokeapi.PokedexData
}

func NewCommandManager() commandManager {
	return commandManager{
		config:  config{},
		cache:   pokecache.NewCache(30 * time.Second),
		pokedex: make(map[string]pokeapi.PokedexData),
	}
}

func (c *commandManager) GetConfig() *config {
	return &c.config
}

func (c *commandManager) SetConfig(next string, prev string) {
	c.config.next = &next
	c.config.previous = &prev
}

func (c *commandManager) GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			Callback:    c.commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints all command and description",
			Callback:    c.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints the next 20 locations",
			Callback:    c.commandMap_forward,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous 20 locations",
			Callback:    c.commandMap_backward,
		},
		"explore": {
			name:        "explore map_name",
			description: "Explores a map and print all pokemon in the map",
			Callback:    c.commandMap_explore,
		},
		"catch": {
			name:        "catch pokemon_name",
			description: "catches a pokemon",
			Callback:    c.commandMap_catch,
		},
		"pokedex": {
			name:        "pokedex [optional: pokemon_name]",
			description: "prints all the name of all caught. if pokemon_name is provided, prints stats of pokedex instead",
			Callback:    c.commandMap_pokedex,
		},
	}
}
