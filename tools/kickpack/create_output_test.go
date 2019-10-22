package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Test1 run this test first
// then run Test2
func Test1(t *testing.T) {
	var err error
	pwd, _ := os.Getwd()
	wf := filepath.Base(pwd)
	switch wf {
	case "kickwasm":
		if err = os.Chdir("tools/kickpack/testdata/deep"); err != nil {
			t.Fatal(err.Error())
		}
	case "kickpack":
		if err = os.Chdir("./testdata/deep"); err != nil {
			t.Fatal(err.Error())
		}
	case "deep":
	default:
		t.Fatal("wrong folder: " + wf)
	}
	cmd := exec.Command("kickpack", "-o", "./gooutput", "./src1", "./me.txt", "../src2")
	if err = cmd.Start(); err != nil {
		t.Fatal(err.Error())
	}
}
