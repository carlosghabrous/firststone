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

	command := os.Args[1]
	if command == "init" {
		initCmd(os.Args[2:]...)

	} else if command == "help" {
		helpCmd(os.Args[2:]...)

	} else if command == "clean" {
		cleanCmd(os.Args[2:]...)

	} else {
		fmt.Println("Unknown command")
		usage()
		os.Exit(1)
	}
}

func initCmd(commands ...string) {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initHelp := initCommand.Bool("help", false, "help")

	initCommand.Parse(commands)
	projectLanguage := unpackArgs(initCommand.Args()...)

	fmt.Printf("help %v, language %v\n", *initHelp, projectLanguage)

	if err := skeletons.CreateProject(projectLanguage); err != nil {
		fmt.Printf("Could not create project -> %v\n", err)
	}
}

func helpCmd(commands ...string) {
	usage()

}

func cleanCmd(commands ...string) {
	cleanCommand := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanHelp := cleanCommand.Bool("help", false, "help")
	cleanCommand.Parse(commands)

	projectLanguage := unpackArgs(cleanCommand.Args()...)
	fmt.Printf("help %v, language %v\n", *cleanHelp, projectLanguage)

	if err := skeletons.CleanProject(projectLanguage); err != nil {
		fmt.Printf("Could not clean project -> %v\n", err)
	}

}

func unpackArgs(args ...string) string {
	if len(args) < 1 {
		fmt.Println("Two arguments expected, project language!")
		os.Exit(1)
	}

	return args[0]
}
