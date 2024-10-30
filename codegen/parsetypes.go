package codegen

import (
	"cmp"
	"log/slog"
	"slices"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

// TODO: enable custom typeMappings through config
var typeMappings = map[string]string{
	"Boolean": "bool",
	"ID":      "string",
	"Int":     "int",
	"Float":   "float64",
	"String":  "string",
}

func (gen *Generator) parseTypes(namedTypes []string) error {
	// add field types of function input/return types
	namedTypes, err := gen.findFieldTypes(namedTypes)
	if err != nil {
		return err
	}

	for _, t := range gen.schemas.Types {
		// skip unused types
		if !slices.Contains(namedTypes, t.Name) {
			continue
		}
		// skip types with mappings
		if _, ok := typeMappings[t.Name]; ok {
			continue
		}

		slog.Debug("parseTypes", "type", t.Name)

		// TODO: support all kinds of types
		// only support scalar and object kind
		if !slices.Contains([]ast.DefinitionKind{ast.Object, ast.InputObject, ast.Scalar}, t.Kind) {
			slog.Warn("unsupported type kind", "kind", t.Kind)
			continue
		}

		schemaType := SchemaType{
			Name:        t.Name,
			Type:        t.Name,
			Description: t.Description,
			Kind:        string(t.Kind),
			Fields:      map[string]SchemaType{},
		}

		for _, field := range t.Fields {
			if tt, ok := gen.schemas.Types[field.Type.Name()]; ok {
				name := tt.Name
				// override type name with golang type for types listed in typeMappings
				if val, ok := typeMappings[name]; ok {
					name = val
				}

				fieldType := SchemaType{
					Name: name,
					Type: name,
					Kind: string(tt.Kind),

					Description: field.Description,
					NonNull:     field.Type.NonNull,
				}

				schemaType.Fields[field.Name] = fieldType
			}
		}

		gen.Schema.Types = append(gen.Schema.Types, schemaType)
	}

	// sort types so generated code stays stable
	slices.SortFunc(gen.Schema.Types, func(a, b SchemaType) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	return nil
}

func (gen *Generator) findFieldTypes(namedTypes []string) ([]string, error) {
	visited := []string{}

	// keep iterating while appending slice
	for i := 0; i < len(namedTypes); i++ {
		namedType := namedTypes[i]

		// only check type ones, graphql types is able to have circular dependencies
		if slices.Contains(visited, namedType) {
			continue
		}
		visited = append(visited, namedType)

		// add all fields to namedTypes if missing
		for _, field := range gen.schemas.Types[namedType].Fields {
			var name string
			if field.Type.NamedType == "" {
				name = field.Type.Elem.NamedType
			} else {
				name = field.Type.NamedType

			}

			if slices.Contains(visited, name) {
				continue
			}
			namedTypes = append(namedTypes, name)
		}
	}

	return namedTypes, nil
}
