package codegen

import (
	"cmp"
	"log/slog"
	"slices"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

func (gen *Generator) parseFunctions(definition *ast.Definition) ([]string, error) {
	namedTypes := []string{}

	// add client query functions
	for _, field := range definition.Fields {
		if strings.HasPrefix(field.Name, "__") {
			// skip __schema, __type, etc
			continue
		}

		slog.Debug("parseFunction", "type", definition.Name, "name", field.Name)

		var namedType string
		var schemaType SchemaType
		if field.Type.NamedType == "" {
			namedType = field.Type.Elem.NamedType
			goType := namedType
			kind := "OBJECT"
			if mapping, ok := gen.Config.TypeMappings[namedType]; ok {
				goType = typeName(mapping)
				kind = "SCALAR"
			}

			schemaType = SchemaType{
				Name:    field.Type.Elem.NamedType,
				Type:    goType,
				Kind:    kind,
				NonNull: field.Type.Elem.NonNull,

				List:        true,
				ListNonNull: field.Type.NonNull,
			}
		} else {
			namedType = field.Type.NamedType
			goType := namedType
			kind := "OBJECT"
			if mapping, ok := gen.Config.TypeMappings[namedType]; ok {
				goType = typeName(mapping)
				kind = "SCALAR"
			}

			schemaType = SchemaType{
				Name:    field.Type.NamedType,
				Type:    goType,
				Kind:    kind,
				NonNull: field.Type.NonNull,
			}
		}

		// add returned type to types
		if !slices.Contains(namedTypes, namedType) {
			namedTypes = append(namedTypes, namedType)
		}

		arguments := map[string]SchemaType{}
		for _, argument := range field.Arguments {
			slog.Debug("parseFunction argument", "name", field.Name, "argument", argument.Name)

			var namedType string
			var argumentType SchemaType
			if argument.Type.NamedType == "" {
				namedType = argument.Type.Elem.NamedType
				goType := namedType
				kind := "OBJECT"
				if mapping, ok := gen.Config.TypeMappings[namedType]; ok {
					goType = typeName(mapping)
					kind = "SCALAR"
				}

				argumentType = SchemaType{
					Name:    namedType,
					Type:    goType,
					Kind:    kind,
					NonNull: argument.Type.Elem.NonNull,

					List:        true,
					ListNonNull: field.Type.NonNull,
				}
			} else {
				namedType = argument.Type.NamedType
				goType := namedType
				kind := "OBJECT"
				if mapping, ok := gen.Config.TypeMappings[namedType]; ok {
					goType = typeName(mapping)
					kind = "SCALAR"
				}

				argumentType = SchemaType{
					Name:    namedType,
					Type:    goType,
					Kind:    kind,
					NonNull: argument.Type.NonNull,
				}
			}

			// add input type to types
			if !slices.Contains(namedTypes, namedType) {
				namedTypes = append(namedTypes, namedType)
			}

			arguments[argument.Name] = argumentType
		}

		gen.Schema.Functions = append(gen.Schema.Functions, SchemaFunction{
			Name:        field.Name,
			QueryType:   strings.ToLower(definition.Name),
			Type:        schemaType,
			Description: field.Description,
			Arguments:   arguments,
		})
	}

	// sort types so generated code stays stable
	slices.SortFunc(gen.Schema.Functions, func(a, b SchemaFunction) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	return namedTypes, nil
}
