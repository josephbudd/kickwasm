package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg"
	"github.com/josephbudd/kickwasm/pkg/kickwasm"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

const (
	outputFolder = "output"
)

var (
	version = []string{
		`kickwasm:`,
		`  Version: 0.2.4`,
		`  Unstable and probably buggy. 8^)`,
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
	flag.StringVar(&YAMLFileFlag, "gf", "", "The name of your main YAML file. Kickwasm will generate the framework source code using that file.")
	flag.BoolVar(&LocationsFlag, "cc", false, "Add cookie crumbs. Use with -gf.")
	flag.BoolVar(&VersionFlag, "v", false, "Version information.")
}

func main() {
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
	if filepath.Ext(YAMLFileFlag) != ".yaml" {
		log.Println("Kickwasm needs a YAML file to build the framework so the file extension must be .yaml")
	}
	// initialize paths
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Tried to get the working directory but couldn't, ", err)
		return
	}
	if _, err = kickwasm.Do(pwd, outputFolder, YAMLFileFlag, LocationsFlag, pkg.LocalHost, pkg.LocalPort, pkg.HeadTemplateFile); err != nil {
		log.Println(err.Error())
		return
	}
}
