package codegen

import (
	"embed"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

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
		Schema []string
		Client ConfigClient
	}

	ConfigClient struct {
		Dir     string
		Package string
	}

	Schema struct {
		Functions []SchemaFunction
		Types     []SchemaType
	}

	SchemaFunction struct {
		Name string
		// QueryType is graphql type [query, mutation]
		QueryType string

		Arguments []SchemaFunctionArgument
		Type      SchemaType
	}

	SchemaFunctionArgument struct {
		Name string
		Type SchemaType
	}

	SchemaType struct {
		Name    string
		NonNull bool

		List        bool
		ListNonNull bool
	}
)

//go:embed *.gotpl
var tplFiles embed.FS

func New(filename string) (*Generator, error) {
	gen := Generator{
		Config: Config{
			Client: ConfigClient{
				Dir:     "client",
				Package: "client",
			},
		},
	}

	// TODO: use relative path from config
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &gen.Config); err != nil {
		return nil, err
	}

	// TODO: parse graphql files
	if err := gen.loadSchemas(); err != nil {
		return nil, err
	}

	return &gen, nil
}

func (gen *Generator) Generate() error {
	slog.Info("generating.")

	if err := gen.loadTemplates(); err != nil {
		return err
	}

	if err := gen.parseSchemas(); err != nil {
		return err
	}

	// TODO: use relative path from config
	path := filepath.Dir(gen.Config.Client.Dir)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	// TODO: use relative path from config
	filename := filepath.Join(gen.Config.Client.Dir, "graphqlclient.go")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := gen.templates.Execute(f, gen); err != nil {
		return err
	}

	if err := gofmt(filename); err != nil {
		return err
	}

	return nil
}

func (gen *Generator) loadTemplates() error {
	tpl := template.New("graphqlclient.gotpl").Funcs(template.FuncMap{
		"capitalize": func(in string) (string, error) {
			runes := []rune(in)
			runes[0] = unicode.ToUpper(runes[0])
			return string(runes), nil
		},
	})

	var err error
	gen.templates, err = tpl.ParseFS(tplFiles, "graphqlclient.gotpl")
	return err
}

func (gen *Generator) loadSchemas() error {
	var sources []*ast.Source
	for _, schemaPath := range gen.Config.Schema {
		// TODO: use relative path from config
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

func (gen *Generator) parseSchemas() error {
	// TODO: add schema comments to golang functions
	types := map[string]bool{}

	// add client query functions
	for _, field := range gen.schemas.Query.Fields {
		if strings.HasPrefix(field.Name, "__") {
			// skip __schema, __type, etc
			continue
		}

		var schemaType SchemaType
		if field.Type.NamedType == "" {
			schemaType = SchemaType{
				Name:    field.Type.Elem.NamedType,
				NonNull: field.Type.Elem.NonNull,

				List:        true,
				ListNonNull: field.Type.NonNull,
			}
		} else {
			schemaType = SchemaType{
				Name:    field.Type.NamedType,
				NonNull: field.Type.NonNull,
			}
		}

		schemaFunction := SchemaFunction{
			Name:      field.Name,
			QueryType: "query",
			Type:      schemaType,
		}

		// add returned type to types
		types[field.Type.Elem.NamedType] = true

		slog.Debug("query field", "name", field.Name)
		gen.Schema.Functions = append(gen.Schema.Functions, schemaFunction)
	}

	// add client mutation functions
	for _, field := range gen.schemas.Mutation.Fields {
		if strings.HasPrefix(field.Name, "__") {
			// skip __schema, __type, etc
			continue
		}

		var schemaType SchemaType
		if field.Type.NamedType == "" {
			schemaType = SchemaType{
				Name:    field.Type.Elem.NamedType,
				NonNull: field.Type.Elem.NonNull,

				List:        true,
				ListNonNull: field.Type.NonNull,
			}
		} else {
			schemaType = SchemaType{
				Name:    field.Type.NamedType,
				NonNull: field.Type.NonNull,
			}
		}

		schemaFunction := SchemaFunction{
			Name:      field.Name,
			QueryType: "mutation",

			Type: schemaType,
		}

		// add returned type to types
		types[field.Type.NamedType] = true

		slog.Debug("mutation field", "name", field.Name)
		gen.Schema.Functions = append(gen.Schema.Functions, schemaFunction)
	}

	// generate all required types definitions
	for t := range types {
		gen.Schema.Types = append(gen.Schema.Types, SchemaType{
			Name: t,
		})
	}

	return nil
}

func gofmt(filename string) error {
	cmd := exec.Command("go", "fmt", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
