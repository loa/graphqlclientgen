package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

var (
	tmplContent = "package {{ . }}\n\n//go:generate go tool graphqlclientgen generate\n"
)

func actionInit(ctx context.Context, c *cli.Command) error {
	packageName := c.String("package-name")
	if len(packageName) < 1 {
		// take current directory name, if packageName is empty
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		packageName = filepath.Base(path)
	}

	{
		f, err := os.Create("graphqlclientgen.yaml")
		if err != nil {
			return err
		}
		defer f.Close()

		config := struct {
			Schema       []string                     `yaml:"schema"`
			Client       map[string]string            `yaml:"client"`
			TypeMappings map[string]map[string]string `yaml:"typeMappings"`
		}{
			Schema: []string{c.String("schema-path")},
			Client: map[string]string{
				"dir":     ".",
				"package": packageName,
			},
			TypeMappings: map[string]map[string]string{
				"Int64": {
					"name": "int",
				},
				"Time": {
					"name":   "Time",
					"import": "time",
				},
				"UUID": {
					"name":   "UUID",
					"import": "github.com/google/uuid",
				},
			},
		}

		if err := yaml.NewEncoder(f).Encode(config); err != nil {
			return err
		}
	}

	{
		f, err := os.Create(fmt.Sprintf("%s.go", packageName))
		if err != nil {
			return err
		}
		defer f.Close()

		tmpl, err := template.New("").Parse(tmplContent)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, packageName); err != nil {
			return err
		}
	}

	return nil
}
