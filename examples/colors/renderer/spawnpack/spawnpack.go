package spawnpack

// Contents returns the contents of the file at path and if found.
func Contents(path string) (contents []byte, found bool) {
	contents, found = fileStore[path]
	return
}

// Paths returns a slice of the file paths.
func Paths() (paths []string) {
	l := len(fileStore)
	paths = make([]string, 0, l)
	for k := range fileStore {
		paths = append(paths, k)
	}
	return
}

// fileStore is a store of various files.
var fileStore =  map[string][]byte{
    "spawnTemplates":[]byte{},
}
