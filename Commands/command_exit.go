package commands

import (
	"fmt"
	"os"
)

func (c *commandManager) commandExit(ptr_config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
