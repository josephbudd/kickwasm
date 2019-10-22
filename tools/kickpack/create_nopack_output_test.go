package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestNoPack1 run this test first
// then run TestNoPack2
func TestNoPack1(t *testing.T) {
	var err error
	if err = os.Chdir("./testdata/deep"); err != nil {
		t.Fatal(err.Error())
	}
	cmd := exec.Command("kickpack", "-o", "./gooutputnopack", "-nopack", "./src1", "./me.txt", "../src2")
	if err = cmd.Start(); err != nil {
		t.Fatal(err.Error())
	}
}
