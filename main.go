package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/flagdata"
	"github.com/josephbudd/kickwasm/foldercp"
	"github.com/josephbudd/kickwasm/mainprocess"
	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer"
	"github.com/josephbudd/kickwasm/slurp"
	"github.com/josephbudd/kickwasm/tap"
)

const (
	headTemplateFile = "head.tmpl"
	outputFolder     = "output"
	host             = "127.0.0.1"
	port             = uint(9090)
)

var (
	version = []string{
		`kickwasm:`,
		`  Version: 0.1.0`,
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
	Repos      []string             `yaml:"repos"`
	Services   []*slurp.ServiceInfo `yaml:"services"`
}

//FileFlag is the file.
var FileFlag string

// AboutFlag means add the about section.
var AboutFlag bool

// VersionFlag means show the version.
var VersionFlag bool

// LocationsFlag means add cookie crumbs.
var LocationsFlag bool

func init() {
	flag.StringVar(&FileFlag, "gf", "", "The name of your application yaml file. Kick will generate source code using that file.")
	flag.BoolVar(&AboutFlag, "about", false, "Kick will add it's default About section to the generated source code. Use with -gf or -gx.")
	flag.BoolVar(&VersionFlag, "v", false, "Kick will display it's version information.")
	flag.BoolVar(&LocationsFlag, "cc", false, "Kick will add cookie crumbs to the generated source code. Use with -gf or -gx.")
}

func main() {
	flag.Parse()
	if VersionFlag {
		for _, v := range version {
			fmt.Println(v)
		}
		return
	}
	if len(FileFlag) == 0 {
		flag.PrintDefaults()
		return
	}
	if filepath.Ext(FileFlag) != ".yaml" {
		log.Println("Kick needs a YAML file to build your application so the file extension must be .yaml")
	}
	// initialize paths
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Tried to get the working directory but couldn't, ", err)
		return
	}
	sl := slurp.NewSlurper()
	builder, err := sl.Gulp(FileFlag)
	if err != nil {
		log.Println("Tried to slurp the yaml file(s) but counldn't, ", err)
		return
	}
	parts := strings.Split(builder.ImportPath, "/")
	appName := parts[len(parts)-1]
	appPaths := &paths.ApplicationPaths{}
	appPaths.Initialize(pwd, "output", appName)
	if err = appPaths.MakeOutput(); err != nil {
		log.Println(err)
		return
	}
	if err := create(appPaths, builder); err != nil {
		log.Println(err)
		return
	}
	folderPaths := appPaths.GetPaths()
	// create the .kick/flags.yaml file
	flagsPath := filepath.Join(folderPaths.OutputDotKick, "flags.yaml")
	yamlStartFileName := filepath.Base(FileFlag)
	if err := flagdata.SaveFlags(flagsPath, AboutFlag, LocationsFlag, yamlStartFileName); err != nil {
		log.Println(err)
		return
	}
	// build the panel file paths
	appYAMLFilePath := filepath.Join(pwd, FileFlag)
	panelFilePaths := sl.GetPanelFilePaths()
	for i, p := range panelFilePaths {
		panelFilePaths[i] = filepath.Join(pwd, p)
	}
	if err := foldercp.CopyYAML(appPaths, appYAMLFilePath, panelFilePaths); err != nil {
		log.Println(err)
		return
	}
}

func create(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	// get the application name from the import path.
	if err := renderer.Create(appPaths, builder, AboutFlag, LocationsFlag, headTemplateFile, host, port); err != nil {
		return err
	}
	if err := mainprocess.Create(appPaths, builder, AboutFlag, headTemplateFile, host, port); err != nil {
		return err
	}
	return nil
}
