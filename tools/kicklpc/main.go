package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/slurp"

	"github.com/josephbudd/kickwasm/tools/common"
	"github.com/josephbudd/kickwasm/tools/kicklpc/message"
)

const (
	applicationName        = "kicklpc"
	versionBreaking        = 16 // Kicwasm Breaking Version. (Backwards compatibility.)
	versionFeature         = 0  // Added features. Still backwards compatible.
	versionPatch           = 0  // Bug fix. No added features.
	minumunKickwasmVersion = 14 // Minumum kickwasm version.
)

// VersionFlag means show the version.
var VersionFlag bool

// ListFlag means list the lpc message names.
var ListFlag bool

// AddFlag means add the named lpc message.
var AddFlag string

// DelFlag means delete the named lpc message.
var DelFlag string

func main() {
	// initialize the flags
	flag.BoolVar(&ListFlag, "l", false, "Lists the current LPC messages")
	flag.StringVar(&AddFlag, "add", "", `names the message to add`)
	flag.StringVar(&DelFlag, "delete-forever", "", `names the message to delete`)
	flag.BoolVar(&VersionFlag, "v", false, "version")
	flag.Parse()
	if VersionFlag {
		fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
		return
	}
	// Other flags are required in order to actually use kicklpc.
	if !ListFlag && len(AddFlag) == 0 && len(DelFlag) == 0 {
		help()
		return
	}
	// The user must be running this from inside the framework source code.
	var err error
	var rootFolderPath string
	if rootFolderPath, err = common.FindRoot(); err != nil {
		help()
		return
	}
	// Must not use it while rekickwasm is being used.
	if common.HaveRekickwasmFolder(rootFolderPath) {
		common.PrintRekickwasmError(applicationName)
		help()
		err = common.ErrRekickwasmExists
		return
	}
	// This framework must have been built with a recent version of kickwasm.
	if kwversion := common.AppKickwasmVersion(); kwversion < minumunKickwasmVersion {
		common.PrintWrongVersion(applicationName, kwversion, minumunKickwasmVersion)
		err = common.ErrWrongVersion
		return
	}
	// The user has provided flags for using kicklpc.
	// Prepare to handle other flags.
	infoPath := filepath.Join(rootFolderPath, ".kickwasm", "yaml", "kickwasm.yaml")
	var appInfo *slurp.ApplicationInfo
	if appInfo, err = slurp.GetApplicationInfo(infoPath); err != nil {
		fmt.Println(err.Error())
		return
	}
	appName := filepath.Base(rootFolderPath)
	authorwd := filepath.Dir(rootFolderPath)
	mngr := message.NewManager(authorwd, appName, appInfo.ImportPath)
	if ListFlag {
		var defaults, addeds []string
		if defaults, addeds, err = mngr.List(); err != nil {
			fmt.Println(err.Error())
		} else {
			// Default messages.
			var l int
			l = len(defaults)
			switch l {
			case 0:
				fmt.Println("No default messages.")
			case 1:
				fmt.Printf("%d default message:\n", l)
				for _, msg := range defaults {
					fmt.Println("\t", msg)
				}
			default:
				fmt.Printf("%d default messages:\n", l)
				for _, msg := range defaults {
					fmt.Println("\t", msg)
				}
			}
			// Added messages.
			l = len(addeds)
			switch l {
			case 0:
				fmt.Println("No added messages.")
			case 1:
				fmt.Printf("%d added message:\n", l)
				for _, msg := range addeds {
					fmt.Println("\t", msg)
				}
			default:
				fmt.Printf("%d added messages:\n", l)
				for _, msg := range addeds {
					fmt.Println("\t", msg)
				}
			}
		}
		return
	}
	if len(AddFlag) > 0 {
		if err = mngr.Add(AddFlag); err != nil {
			fmt.Println(err.Error())
		}
	}
	if len(DelFlag) > 0 {
		if err = mngr.Del(DelFlag); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func help() {
	fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
	fmt.Println(common.UseItAnyWhere)
	flag.Usage()
}
