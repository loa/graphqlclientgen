package main

import (
	"log/slog"
	"os"

	"github.com/loa/graphqlclientgen/codegen"
	"github.com/urfave/cli/v2"
)

func generate(*cli.Context) error {
	opts := &slog.HandlerOptions{
		// AddSource: true,
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	// TODO: add positional argument of graphqlclient.yml
	// TODO: find location of graphqlclient.yml recursively downward
	filename := "graphqlclientgen.yml"

	generator, err := codegen.New(filename)
	if err != nil {
		return err
	}

	return generator.Generate()
}
