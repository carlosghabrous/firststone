package main

import (
	"fmt"
	"os"
)

const minimumArgumentsNumber int = 2

type FirstStoneCommand struct {
	cmdName string
}

func (fsCmd FirstStoneCommand) run(cmdFlagsSubs ...string) {
	fmt.Printf("Running command %v\n", fsCmd.cmdName)
}

var commandsRegistry = make(map[string]FirstStoneCommand)

func init() {
	commandsRegistry["init"] = FirstStoneCommand{cmdName: "init"}
	commandsRegistry["help"] = FirstStoneCommand{cmdName: "help"}
}

func main() {

	helpCommand := commandsRegistry["help"]

	if len(os.Args) < minimumArgumentsNumber {
		fmt.Println("Missing command!")
		helpCommand.run()
		os.Exit(1)
	}

	commandName := os.Args[1]
	command, ok := commandsRegistry[commandName]

	if !ok {
		fmt.Printf("Unknown command %v\n", commandName)
		helpCommand.run()
		os.Exit(1)
	}

	command.run(os.Args[2:]...)
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
