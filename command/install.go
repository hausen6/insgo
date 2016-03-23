package command

import (
	"strings"
)

type InstallCommand struct {
	Meta
}

func (c *InstallCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *InstallCommand) Synopsis() string {
	return "Install sume programs"
}

func (c *InstallCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
