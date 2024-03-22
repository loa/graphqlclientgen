package main

import (
	"log"
	"os"

	"github.com/loa/graphqlclientgen/cmd/generate"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "graphqlclientgen",
		Usage:  "Generate golang graphql clients",
		Action: generate.Cmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
