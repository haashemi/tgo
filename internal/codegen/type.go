package main

import (
	"fmt"
	"strings"
)

func GenerateType(g Group) string {
	var fields string
	for _, f := range g.Fields {
		tag := f.Name
		if f.IsOptional {
			tag += ",omitempty"
		}

		desc := strings.ReplaceAll(f.Description, "\n", " ")
		if f.IsOptional && !strings.HasPrefix(desc, "Optional.") {
			desc = "Optional. " + desc
		}

		fields += fmt.Sprintf("    %s %s `json:\"%s\"` // %s\n", snakeToPascal(f.Name), typeOf(f.Name, f.Type), tag, desc)
	}

	return fmt.Sprintf("// %s %s\ntype %s struct {\n%s}", g.Name, g.Description, g.Name, fields)
}
