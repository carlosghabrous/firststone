package main

import (
	"fmt"
	"os"
)

const minimumArgumentsNumber int = 1

func main() {
	if len(os.Args[1:]) < minimumArgumentsNumber {
		fmt.Println("Cannot do anything without arguments!")
		help_cmd()
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "help" {
		help_cmd()

	} else if command == "init" {
		err := init_cmd(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {
		fmt.Println("unknown command")
		help_cmd()
		os.Exit(1)
	}

}

func help_cmd() {
	usage := "firststone <command> [flags] <project-name> <language>\n" +
		"Available commands:\n" +
		"init -> starts project\n" +
		"help -> prints this or a command's help\n"

	fmt.Println(usage)
}

func init_cmd(commandsAndFlags []string) error {
	projectName, language := commandsAndFlags[0], commandsAndFlags[1]

	if language == "python" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current working directory!")
			os.Exit(1)
		}

		err = initProject(cwd, language, projectName)
		if err != nil {
			fmt.Println("Error while copying files")
		}

	} else if language == "go" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current working directory!")
			os.Exit(1)
		}

		err = initProject(cwd, language, projectName)
		if err != nil {
			fmt.Println("Error while copying files")
		}

	} else {
		fmt.Println("Language not supported")
		os.Exit(1)
	}

	return nil
}

func initProject(dest, language, projectName string) error {

	return nil
}
