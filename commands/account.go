/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/digitalocean/doctl/commands/displayers"
	"github.com/spf13/cobra"
)

// Account creates the account commands hierarchy.
func Account() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:   "account",
			Short: "account commands",
			Long:  "account is used to access account commands",
		},
		DocCategories: []string{"account"},
		IsIndex:       true,
	}

	CmdBuilder(cmd, RunAccountGet, "get", "get account", Writer,
		aliasOpt("g"), displayerType(&displayers.Account{}), docCategories("account"))

	CmdBuilder(cmd, RunAccountRateLimit, "ratelimit", "get API rate limits", Writer,
		aliasOpt("rl"), displayerType(&displayers.RateLimit{}), docCategories("account"))

	return cmd
}

// RunAccountGet runs account get.
func RunAccountGet(c *CmdConfig) error {
	a, err := c.Account().Get()
	if err != nil {
		return err
	}

	return c.Display(&displayers.Account{Account: a})
}

// RunAccountRateLimit retrieves API rate limits for the account.
func RunAccountRateLimit(c *CmdConfig) error {
	rl, err := c.Account().RateLimit()
	if err != nil {
		return err
	}

	return c.Display(&displayers.RateLimit{RateLimit: rl})
}
