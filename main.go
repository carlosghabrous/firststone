package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/carlosghabrous/firststone/languages"
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
	initOverwrite := initCommand.Bool("overwrite", false, "overwrite")
	initGitStart := initCommand.Bool("gitstart", false, "gitstart")

	initCommand.Parse(commands)
	projectName, projectLanguage := unpackArgs(initCommand.Args()...)

	fmt.Printf("help %v, overwrite %v, gitstart %v, name %v, language %v\n", *initHelp, *initOverwrite, *initGitStart, projectName, projectLanguage)

	if !languages.IsSupportedLanguage(projectLanguage) {
		fmt.Printf("Language %v is not supported\n", projectLanguage)
		os.Exit(1)
	}

	languages.CreateProject(projectName, projectLanguage)
}

func helpCmd(commands ...string) {
	usage()

}

func cleanCmd(commands ...string) {
	cleanCommand := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanCommand.Parse(commands)

	projectName, projectLanguage := unpackArgs(cleanCommand.Args()...)
	fmt.Printf("name %v, language %v\n", projectName, projectLanguage)

	if !languages.IsSupportedLanguage(projectLanguage) {
		fmt.Printf("Language %v is not supported\n", projectLanguage)
		os.Exit(1)
	}

}

func unpackArgs(args ...string) (string, string) {
	if len(args) < 2 {
		fmt.Println("Two arguments expected, project name and language!")
		os.Exit(1)
	}

	return args[0], args[1]
}
