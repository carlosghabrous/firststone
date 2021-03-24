package main

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/firststone/commands"
)

const minimumArgumentsNumber int = 2

var commandsRegistry = make(map[string]commands.FirstStoneCommand)

func init() {
	var c commands.FirstStoneCommand

	c = commands.NewInitCmd()
	commandsRegistry[c.CmdName()] = c

	c = commands.NewHelpCmd()
	commandsRegistry[c.CmdName()] = c
}

func main() {

	helpCommand := commandsRegistry["help"]

	if len(os.Args) < minimumArgumentsNumber {
		fmt.Println("Missing command!")
		helpCommand.Run()
		os.Exit(1)
	}

	commandName := os.Args[1]
	command, ok := commandsRegistry[commandName]

	if !ok {
		fmt.Printf("Unknown command %v\n", commandName)
		helpCommand.Run()
		os.Exit(1)
	}

	err := command.Run(os.Args[2:]...)
	if err != nil {
		fmt.Printf("Error while running command %v: %v\n", command.CmdName(), err)
	}
}

// command := os.Args[1]
// if command == "help" {
// 	help_cmd()

// } else if command == "init" {
// 	err := init_cmd(os.Args[2:])
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// } else {
// 	fmt.Println("unknown command")
// 	help_cmd()
// 	os.Exit(1)
// }

// func help_cmd() {
// 	usage := "firststone <command> [flags] <project-name> <language>\n" +
// 		"Available commands:\n" +
// 		"init -> starts project\n" +
// 		"help -> prints this or a command's help\n"

// 	fmt.Println(usage)
// }

// func init_cmd(commandsAndFlags []string) error {
// 	projectName, language := commandsAndFlags[0], commandsAndFlags[1]

// 	if language == "python" {
// 		cwd, err := os.Getwd()
// 		if err != nil {
// 			fmt.Println("Could not get current working directory!")
// 			os.Exit(1)
// 		}

// 		err = initProject(cwd, language, projectName)
// 		if err != nil {
// 			fmt.Println("Error while copying files")
// 		}

// 	} else if language == "go" {
// 		cwd, err := os.Getwd()
// 		if err != nil {
// 			fmt.Println("Could not get current working directory!")
// 			os.Exit(1)
// 		}

// 		err = initProject(cwd, language, projectName)
// 		if err != nil {
// 			fmt.Println("Error while copying files")
// 		}

// 	} else {
// 		fmt.Println("Language not supported")
// 		os.Exit(1)
// 	}

// 	return nil
// }
