package main

import (
	"context"
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/protocol/mqtt"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var cfgPath string
	function := &cli.App{
		Name:    "Homo Function",
		Version: "0.0.1",
		Usage:   "Function manager for homo",
		Flags: []cli.Flag{
			&cli.StringFlag{
				EnvVars:     []string{"HOMO_FUNCTION_CONFIG_FILE"},
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: homo.DefaultConfFile,
				Usage:       "set homo function config file path",
				Destination: &cfgPath,
			},
		},
		Action: func(c *cli.Context) error {
			homo.Run(homo.Service{CfgPath: cfgPath}, func(ctx homo.Context) error {
				var cfg Config
				err := ctx.LoadConfig(cfgPath, &cfg)
				if err != nil {
					return err
				}
				functions := make(map[string]*Function)
				for _, fi := range cfg.Functions {
					functions[fi.Name] = NewFunction(fi, newProducer(ctx, fi), ctx.Log())
				}
				defer func() {
					for _, f := range functions {
						f.Close()
					}
				}()
				rulers := make([]*ruler, 0)
				defer func() {
					for _, ruler := range rulers {
						ruler.close()
					}
				}()
				for _, ri := range cfg.Rules {
					f, ok := functions[ri.Function.Name]
					if !ok {
						return fmt.Errorf("function (%s) not found", ri.Function.Name)
					}
					c, err := ctx.NewHubClient(ri.ClientID, []mqtt.TopicInfo{ri.Subscribe})
					if err != nil {
						return fmt.Errorf("failed to create hub client: %s", err.Error())
					}
					rulers = append(rulers, newRuler(ri, c, f, ctx.Log()))
				}
				for _, ruler := range rulers {
					err := ruler.start()
					if err != nil {
						return err
					}
				}
				if cfg.Server.Address != "" {
					svr, err := homo.NewFServer(cfg.Server, func(ctx context.Context, msg *homo.FunctionMessage) (*homo.FunctionMessage, error) {
						f, ok := functions[msg.FunctionName]
						if !ok {
							return nil, fmt.Errorf("function (%s) not found", msg.FunctionName)
						}
						return f.Call(msg)
					})
					if err != nil {
						return err
					}
					defer svr.Close()
				}
				ctx.Wait()
				return nil
			})
			return nil
		},
	}
	if err := function.Run(os.Args); err != nil {
		logger.S.Fatal(err)
	}
}
