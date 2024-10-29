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
				Action: generate,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
