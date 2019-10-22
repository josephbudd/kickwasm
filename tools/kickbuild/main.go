package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/tools/common"
	"github.com/josephbudd/kickwasm/tools/kickbuild/jobs"
)

const (
	success = "* SUCCESS"
)

var (
	initError      error
	startwd        string // renderer
	rootFolderPath string
	appFolderName  string
	authorwd       string

	relativeSitePath string
	fullSitePath     string
	rendererPath     string
	fullRendererPath string

	sitepackPackageName string
	sitepackPackagePath string

	spawnpackPackageName string
	spawnpackPackagePath string

	spawnTemplatesFolder string
)

const (
	applicationName        = "kickbuild"
	versionBreaking        = 11 // Kicwasm Breaking Version. (Backwards compatibility.)
	versionFeature         = 0  // Added features. Still backwards compatible.
	versionPatch           = 0  // Bug fix. No added features.
	minumunKickwasmVersion = 11 // Minumum kickwasm version.
	here                   = "."
)

// VersionFlag means show the version.
var VersionFlag bool

// RendererFlag means only build the renderer process.
var RendererFlag bool

// MainFlag means only build the main process.
var MainFlag bool

// PackFlag means pack the renderer process into a package.
var PackFlag bool

//GOOSFlag is the build GOOS=linux GOARCH flag.
var GOOSFlag string

// GOARCHFlag is the go build GOARCH flag.
var GOARCHFlag string

// RunFlag means run the executable.
var RunFlag bool

func main() {

	// flags
	flag.BoolVar(&VersionFlag, "v", false, "display the version")
	flag.BoolVar(&RendererFlag, "rp", false, "quick build the renderer process")
	flag.BoolVar(&MainFlag, "mp", false, "build the main process")
	flag.BoolVar(&PackFlag, "packrp", false, "pack and build the renderer process")
	flag.BoolVar(&RunFlag, "run", false, "run the exectuable")
	flag.StringVar(&GOOSFlag, "GOOS", "", "pass GOOS=<string> to go build. Use with -mp")
	flag.StringVar(&GOARCHFlag, "GOARCH", "", "pass GOARCH=<string> to go build. Use with -mp")
	flag.Parse()

	if VersionFlag {
		// Handle the version flag.
		fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
		return
	}
	// The user must be running this from inside the framework source code.
	var err error
	if rootFolderPath, err = common.FindRoot(); err != nil {
		help()
		return
	}

	// Must not use it while rekickwasm is being used.
	if common.HaveRekickwasmFolder(rootFolderPath) {
		common.PrintRekickwasmError(applicationName)
		help()
		return
	}

	// This framework must have been built with a recent version of kickwasm.
	if kwversion := common.AppKickwasmVersion(); kwversion < minumunKickwasmVersion {
		common.PrintWrongVersion(applicationName, kwversion, minumunKickwasmVersion)
		return
	}
	folderNames := paths.GetFolderNames()
	appFolderName = filepath.Base(rootFolderPath)
	authorwd = filepath.Dir(rootFolderPath)
	relativeSitePath = folderNames.RendererSite
	fullSitePath = filepath.Join(rootFolderPath, folderNames.RendererSite)
	rendererPath = folderNames.Renderer
	fullRendererPath = filepath.Join(rootFolderPath, folderNames.Renderer)
	sitepackPackageName = appFolderName + folderNames.SitePack
	sitepackPackagePath = filepath.Join(authorwd, sitepackPackageName)
	spawnpackPackageName = folderNames.SpawnPack
	spawnpackPackagePath = filepath.Join(fullRendererPath, spawnpackPackageName)
	spawnTemplatesFolder = folderNames.SpawnTemplates

	// Process the remaining flags.
	if PackFlag {
		RendererFlag = true
	}
	if !RendererFlag && !MainFlag && !RunFlag {
		help()
		return
	}
	if PackFlag && !RendererFlag {
		help()
	}
	if len(GOOSFlag) > 0 && !MainFlag {
		help()
		return
	}
	if len(GOARCHFlag) > 0 && !MainFlag {
		help()
		return
	}
	if MainFlag {
		if len(GOOSFlag) == 0 {
			GOOSFlag = runtime.GOOS
		}
		if len(GOARCHFlag) == 0 {
			GOARCHFlag = runtime.GOARCH
		}
	}
	if RendererFlag {
		if err = buildRendererProcess(); err != nil {
			return
		}
	}
	if MainFlag {
		if err = buildMainProcess(); err != nil {
			return
		}
	}
	if RunFlag {
		if err = jobs.RunExecutable(rootFolderPath, appFolderName, GOOSFlag); err != nil {
			return
		}
	}
}

func buildMainProcess() (err error) {
	jobs.PrintLine("BUILD THE MAIN PROCESS.")
	jobs.PrintLine("")
	err = jobs.BuildMain(rootFolderPath, appFolderName, GOOSFlag, GOARCHFlag)
	return
}

func buildRendererProcess() (err error) {
	jobs.PrintLine("BUILD THE RENDERER PROCESS.")
	jobs.PrintLine("")
	// Remove the old spawn package if it's there.
	if err = jobs.RemoveOldSpawnPackPackage(spawnpackPackageName, spawnpackPackagePath); err != nil {
		return
	}
	// Write the new spawn package.
	if err = jobs.WriteNewSpawnPackage(spawnpackPackageName, spawnpackPackagePath, spawnTemplatesFolder, fullSitePath); err != nil {
		return
	}
	// Build the wasm.
	if err = jobs.BuildWASM(rootFolderPath, rendererPath); err != nil {
		return
	}
	// Remove the old site pack package.
	if err = jobs.RemoveOldSitePackPackage(sitepackPackageName, sitepackPackagePath); err != nil {
		return
	}
	// Write the new site pack package.
	if PackFlag {
		if err = jobs.WriteSitePackPackagePack(rootFolderPath, sitepackPackageName, sitepackPackagePath); err != nil {
			return
		}
	} else {
		if err = jobs.WriteSitePackPackageDontPack(rootFolderPath, sitepackPackageName, sitepackPackagePath); err != nil {
			return
		}
	}

	if err = jobs.BuildSitePackPackage(sitepackPackageName, sitepackPackagePath); err != nil {
		return
	}
	return
}

func help() {
	fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
	fmt.Println(common.UseItAnyWhere)
	flag.Usage()
}
