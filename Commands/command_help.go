package commands

import "fmt"

func commandHelp(ptr_config *Config) error {
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
