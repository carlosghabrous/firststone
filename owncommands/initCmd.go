package owncommands

import "github.com/carlosghabrous/firststone/firststonecmd"

var initCmd = &firststonecmd.Command{
	Name:     "init",
	ShortDoc: "Creates a new project",
	LongDoc:  "Creates the project template for a project in a specific language",
	Usage:    `firststone init LANGUAGE`,
}
