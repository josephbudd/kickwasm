package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/common"
	"github.com/josephbudd/kickwasm/tools/script"
)

var testFolders = []string{

	// panels from tabs to buttons and buttons to tabs.
	"tp2b",

	// initial button pad.
	"initaddtabs",
	"initrot",
	"initswap",

	// buttons.
	"bprot",
	"bprotpan",
	"bpswap",
	"bpswappan",

	// tabs.
	"tbrot",
	"tbrotpan",
	"tbswap",
	"tbswappan",

	// spawn tabs.
	"stbrot",
	"stbrotpan",
	"stbswap",
	"stbswappan",

	// tabs and spawn tabs.
	"tbstbrot",
}

func main() {

	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

	var wd string
	if wd, err = os.Getwd(); err != nil {
		return
	}
	if filepath.Base(wd) == "kickwasm" {
		// fix if funning from the kickwasm folder.
		wd = filepath.Join(wd, "proofs")
	}
	fmt.Println("PROOFS.")
	fmt.Println("A SERIES OF REKICKWASM TESTS PERFORMED IN RUNNING FRAMEWORKS.")
	for _, f := range testFolders {
		path := filepath.Join(wd, f)
		if _, err = script.RunDump(nil, path, "go", "build"); err != nil {
			return
		}
		executable := common.Executable("." + string(os.PathSeparator) + f)
		var dump []byte
		if dump, err = script.RunDump(nil, path, executable); err != nil {
			return
		}
		fmt.Println(string(dump))
	}
	fmt.Println("SUCCESS.")
	fmt.Println("Each refactoring with rekickwasm worked successfully.")
	fmt.Println("End of Proofs.")
}
