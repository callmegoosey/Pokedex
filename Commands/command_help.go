package commands

import "fmt"

func (c *commandManager) commandHelp(ptr_config *config) error {
	for _, command := range c.GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
