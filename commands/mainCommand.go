package commands

import (
	"fmt"
	"os"
)

var mainCommand = commands.Command()

func Run() {
	if err := mainCommand.Run(); err != nil {
		fmt.Printf("Error while running the command: %v\n", err)
		os.Exit(1)
	}
}
