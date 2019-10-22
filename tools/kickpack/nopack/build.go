package nopack

import (
	"github.com/pkg/errors"
)

// Build creates and writes the sources into a single package.
// Param output is the output packages folder path.
// Param packageName is the name of the output package.
func Build(output string, packageName string) (err error) {
	if err = write(output, packageName); err != nil {
		err = errors.WithMessage(err, "partial.Build write")
	}
	return
}
