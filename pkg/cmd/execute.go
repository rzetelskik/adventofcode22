package cmd

import (
	"fmt"
	"github.com/rzetelskik/adventofcode22/pkg/execute"
	"github.com/spf13/cobra"
)

const (
	pluginPathFlag      = "path"
	pluginPathFlagShort = "p"
)

type ExecuteOptions struct {
	path string
}

func NewExecuteOptions() *ExecuteOptions {
	return &ExecuteOptions{}
}

func NewExecuteCommand() *cobra.Command {
	so := NewExecuteOptions()

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a solution to the Advent of Code challenge",
		Long:  "Execute a solution to the Advent of Code challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error

			err = so.Validate()
			if err != nil {
				return err
			}

			err = so.Complete()
			if err != nil {
				return err
			}

			err = so.Run()
			if err != nil {
				return err
			}

			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Flags().StringVarP(&so.path, pluginPathFlag, pluginPathFlagShort, so.path, "Path to the the plugin")
	_ = cmd.MarkFlagRequired(pluginPathFlag)
	_ = cmd.MarkFlagDirname(pluginPathFlag)

	return cmd
}

func (so *ExecuteOptions) Validate() error {
	return nil
}

func (so *ExecuteOptions) Complete() error {
	return nil
}

func (so *ExecuteOptions) Run() error {
	g, err := execute.NewExecutor(so.path)
	if err != nil {
		return fmt.Errorf("can't create new executor: %w", err)
	}

	err = g.Run()
	if err != nil {
		return fmt.Errorf("can't run executor: %w", err)
	}

	return nil
}
