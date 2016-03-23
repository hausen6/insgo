package command

import (
	"strings"
)

// InstallCommand command definition
type InstallCommand struct {
	Meta
}

// Run is a install command.
func (c *InstallCommand) Run(args []string) int {
	// Write your code here
	c.Ui.Info("[info] insgo start...")
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
