package cmd

import (
	"github.com/spf13/cobra"
)

type AdventOfCodeOptions struct {
	day string
}

func NewAdventOfCodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aocctl",
		Short: "Run the Advent of Code 2022 CLI",
		Long:  "Run the Advent of Code 2022 CLI",
	}

	cmd.AddCommand(NewGenerateCommand())
	cmd.AddCommand(NewExecuteCommand())

	return cmd
}
