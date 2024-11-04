package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/loa/graphqlclientgen/codegen"
	"github.com/urfave/cli/v3"
)

func actionGenerate(ctx context.Context, c *cli.Command) error {
	configDir, err := findConfigDir(c.String("filename"))
	if err != nil {
		return err
	}

	generator, err := codegen.New(configDir, c.String("filename"))
	if err != nil {
		return err
	}

	return generator.Generate()
}

func findConfigDir(filename string) (string, error) {
	// check if user specified abs/relative path to configfile
	if f, err := os.Stat(filename); err == nil {
		if f.Mode().IsDir() {
			return "", fmt.Errorf("%s is not a file", filename)
		}

		file, err := filepath.Abs(filename)
		if err != nil {
			return "", err
		}

		return filepath.Dir(file), nil
	}

	if filepath.Dir(filename) != "." {
		// user provided filename with path that did not exist
		return "", fmt.Errorf("%s not found", filename)
	}

	// start search in current working directory
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for path != "/" {
		file := filepath.Join(path, filename)

		if _, err := os.Stat(file); err == nil {
			// config file found
			return path, nil
		}

		// step up to parent directory
		path = filepath.Dir(path)
	}

	return "", fmt.Errorf("%s not found", filename)
}
