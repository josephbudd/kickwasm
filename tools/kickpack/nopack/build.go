package nopack

import "fmt"

// Build creates and writes the sources into a single package.
// Param output is the output packages folder path.
// Param packageName is the name of the output package.
func Build(output string, packageName string) (err error) {
	if err = write(output, packageName); err != nil {
		err = fmt.Errorf("partial.Build write: %w", err)
	}
	return
}
