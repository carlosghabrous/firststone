package firststonecmd

import "fmt"

type Command struct {
	Name        string
	ShortDoc    string
	LongDoc     string
	Usage       string
	Run         func() error
	subcommands map[string]*Command
}

// add adds a subcommand to another one
func (cmd *Command) Add(command *Command) error {
	if cmd.subcommands == nil {
		cmd.subcommands = make(map[string]*Command)
	}

	if _, ok := cmd.subcommands[command.Name]; ok {
		return fmt.Errorf("Command %v already registered under %v\n", command.Name, cmd.Name)
	}

	cmd.subcommands[command.Name] = command
	return nil
}
