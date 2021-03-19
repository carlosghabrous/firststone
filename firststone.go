package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const minimumArgumentsNumber int = 1

func main() {
	if len(os.Args[1:]) < minimumArgumentsNumber {
		fmt.Println("Cannot do anything with that!")
		fmt.Println("Print help/usage or whatever")
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
		fmt.Println("Print help/usage or whatever")
		os.Exit(1)
	}

}

func help_cmd() {
	usage := "firststone <command> [flags] <project-name> <language>"
	usage += "Available commands:"
	usage += "init -> starts project"
	usage += "help -> prints this or a command's help"
	fmt.Println(usage)
}

func init_cmd(commandsAndFlags []string) error {
	fmt.Println("Execute init command with arguments", commandsAndFlags)
	projectName, language := commandsAndFlags[0], commandsAndFlags[1]
	defaultPermissions := 0755

	err := os.Mkdir(projectName, os.FileMode(defaultPermissions))
	if err != nil {
		fmt.Println("Error while creating directory ", projectName)
		return err
	}

	if language == "python" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current working directory!")
			os.Exit(1)
		}

		err = recursiveCopy(cwd, filepath.Join("_languages", language))
		if err != nil {
			fmt.Println("Error while copying files")
		}

	} else if language == "go" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current working directory!")
			os.Exit(1)
		}

		err = recursiveCopy(cwd, filepath.Join("_languages", language))
		if err != nil {
			fmt.Println("Error while copying files")
		}

	} else {
		fmt.Println("Language not supported")
		os.Exit(1)
	}

	return nil
}

func recursiveCopy(dest, src string) error {
	fmt.Println("reading directory ", src)
	files, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println("Something very wrong happened!")
	}

	for _, entry := range files {
		fmt.Println(entry.Name())
	}
	return nil
}
