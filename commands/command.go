package commands

import "fmt"

type firstStoneCommand struct {
	cmdName string
}

type FirstStoneCommand interface {
	CmdName() string
	Run(cmdFlagsSubs ...string) error
}

func NewFirstStoneCmd(cmdName string) firstStoneCommand {
	return firstStoneCommand{cmdName: cmdName}
}

func (fsCmd firstStoneCommand) CmdName() string {
	return fsCmd.cmdName
}

func (fsCmd firstStoneCommand) Run(cmdFlagsSubs ...string) error {
	fmt.Printf("Running command %v with arguments %v\n", fsCmd.cmdName, cmdFlagsSubs)
	return nil
}

type initCommand firstStoneCommand

func NewInitCmd() initCommand {
	initCmd := NewFirstStoneCmd("init")
	return initCommand(initCmd)
}

func (initCmd initCommand) CmdName() string {
	return initCmd.cmdName
}

func (initCmd initCommand) Run(cmdFlagsSubs ...string) error {
	fmt.Printf("Running init command with args %v\n", cmdFlagsSubs)
	return nil
}

type helpCommand firstStoneCommand

func (helpCmd helpCommand) CmdName() string {
	return helpCmd.cmdName
}

func NewHelpCmd() helpCommand {
	helpCmd := NewFirstStoneCmd("help")
	return helpCommand(helpCmd)
}

func (helpCmd helpCommand) Run(cmdFlagsSubs ...string) error {
	fmt.Printf("Running help command with args %v\n", cmdFlagsSubs)
	return nil
}
