package format

import (
	"fmt"
	"sort"
	"strings"
)

// FixImports sorts and prefixes and quotes imports.
func FixImports(imports []string) (fixed []string) {
	fixed = make([]string, len(imports))
	sort.Strings(imports)
	for i, s := range imports {
		parts := strings.Split(s, "/")
		name := parts[len(parts)-1]
		lname := strings.ToLower(name)
		if name != lname {
			fixed[i] = fmt.Sprintf("%s %q", lname, s)
		} else {
			fixed[i] = fmt.Sprintf("%q", s)
		}
	}
	return
}
