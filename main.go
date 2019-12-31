package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/josephbudd/kickwasm/pkg"
	"github.com/josephbudd/kickwasm/pkg/kickwasm"
	"github.com/josephbudd/kickwasm/tools/common"
)

const (
	// outputFolder = "output"
	outputFolder = ""
	yamlFilePath = "kickwasm.yaml"

	versionBreaking = 15 // Each new version breaks backwards compatibility.
	versionFeature  = 0  // Each new version adds features. Retains backwards compatibility.
	versionPatch    = 1  // Each new version only fixes bugs. No added features. Retains backwards compatibility.
)

var (
	versionDescription = []string{
		"Experimental because the go package syscall/js is still experimental.",
		"Updated to the experimental go version 1.13 syscall/js package.",
		"15.0.1 fixed 2 Type-Os in 15.0.0 inline documentation.",
		fmt.Sprintf("%d: Backwards Compatibility: Broken from the previous version %d.", versionBreaking, (versionBreaking - 1)),
		fmt.Sprintf("%d: Features: ", versionFeature),
		fmt.Sprintf("%d: Bug Patch:", versionPatch),
	}
	version = []string{
		`kickwasm:`,
		fmt.Sprintf("  Version: %d.%d.%d", versionBreaking, versionFeature, versionPatch),
		fmt.Sprint("  ", strings.Join(versionDescription, "\n  ")),
	}
	usage = `Step 1: Create a folder in your go path.
Step 2: cd to that folder.
Step 3: Create a "kickwasm.yaml" file.
Step 4: Build the framework with "kickwasm [-cc]"
`
	nlSrcBB = []byte("\n")
	nlRepBB = []byte("\\n")
	qtSrcBB = []byte("\"")
	qtRepBB = []byte("\\\"")
	ticBB   = []byte("`")
)

// type info struct {
// 	Title      string              `yaml:"title"`
// 	ImportPath string              `yaml:"importPath"`
// 	Homes      []*slurp.ButtonInfo `yaml:"buttons"`
// }

// VersionFlag means show the version.
var VersionFlag bool

// LocationsFlag means add cookie crumbs.
var LocationsFlag bool

func init() {
	flag.BoolVar(&LocationsFlag, "cc", false, "Add cookie crumbs.")
	flag.BoolVar(&VersionFlag, "v", false, "Version information.")
}

func usageFunc() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
	fmt.Fprintln(out, usage)
	flag.PrintDefaults()
}

func main() {

	var err error
	defer func() {
		if err != nil {
			os.Exit(1)
		}
	}()

	flag.Usage = usageFunc
	flag.Parse()
	if VersionFlag {
		for _, v := range version {
			fmt.Println(v)
		}
		return
	}
	// initialize paths
	var pwd string
	pwd, err = os.Getwd()
	if err != nil {
		log.Println("Tried to get the working directory but couldn't, ", err)
		return
	}
	if !common.PathFound(yamlFilePath) {
		flag.Usage()
		err = common.ErrNoKickWasm
		return
	}
	// var appPaths *paths.ApplicationPaths
	// var importPath string
	// if appPaths, importPath, err = kickwasm.Do(pwd, outputFolder, yamlFilePath, LocationsFlag, versionBreaking, versionFeature, versionPatch, pkg.LocalHost, pkg.LocalPort); err != nil {
	if _, _, err = kickwasm.Do(pwd, outputFolder, yamlFilePath, LocationsFlag, versionBreaking, versionFeature, versionPatch, pkg.LocalHost, pkg.LocalPort); err != nil {
		log.Println(err.Error())
		return
	}
}
