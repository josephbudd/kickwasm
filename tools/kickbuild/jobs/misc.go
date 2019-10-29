package jobs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	success = " * SUCCESS"
	newline = ""
)

var step uint

func printNextStep(message string) {
	step++
	PrintLine(fmt.Sprintf("STEP %d: %s", step, message))
}

// PrintLine prints a line.
func PrintLine(text string) {
	fmt.Println("kickbuild: ", text)
}

func printError(err error) {
	PrintLine("Error: " + err.Error())
	PrintLine(newline)
}

func printSuccess() {
	PrintLine(success)
	PrintLine(newline)
}

func printDump(dump []byte, err error) {
	if err == nil {
		return
	}
	PrintLine("Error: " + err.Error())
	fmt.Fprintln(os.Stderr, string(dump))
}

func fixPrintDump(rootFolderPath string, dump []byte, err error) {
	if err == nil {
		return
	}
	PrintLine("Error: " + err.Error())
	lines := strings.Split(string(dump), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if matched, err := regexp.MatchString(`.*\.go\:\d+\:\d+\:\s+`, line); matched == true && err == nil {
			fmt.Fprintln(os.Stderr, filepath.Join(rootFolderPath, line))
		} else {
			fmt.Fprintln(os.Stderr, strings.TrimSpace(line))
		}
	}
}
