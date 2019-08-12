package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg"
	"github.com/josephbudd/kickwasm/pkg/kickwasm"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

const (
	outputFolder = "output"

	versionBreaking = 6 // Each new version breaks backwards compatibility.
	versionFeature  = 0 // Each new version adds features. Retains backwards compatibility.
	versionPatch    = 1 // Each new version only fixes bugs. No added features. Retains backwards compatibility.

)

var (
	versionDescription = []string{
		"Experimental because the go package syscall/js is still experimental.",
		"Updated to the experimental go version 1.12 syscall/js package.",
		"6: Backwards Compatibility: Broken from version 5.",
		"0: ",
		"1: Minor changes for the go report card.",
	}
	version = []string{
		`kickwasm:`,
		fmt.Sprintf("  Version: %d.%d.%d", versionBreaking, versionFeature, versionPatch),
		fmt.Sprint("  ", strings.Join(versionDescription, "\n  ")),
	}
	nlSrcBB = []byte("\n")
	nlRepBB = []byte("\\n")
	qtSrcBB = []byte("\"")
	qtRepBB = []byte("\\\"")
	ticBB   = []byte("`")
)

type info struct {
	Title      string               `yaml:"title"`
	ImportPath string               `yaml:"importPath"`
	Stores     []string             `yaml:"stores"`
	Services   []*slurp.ServiceInfo `yaml:"services"`
}

//YAMLFileFlag is the file.
var YAMLFileFlag string

// VersionFlag means show the version.
var VersionFlag bool

// LocationsFlag means add cookie crumbs.
var LocationsFlag bool

func init() {
	flag.StringVar(&YAMLFileFlag, "f", "", "The path to the kickwasm.yaml or kickwasm.yml file.")
	flag.BoolVar(&LocationsFlag, "cc", false, "Add cookie crumbs. Use with -f.")
	flag.BoolVar(&VersionFlag, "v", false, "Version information.")
}

func main() {
	fileNames := paths.GetFileNames()
	flag.Parse()
	if VersionFlag {
		for _, v := range version {
			fmt.Println(v)
		}
		return
	}
	if len(YAMLFileFlag) == 0 {
		flag.PrintDefaults()
		return
	}
	if filename := filepath.Base(YAMLFileFlag); filename != fileNames.KickwasmDotYAML && filename != fileNames.KickwasmDotYML {
		log.Printf("Kickwasm needs a YAML file named %s or %s to build the framework not a file named %q", fileNames.KickwasmDotYAML, fileNames.KickwasmDotYML, filename)
		return
	}
	// initialize paths
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Tried to get the working directory but couldn't, ", err)
		return
	}
	if _, err = kickwasm.Do(pwd, outputFolder, YAMLFileFlag, LocationsFlag, versionBreaking, versionFeature, versionPatch, pkg.LocalHost, pkg.LocalPort); err != nil {
		log.Println(err.Error())
		return
	}
}
