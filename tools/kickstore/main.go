package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/slurp"

	"github.com/josephbudd/kickwasm/tools/common"
	"github.com/josephbudd/kickwasm/tools/kickstore/store"
)

const (
	applicationName        = "kickstore"
	versionBreaking        = 16 // Kicwasm Breaking Version. (Backwards compatibility.)
	versionFeature         = 0  // Added features. Still backwards compatible.
	versionPatch           = 0  // Bug fix. No added features.
	minumunKickwasmVersion = 14 // Minumum kickwasm version.
)

// VersionFlag means show the version.
var VersionFlag bool

// ListFlag means list the store names.
var ListFlag bool

// AddFlag means add the named store.
var AddFlag string

// DelFlag means delete the named store.
var DelFlag string

// AddRemoteDBFlag means add the named store.
var AddRemoteDBFlag string

// DelRemoteDBFlag means delete the named store.
var DelRemoteDBFlag string

// AddRemoteRecordFlag means add the named store.
var AddRemoteRecordFlag string

// DelRemoteRecordFlag means delete the named store.
var DelRemoteRecordFlag string

func main() {

	var err error
	defer func() {
		if err != nil {
			os.Exit(1)
		}
	}()

	// initialize the flags
	flag.BoolVar(&ListFlag, "l", false, "Lists the current stores")
	flag.StringVar(&AddFlag, "add", "", `names the local bolt store to add`)
	flag.StringVar(&DelFlag, "delete-forever", "", `names the local bolt store to delete`)
	flag.StringVar(&AddRemoteDBFlag, "add-remote-api", "", `names the remote API to add`)
	flag.StringVar(&DelRemoteDBFlag, "delete-forever-remote-api", "", `names the remote API to delete`)
	flag.StringVar(&AddRemoteRecordFlag, "add-remote-record", "", `names the remote record to add`)
	flag.StringVar(&DelRemoteRecordFlag, "delete-forever-remote-record", "", `names the remote record to delete`)
	flag.BoolVar(&VersionFlag, "v", false, "version")
	flag.Parse()
	// The version and help flags work alone.
	// If they are used then process them and quit.
	if VersionFlag {
		fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
		return
	}
	// Other flags are required in order to actually use kickstore.
	if !ListFlag && len(AddFlag) == 0 && len(DelFlag) == 0 && len(AddRemoteDBFlag) == 0 && len(DelRemoteDBFlag) == 0 && len(AddRemoteRecordFlag) == 0 && len(DelRemoteRecordFlag) == 0 {
		help()
		return
	}
	// The user must be running this from inside the framework source code.
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
	var appInfo *slurp.ApplicationInfo
	if appInfo, err = common.ApplicationInfo(rootFolderPath); err != nil {
		fmt.Println(err.Error())
		return
	}
	appName := filepath.Base(rootFolderPath)
	authorwd := filepath.Dir(rootFolderPath)
	var mngr *store.Manager
	if mngr, err = store.NewManager(authorwd, appName, appInfo.ImportPath); err != nil {
		fmt.Println(err.Error())
		return
	}
	if ListFlag {
		boltStores, remoteDBs, remoteRecords := mngr.List()
		fmt.Printf("%d Local Bolt Stores:\n", len(boltStores))
		for i, s := range boltStores {
			fmt.Printf(" %2d. %s\n", i+1, s)
		}
		fmt.Printf("%d Remote APIs:\n", len(remoteDBs))
		for i, s := range remoteDBs {
			fmt.Printf(" %2d. %s\n", i+1, s)
		}
		fmt.Printf("%d Remote Records:\n", len(remoteRecords))
		for i, s := range remoteRecords {
			fmt.Printf(" %2d. %s\n", i+1, s)
		}
		return
	}
	if len(AddFlag) > 0 {
		if err = mngr.Add(AddFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(DelFlag) > 0 {
		if err = mngr.Del(DelFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(AddRemoteDBFlag) > 0 {
		if err = mngr.AddRemoteDB(AddRemoteDBFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(DelRemoteDBFlag) > 0 {
		if err = mngr.DelRemoteDB(DelRemoteDBFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(AddRemoteRecordFlag) > 0 {
		if err = mngr.AddRemoteRecord(AddRemoteRecordFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if len(DelRemoteRecordFlag) > 0 {
		if err = mngr.DelRemoteRecord(DelRemoteRecordFlag); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if err = mngr.Finish(); err != nil {
		fmt.Println(err.Error())
	}
}

func help() {
	fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
	fmt.Println(common.UseItAnyWhere)
	flag.Usage()
}
