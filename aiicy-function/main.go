package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy/protocol/mqtt"
	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/urfave/cli/v2"
)

func main() {
	var cfgPath string
	function := &cli.App{
		Name:    "Aiicy Function",
		Version: "0.0.1",
		Usage:   "Function manager for aiicy",
		Flags: []cli.Flag{
			&cli.StringFlag{
				EnvVars:     []string{"AIICY_FUNCTION_CONFIG_FILE"},
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: aiicy.DefaultConfFile,
				Usage:       "set aiicy function config file path",
				Destination: &cfgPath,
			},
		},
		Action: func(c *cli.Context) error {
			aiicy.Run(aiicy.Service{CfgPath: cfgPath}, func(ctx aiicy.Context) error {
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
					svr, err := aiicy.NewFServer(cfg.Server, func(ctx context.Context, msg *aiicy.FunctionMessage) (*aiicy.FunctionMessage, error) {
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
