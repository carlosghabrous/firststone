package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/carlosghabrous/firststone/skeletons"
)

const usageDoc string = "firststone <command> [flags]"

func usage() {
	fmt.Println(usageDoc)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command found!")
		usage()
		os.Exit(1)
	}

	if err := dispatchCommand(os.Args[1]); err != nil {
		fmt.Printf("Error: %v\n", err)
		usage()
		os.Exit(1)
	}
}

func dispatchCommand(command string) (err error) {

	if command == "init" {
		err = initCmd(os.Args[2:]...)

	} else if command == "help" {
		err = helpCmd(os.Args[2:]...)

	} else if command == "clean" {
		err = cleanCmd(os.Args[2:]...)

	} else {
		err = fmt.Errorf("Unknown command")
	}

	return
}

func initCmd(commands ...string) (err error) {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initHelp := initCommand.Bool("help", false, "help")

	initCommand.Parse(commands)
	projectLanguage, err := unpackArgs(initCommand.Args()...)
	if err != nil {
		return fmt.Errorf("Wrong input commands %v\n", err)
	}

	if *initHelp {
		usageInitCmd()
		return
	}

	if err := skeletons.CreateProject(projectLanguage); err != nil {
		return fmt.Errorf("Could not create project -> %v\n", err)
	}

	return
}

func helpCmd(commands ...string) (err error) {
	usage()
	return
}

func cleanCmd(commands ...string) (err error) {
	cleanCommand := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanHelp := cleanCommand.Bool("help", false, "help")
	cleanCommand.Parse(commands)

	projectLanguage, err := unpackArgs(cleanCommand.Args()...)
	if err != nil {
		return fmt.Errorf("Wrong input commands %v\n", err)
	}

	if *cleanHelp {
		usageHelpCmd()
		return
	}

	if err := skeletons.CleanProject(projectLanguage); err != nil {
		return fmt.Errorf("Could not clean project -> %v\n", err)
	}

	return
}

func unpackArgs(args ...string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("Two arguments expected, project language!")
	}

	return args[0], nil
}

func usageInitCmd() {
	fmt.Println("Print init command's usage")
}

func usageHelpCmd() {
	fmt.Println("Print help command's usage")
}
