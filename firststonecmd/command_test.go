package firststonecmd

import (
	"fmt"
	"testing"
)

func TestAddSingleCommand(t *testing.T) {

	var testCmd = &Command{
		Name:     "main",
		ShortDoc: "Some short documentation string",
		LongDoc:  "Some long documentation string",
		Usage:    "The command is used like this",
	}

	if err := testCmd.Add(&Command{Name: "main subcommand"}); err != nil {
		t.Errorf("add(&command) returned %v, expected nil\n", err)
	}
}

func TestAddMultipleCommands(t *testing.T) {

	var testCmd = &Command{
		Name:     "main",
		ShortDoc: "Some short documentation string",
		LongDoc:  "Some long documentation string",
		Usage:    "The command is used like this",
	}

	numberOfSubcommands := 3

	for i := 0; i < numberOfSubcommands; i++ {
		subCmdName := fmt.Sprintf("main subcommand %d", i)
		testCmd.Add(&Command{Name: subCmdName})
	}

	lenSubs := len(testCmd.subcommands)
	if lenSubs != numberOfSubcommands {
		t.Errorf("expected number of subcommands: %d, got: %d\n", numberOfSubcommands, lenSubs)
	}
}

func TestAddExistingSubCommand(t *testing.T) {

	var testCmd = &Command{
		Name:     "main",
		ShortDoc: "Some short documentation string",
		LongDoc:  "Some long documentation string",
		Usage:    "The command is used like this",
	}

	testCmd.Add(&Command{Name: "subcommand#1"})
	if err := testCmd.Add(&Command{Name: "subcommand#1"}); err == nil {
		t.Errorf("expected error when adding duplicated command, got nil\n")
	}
}
