// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

var _ cli.Command = (*ConfigCommand)(nil)

type ConfigCommand struct {
	*BaseCommand
}

func (c *ConfigCommand) Synopsis() string {
	return "Interact with config contexts"
}

func (c *ConfigCommand) Help() string {
	helpText := `
Usage: bao config <subcommand> [options] [args]

  This command groups subcommands for interacting with Vault configs.
  These subcommands operate on the contexts of Vault servers

  Describe one or many contexts:

      $ bao config list

  Set a context entry in bao config:

      $ bao config create vault-1 -server https://1.2.3.4 -namespace xyz

  Set the current-context in bao config:

      $ bao config use vault-1

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (c *ConfigCommand) Run(args []string) int {
	return cli.RunResultHelp
}
