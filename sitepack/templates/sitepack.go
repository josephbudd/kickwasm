package templates

// SitePack is the ../{{.PackageName}}.go file.
const SitePack = `package {{.PackageName}}

/*
	Note Well:

	This {{.PackageName}} package was created by kickwasm.
	It is only a temporary place holder.
	You will replace this file with your true {{.PackageName}} package
	  when you build the renderer process.
*/

func Contents(path string) (markupbb []byte, found bool) {
	return
}

func Paths() (paths []string) {
	paths = make([]string, 0)
	return
}
`
