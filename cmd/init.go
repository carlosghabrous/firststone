/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"

	"github.com/carlosghabrous/firststone/lang"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <app name> <language>",
	Short: "Creates the structure for a project",
	Long: `Creates files and folders to start the project. For example:

	firsttone init <app name> <language>
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Only one argument is required")
		}

		appLanguage := args[1]
		if err := lang.LanguageSupported(appLanguage); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		appName, appLanguage := args[0], args[1]

		//TODO: Check how to create a variable of ProjectBuilder type dynamically
		//TODO: general.go should have the data structures to get the appropriate builder from them
		var builder lang.ProjectBuilder

		switch appLanguage {
		case "python":
			builder = &lang.PythonProject{Name: appName, Language: appLanguage}

		case "golang":
			builder = &lang.GolangProject{Name: appName, Language: appLanguage}

		default:
			return errors.New("Unrecognized language! This should not have happened!")
		}

		if err := builder.CheckNamingConventions(); err != nil {
			return nil
		}

		builder.SetAppName(appName)

		return builder.Build()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
