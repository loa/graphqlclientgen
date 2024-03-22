package codegen

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

func (gen *Generator) parseSchema() error {
	// TODO: add schema comments to golang functions
	namedTypes := []string{}

	nt, err := gen.parseFunction(gen.schemas.Query)
	if err != nil {
		return err
	}
	namedTypes = append(namedTypes, nt...)

	nt, err = gen.parseFunction(gen.schemas.Mutation)
	if err != nil {
		return err
	}
	for _, namedType := range nt {
		if !slices.Contains(namedTypes, namedType) {
			namedTypes = append(namedTypes, namedType)
		}
	}

	// generate all required types definitions
	for _, t := range namedTypes {
		gen.Schema.Types = append(gen.Schema.Types, SchemaType{
			Name: t,
		})
	}

	return nil
}

func (gen *Generator) parseFunction(definition *ast.Definition) ([]string, error) {
	namedTypes := []string{}

	// add client query functions
	for _, field := range definition.Fields {
		if strings.HasPrefix(field.Name, "__") {
			// skip __schema, __type, etc
			continue
		}

		var namedType string
		var schemaType SchemaType
		if field.Type.NamedType == "" {
			namedType = field.Type.Elem.NamedType
			schemaType = SchemaType{
				Name:    field.Type.Elem.NamedType,
				NonNull: field.Type.Elem.NonNull,

				List:        true,
				ListNonNull: field.Type.NonNull,
			}
		} else {
			namedType = field.Type.NamedType
			schemaType = SchemaType{
				Name:    field.Type.NamedType,
				NonNull: field.Type.NonNull,
			}
		}

		schemaFunction := SchemaFunction{
			Name:      field.Name,
			QueryType: strings.ToLower(definition.Name),
			Type:      schemaType,
		}

		// add returned type to types
		if !slices.Contains(namedTypes, namedType) {
			namedTypes = append(namedTypes, namedType)
		}

		slog.Debug("parseFunction", "type", definition.Name, "name", field.Name)
		gen.Schema.Functions = append(gen.Schema.Functions, schemaFunction)
	}

	return namedTypes, nil
}
