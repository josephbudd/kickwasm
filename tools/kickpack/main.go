package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/tools/kickpack/nopack"
	"github.com/josephbudd/kickwasm/tools/kickpack/pack"
	"github.com/pkg/errors"
)

// DestinationFlag is the output package name.
var DestinationFlag string

// MustExistFlag indicates that missing folders and files to not generate errors.
var MustExistFlag bool

// NoUsageFlag indicates that the usage should not be printed with errors.
var NoUsageFlag bool

// DontPackFlag indicates that the source folder is not packed.
var DontPackFlag bool

var intro = `
About kickpack:

    kickpack packs files and files in folders into a go package containing
	  a map of file paths to file bytes.
	
	example:
	the current folder contains the file me.txt.
	the folder ./MyStuff/ contains the files things.go and images/Flower.jpg.
	the folder ../MyOtherStuff/ contains the files others.json and css/forms.css.
	$ kickpack -o ../mysite ./MyStuff ./me.txt ../MyOtherStuff
	will
	 * create a file "../mysite/mysite.go
	 * where package mysite defines
	
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

	var fileStore = map[string][]byte{
		{"MyStuff/things.go":[]byte{...}},
		{"MyStuff/images/Flower.jpg":[]byte{...}},
		{"me.txt":[]byte{...}},
		{"../MyOtherStuff/others.json":[]byte{...}},
		{"../MyOtherStuff/css/forms.css":[]byte{...}},
	}
`

func init() {
	// DontPackFlag
	flag.BoolVar(&DontPackFlag, "nopack", false, "source files are not packed into the package")
	flag.StringVar(&DestinationFlag, "o", "", "the output go package folder")
	flag.BoolVar(&MustExistFlag, "strict", false, "source files and folders must exist")
	flag.BoolVar(&NoUsageFlag, "nu", false, "don't print usage with errors")
}

func main() {
	var err error

	defer func() {
		if err != nil {
			os.Exit(-1)
		}
	}()

	flag.Parse()
	var output, packageName string
	if output, packageName, err = getOutput(); err != nil {
		usage(err)
		return
	}
	var sources []string
	if sources, err = getSources(); err != nil {
		usage(err)
		return
	}
	if DontPackFlag {
		if err = nopack.Build(output, packageName); err != nil {
			fmt.Println("nopack.Build error is ", err.Error())
		}
	} else {
		if err = pack.Build(output, sources, packageName, MustExistFlag); err != nil {
			fmt.Println("pack.Build error is ", err.Error())
		}
	}

}

func getOutput() (output, packageName string, err error) {
	if len(DestinationFlag) == 0 {
		err = errors.New("missing output path -o")
		return
	}
	if strings.IndexAny(DestinationFlag, `;:`) >= 0 {
		err = errors.New("-o has illegal characters")
		return
	}
	dir := filepath.Dir(DestinationFlag)
	packageName = filepath.Base(DestinationFlag)
	if packageName != strings.ToLower(packageName) {
		err = errors.New("the output folder must be lower case")
	}
	first := packageName[:1]
	// if (first >= "0" && first <= "9") || strings.IndexAny(packageName, ` ~!@#$%^&*()-_+={}[]|'",./<>?`) >= 0 {
	// 	err = errors.New(fmt.Sprintf("the output folder must be a valid go package name not %q", packageName))
	// }
	if first >= "0" && first <= "9" {
		err = errors.New(fmt.Sprintf("1 the output folder must be a valid go package name not %q", packageName))
	}
	if strings.IndexAny(packageName, ` ~!@#$%^&*()-_+={}[]|'",./<>?`) >= 0 {
		err = errors.New(fmt.Sprintf("2 the output folder must be a valid go package name not %q", packageName))
	}
	output = filepath.Join(dir, packageName)
	return
}

func getSources() (sources []string, err error) {
	sources = flag.Args()
	if len(sources) == 0 {
		err = errors.New("missing sources")
		return
	}
	for _, s := range sources {
		if strings.IndexAny(s, `;:`) >= 0 {
			err = errors.New(fmt.Sprintf("%q is not a valid source.", s))
			return
		}
	}
	return
}

func usage(err error) {
	if err != nil {
		fmt.Println("kickpack has encountered an error.")
		fmt.Println(" * ", err.Error())
		fmt.Println()
	}
	if NoUsageFlag {
		return
	}
	fmt.Println("Usage:")
	flag.PrintDefaults()
	fmt.Println(intro)
}
