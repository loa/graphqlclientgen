package codegen

import (
	"fmt"
	"strings"
)

func typeName(mapping TypeMapping) string {
	if mapping.Alias != nil {
		return fmt.Sprintf("%s.%s", *mapping.Alias, mapping.Name)
	}

	if mapping.Import != nil {
		base := (*mapping.Import)[strings.LastIndex(*mapping.Import, "/")+1:]
		return fmt.Sprintf("%s.%s", base, mapping.Name)
	}

	return mapping.Name
}
