package codegen

import (
	"embed"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"gopkg.in/yaml.v3"
)

type (
	Generator struct {
		Config Config
		Schema Schema

		templates *template.Template
		schemas   *ast.Schema
	}

	Config struct {
		Schema       []string               `yaml:"schema"`
		Client       ConfigClient           `yaml:"client"`
		TypeMappings map[string]TypeMapping `yaml:"typeMappings"`

		SkipGofmt        bool `yaml:"skipGofmt"`
		CreateSchemaYaml bool `yaml:"createSchemaYaml"`
	}

	ConfigClient struct {
		Dir     string
		Package string
	}

	TypeMapping struct {
		Alias  *string `yaml:"alias"`
		Import *string `yaml:"import"`
		Name   string  `yaml:"name"`
	}

	Schema struct {
		Functions []SchemaFunction
		Types     []SchemaType
	}

	SchemaFunction struct {
		Name string
		// QueryType is graphql type [query, mutation]
		QueryType   string
		Description string

		Type      SchemaType
		Arguments map[string]SchemaType
	}

	SchemaType struct {
		Name        string
		Type        string
		NonNull     bool
		Description string
		Kind        string

		List        bool
		ListNonNull bool

		Fields map[string]SchemaType
	}
)

//go:embed *.gotpl
var tplFiles embed.FS

func New(filename string) (*Generator, error) {
	gen := Generator{
		Config: Config{
			Client: ConfigClient{
				Dir:     ".",
				Package: "client",
			},
			TypeMappings: map[string]TypeMapping{},
		},
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &gen.Config); err != nil {
		return nil, err
	}

	defaultTypeMappings := map[string]string{
		"Boolean": "bool",
		"ID":      "string",
		"Int":     "int",
		"Float":   "float64",
		"String":  "string",
	}
	for key, value := range defaultTypeMappings {
		if _, ok := gen.Config.TypeMappings[key]; !ok {
			gen.Config.TypeMappings[key] = TypeMapping{
				Name: value,
			}
		}
	}

	if err := gen.loadSchemas(); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (gen *Generator) Generate() error {
	if err := gen.loadTemplates(); err != nil {
		return err
	}

	if err := gen.parseSchema(); err != nil {
		return err
	}

	path := filepath.Dir(gen.Config.Client.Dir)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	filename := filepath.Join(gen.Config.Client.Dir, "generated.go")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := gen.templates.Execute(f, gen); err != nil {
		return err
	}

	slog.Debug("config",
		"SkipGofmt", gen.Config.SkipGofmt,
		"CreateSchemaYaml", gen.Config.CreateSchemaYaml,
	)
	if !gen.Config.SkipGofmt {
		if err := gofmt(filename); err != nil {
			return err
		}
	}

	if gen.Config.CreateSchemaYaml {
		filename := filepath.Join(gen.Config.Client.Dir, "generated.yaml")
		yf, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer yf.Close()

		if err := yaml.NewEncoder(yf).Encode(gen.Schema); err != nil {
			return err
		}
	}

	return nil
}

func (gen *Generator) loadTemplates() error {
	tpl := template.New("generated.gotpl").Funcs(template.FuncMap{
		"capitalize":  capitalize,
		"initialism":  initialism,
		"stripPrefix": stripPrefix,
	})

	var err error
	gen.templates, err = tpl.ParseFS(tplFiles, "generated.gotpl")
	return err
}

func (gen *Generator) loadSchemas() error {
	var sources []*ast.Source
	for _, schemaPath := range gen.Config.Schema {
		filenames, err := filepath.Glob(schemaPath)
		if err != nil {
			return err
		}

		for _, filename := range filenames {
			content, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			source := ast.Source{
				Name:  filepath.Base(filename),
				Input: string(content),
			}

			sources = append(sources, &source)
			slog.Debug("parse", "name", source.Name)
		}
	}

	var err error
	gen.schemas, err = gqlparser.LoadSchema(sources...)
	return err
}

func gofmt(filename string) error {
	cmd := exec.Command("go", "fmt", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
