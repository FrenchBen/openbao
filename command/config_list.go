// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/openbao/openbao/api"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*ConfigListCommand)(nil)
	_ cli.CommandAutocomplete = (*ConfigListCommand)(nil)
)

type ConfigListCommand struct {
	*BaseCommand
}

func (c *ConfigListCommand) Synopsis() string {
	return "List contexts available"
}

func (c *ConfigListCommand) Help() string {
	helpText := `
Usage: bao config list [options]

  Lists all contexts in a bao config.

  List all contexts in a bao config:

      $ bao config list

` + c.Flags().Help()

	return strings.TrimSpace(helpText)
}

func (c *ConfigListCommand) Flags() *FlagSets {
	set := c.flagSet(FlagSetNone)

	f := set.NewFlagSet("Command Options")

	f.BoolVar(&BoolVar{
		Name:    "detailed",
		Target:  &c.flagDetailed,
		Default: false,
		Usage:   "Print detailed information such as namespace ID.",
	})

	return set
}

func (c *ConfigListCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ConfigListCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *ConfigListCommand) Run(args []string) int {
	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	args = f.Args()
	if len(args) > 0 {
		c.UI.Error(fmt.Sprintf("Too many arguments (expected 0, got %d)", len(args)))
		return 1
	}

	contexts, defaultContext, err := api.ListContexts()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error listing contexts: %s", err))
		return 2
	}

	if Format(c.UI) != "table" {
		if contexts == nil {
			OutputData(c.UI, map[string]interface{}{})
			return 2
		}
	}

	if contexts == nil {
		c.UI.Error("No contexts found")
		return 2
	}
	c.UI.Output(fmt.Sprintf("\nCurrent context: %s\n", string(defaultContext)))

	return OutputData(c.UI, contexts)
}
