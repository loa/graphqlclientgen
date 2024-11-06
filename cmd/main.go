package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"golang.org/x/exp/slog"
)

func main() {
	cmd := &cli.Command{
		Name:  "graphqlclientgen",
		Usage: "Generate golang graphql clients",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "verbose logs",
				Value: false,
			},
		},
		Before: func(ctx context.Context, c *cli.Command) error {
			var logLevel = new(slog.LevelVar)
			if c.Bool("verbose") {
				logLevel.Set(slog.LevelDebug)
			}
			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: logLevel,
			}))
			slog.SetDefault(logger)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "generate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "filename",
						Category: "config",
						Value:    "graphqlclientgen.yaml",
					},
				},
				Action: actionGenerate,
			},
			{
				Name: "init",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "package-name",
					},
					&cli.StringFlag{
						Name:  "schema-path",
						Value: "../graph/*.graphqls",
					},
				},
				Action: actionInit,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
