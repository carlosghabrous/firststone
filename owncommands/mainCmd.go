package owncommands

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/firststone/firststonecmd"
)

var mainCmd = &firststonecmd.Command{
	Name:     "firsttone",
	ShortDoc: "Firsttone automates projects layout's creation",
	LongDoc: `Firststone creates standard project templates for projects written in the
	supported languages`,
	Usage: "firststone COMMAND ARGUMENTS",
}

func init() {
	mainCmd.Add(versionCmd)
	mainCmd.Add(initCmd)
}

func Execute() {
	if err := mainCmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
