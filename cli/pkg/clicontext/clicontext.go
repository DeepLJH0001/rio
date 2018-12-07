package clicontext

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/rio/cli/pkg/clientcfg"
	projectclient "github.com/rancher/rio/types/client/project/v1"
	"github.com/rancher/rio/types/client/rio/v1"
	"github.com/urfave/cli"
)

const dataKey = "config"

type CLIContext struct {
	*clientcfg.Config
	Ctx context.Context
	CLI *cli.Context
	WC  *client.Client
	PC  *projectclient.Client
}

func (c *CLIContext) ClientLookup(typeName string) (clientbase.APIBaseClientInterface, error) {
	if c.WC != nil && c.PC != nil {
		if _, ok := c.WC.Types[typeName]; ok {
			return c.WC, nil
		}
		return c.PC, nil
	}
	wc, err := c.Config.ProjectClient()
	if err == nil {
		if _, ok := wc.Types[typeName]; ok {
			return wc, nil
		}
	}

	cc, err := c.Config.ClusterClient()
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func Lookup(data map[string]interface{}) *CLIContext {
	return data[dataKey].(*CLIContext)
}

func (c *CLIContext) Store(data map[string]interface{}) {
	data[dataKey] = c
}

func Wrap(f func(*CLIContext) error) func(context2 *cli.Context) error {
	return func(app *cli.Context) error {
		cc := Lookup(app.App.Metadata)
		cc.CLI = app
		return f(cc)
	}
}

func DefaultAction(action interface{}) interface{} {
	if fn, ok := action.(func(ctx *CLIContext) error); ok {
		return Wrap(func(ctx *CLIContext) error {
			if ctx.CLI.Bool("help") {
				cli.ShowAppHelp(ctx.CLI)
				return nil
			}
			return fn(ctx)
		})
	}
	return action
}
