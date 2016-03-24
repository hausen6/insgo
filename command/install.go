package command

import (
	"flag"
	"fmt"
	"strings"
)

// InstallCommand command definition
type InstallCommand struct {
	Meta
}

// Run is a install command.
func (c *InstallCommand) Run(args []string) int {
	// check argument
	parser := flag.NewFlagSet("install", flag.ExitOnError)
	parser.Usage = func() {
		c.Help()
	}
	if err := parser.Parse(args); err != nil {
		c.Ui.Error(fmt.Sprintf("[Error]: %v", err))
		return 2
	}
	// need toml file
	if parser.NArg() == 0 {
		c.Ui.Error("[Error] a few argument. install list file must be required.")
		return 2
	}

	return 0
}

// Synopsis is a simple command description.
func (c *InstallCommand) Synopsis() string {
	return "Install sume programs"
}

// Help is a install commands help.
func (c *InstallCommand) Help() string {
	helpText := `
This is a Install commands from file.
Usage:
	$ insgo install install.toml
`
	return strings.TrimSpace(helpText)
}
