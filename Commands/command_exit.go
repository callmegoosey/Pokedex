package commands

import (
	"fmt"
	"os"
)

func commandExit(ptr_config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
