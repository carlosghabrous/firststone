package owncommands

import "github.com/carlosghabrous/firststone/firststonecmd"

var versionCmd = &firststonecmd.Command{
	Name:     "version",
	ShortDoc: "Outputs firststone's version",
	LongDoc:  "Firststone version 0.1",
	Usage:    "firststone version",
}
