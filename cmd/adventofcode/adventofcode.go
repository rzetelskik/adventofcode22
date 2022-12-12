package main

import (
	"fmt"
	"github.com/rzetelskik/adventofcode22/pkg/cmd"
	"os"
)

func main() {
	command := cmd.NewAdventOfCodeCommand()
	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
