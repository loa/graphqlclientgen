package codegen

import (
	"slices"
)

func (gen *Generator) parseSchema() error {
	namedTypes := []string{}

	nt, err := gen.parseFunctions(gen.schemas.Query)
	if err != nil {
		return err
	}
	namedTypes = append(namedTypes, nt...)

	nt, err = gen.parseFunctions(gen.schemas.Mutation)
	if err != nil {
		return err
	}
	for _, namedType := range nt {
		if !slices.Contains(namedTypes, namedType) {
			namedTypes = append(namedTypes, namedType)
		}
	}

	if err := gen.parseTypes(namedTypes); err != nil {
		return err
	}

	return nil
}
