package cmd

import (
	"fmt"
	"github.com/rzetelskik/adventofcode22/pkg/generate"
	"github.com/spf13/cobra"
)

const (
	pathFlag      = "path"
	pathFlagShort = "p"
	dayFlag       = "day"
	dayFlagShort  = "d"
)

type GenerateOptions struct {
	path string
	day  int
}

func NewGenerateOptions() *GenerateOptions {
	return &GenerateOptions{}
}

func NewGenerateCommand() *cobra.Command {
	so := NewGenerateOptions()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate scaffolding for an Advent of Code challenge",
		Long:  "Generate scaffolding for an Advent of Code challenge",
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

	cmd.Flags().IntVarP(&so.day, dayFlag, dayFlagShort, so.day, "Day of the challenge")
	cmd.Flags().StringVarP(&so.path, pathFlag, pathFlagShort, so.path, "Path to the directory that the scaffolding will be generated at")
	_ = cmd.MarkFlagRequired(dayFlag)
	_ = cmd.MarkFlagRequired(pathFlag)
	_ = cmd.MarkFlagDirname(pathFlag)

	return cmd
}

func (so *GenerateOptions) Validate() error {
	if so.day < 0 || so.day > 25 {
		return fmt.Errorf("invalid day: day has to be a number in the range [0,24]")
	}

	return nil
}

func (so *GenerateOptions) Complete() error {
	return nil
}

func (so *GenerateOptions) Run() error {
	g, err := generate.NewGenerator(so.path, so.day)
	if err != nil {
		return fmt.Errorf("can't create new generator: %w", err)
	}

	err = g.Run()
	if err != nil {
		return fmt.Errorf("can't run generator: %w", err)
	}

	return nil
}
