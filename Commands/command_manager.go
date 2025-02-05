package commands

type Config struct {
	next     *string
	previous *string
}

type CliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints all command and description",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints the next 20 locations",
			Callback:    commandMap_forward,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous 20 locations",
			Callback:    commandMap_backward,
		},
	}
}
