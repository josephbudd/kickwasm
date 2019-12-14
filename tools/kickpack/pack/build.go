package pack

import "fmt"

// Build creates and writes the sources into a single package.
// Param output is the output packages folder path.
// Param sources is the source folder paths.
// Param packageName is the name of the output package.
// Param mustExist means that all source folders must exist.
func Build(output string, sources []string, packageName string, mustExist bool) (err error) {
	var pathBytes map[string][]byte
	if pathBytes, err = buildPathBytes(sources, mustExist); err != nil {
		err = fmt.Errorf("pack.Build read: %w", err)
		return
	}
	if err = write(pathBytes, output, packageName); err != nil {
		err = fmt.Errorf("pack.Build write: %w", err)
	}
	return
}

func buildPathBytes(sources []string, mustExist bool) (pathBytes map[string][]byte, err error) {
	pathBytes = make(map[string][]byte, 100)
	for _, s := range sources {
		if err = slurp(s, mustExist, pathBytes); err != nil {
			return
		}
	}
	return
}
