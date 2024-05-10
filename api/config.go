// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/openbao/openbao/sdk/helper/hclutil"
)

const (
	// DefaultConfigPath is the default path to the configuration file
	DefaultConfigPath = "~/.bao"

	// ConfigPathEnv is the environment variable that can be used to
	// override where the Vault configuration is.
	ConfigPathEnv = "BAO_CONFIG_PATH"
)

// Config is the CLI configuration for Vault that can be specified via
// a `$HOME/.vault` file which is HCL-formatted (therefore HCL or JSON).
type DefaultBaoConfig struct {
	// TokenHelper is the executable/command that is executed for storing
	// and retrieving the authentication token for the Vault CLI. If this
	// is not specified, then vault's internal token store will be used, which
	// stores the token on disk unencrypted.
	TokenHelper    string             `hcl:"token_helper"`
	CurrentContext string             `hcl:"current_context"`
	Contexts       map[string]Context `hcl:"contexts"`
}

type Context struct {
	Server    string `hcl:"server"`
	Namespace string `hcl:"namespace"`
}

// Config loads the configuration and returns it. If the configuration
// is already loaded, it is returned.
func BaoConfig() (*DefaultBaoConfig, error) {
	var err error
	config, err := LoadConfig("")
	if err != nil {
		return nil, err
	}

	return config, nil
}

// LoadConfig reads the configuration from the given path. If path is
// empty, then the default path will be used, or the environment variable
// if set.
func LoadConfig(path string) (*DefaultBaoConfig, error) {
	if path == "" {
		path = DefaultConfigPath
	}
	if v := ReadBaoVariable(ConfigPathEnv); v != "" {
		path = v
	}

	// NOTE: requires HOME env var to be set
	path, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("error expanding config path %q: %w", path, err)
	}

	contents, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	conf, err := ParseConfig(string(contents))
	if err != nil {
		return nil, fmt.Errorf("error parsing config file at %q: %w; ensure that the file is valid; Ansible Vault is known to conflict with it", path, err)
	}

	return conf, nil
}

// ParseConfig parses the given configuration as a string.
func ParseConfig(contents string) (*DefaultBaoConfig, error) {
	root, err := hcl.Parse(contents)
	if err != nil {
		return nil, err
	}

	// Top-level item should be the object list
	list, ok := root.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("failed to parse config; does not contain a root object")
	}

	valid := []string{
		"token_helper",
		"current_context",
		"contexts",
	}
	if err := hclutil.CheckHCLKeys(list, valid); err != nil {
		return nil, err
	}

	var c DefaultBaoConfig
	if err := hcl.DecodeObject(&c, list); err != nil {
		return nil, err
	}
	return &c, nil
}

// CurrentContext returns the current context configured for Vault.
// This helper should only be used for non-server CLI commands.
func CurrentContext() (Context, error) {
	config, err := LoadConfig("")
	if err != nil {
		return Context{}, err
	}
	return config.Contexts[config.CurrentContext], nil
}

// ListContexts returns the list of contexts configured for Vault and its default.
// This helper should only be used for non-server CLI commands.
func ListContexts() (map[string]interface{}, string, error) {
	config, err := LoadConfig("")
	if err != nil {
		return nil, "", err
	}
	// fmt.Printf("Decoded Context: %#v\n\n", config.Contexts[config.CurrentContext])
	// Convert the contexts to a printable struct - when using OutputMap (key + value)
	contexts := make(map[string]interface{}, len(config.Contexts))
	for k, v := range config.Contexts {
		newMap := map[string]interface{}{"namespace": v.Namespace, "server": v.Server}
		contexts[k] = newMap
	}
	return contexts, config.CurrentContext, nil
}
