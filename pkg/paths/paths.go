package paths

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Import paths
const (

	// domain data

	importDomainDataFilepaths = "/domain/data/filepaths"
	importDomainDataLogLevels = "/domain/data/loglevels"
	importDomainDataSettings  = "/domain/data/settings"

	// domain lpc

	importDomainLPC        = "/domain/lpc"
	importDomainLPCMessage = "/domain/lpc/message"

	// main process

	importMainProcessHomes = "/mainprocess/homes"

	// main process lpc

	importMainProcessLPC         = "/mainprocess/lpc"
	importMainProcessLPCDispatch = "/mainprocess/lpc/dispatch"

	// renderer

	importRenderer         = "/rendererprocess"
	importRendererPanels   = "/rendererprocess/panels"
	importRendererPaneling = "/rendererprocess/paneling"

	// renderer site

	importRendererSpawnPanels = "/rendererprocess/spawnPanels"

	// stores

	importDomainStore        = "/domain/store"
	importDomainStoreRecord  = "/domain/store/record"
	importDomainStoreStorer  = "/domain/store/storer"
	importDomainStoreStoring = "/domain/store/storing"

	// v 13

	importRendererDOM         = "/rendererprocess/dom"
	importRendererMarkup      = "/rendererprocess/markup"
	importRendererWindow      = "/rendererprocess/window"
	importRendererEvent       = "/rendererprocess/event"
	importRendererDisplay     = "/rendererprocess/display"
	importRendererApplication = "/rendererprocess/application"

	// renderer framework

	importRendererFramework = "/rendererprocess/framework"
	importRendererCallBack  = "/rendererprocess/framework/callback"
	importRendererLocation  = "/rendererprocess/framework/location"
	importRendererLPC       = "/rendererprocess/framework/lpc"
	importRendererSpawnPack = "/rendererprocess/framework/spawnpack"
	importRendererViewTools = "/rendererprocess/framework/viewtools"
	importRendererProofs    = "/rendererprocess/framework/proofs"
)

// Imports is the import paths.
type Imports struct {

	// domain data

	ImportDomainDataFilepaths string
	ImportDomainDataLogLevels string
	ImportDomainDataSettings  string

	// lpc

	ImportDomainLPC              string
	ImportDomainLPCMessage       string
	ImportRendererLPC            string
	ImportMainProcessLPC         string
	ImportMainProcessLPCDispatch string

	// main process

	ImportMainProcessHomes string

	// renderer

	ImportRenderer          string
	ImportRendererPanels    string
	ImportRendererViewTools string
	ImportRendererPaneling  string

	ImportRendererSpawnPack   string
	ImportRendererSpawnPanels string

	// store

	ImportDomainStore        string
	ImportDomainStoreRecord  string
	ImportDomainStoreStorer  string
	ImportDomainStoreStoring string

	// v 13

	ImportRendererDOM         string
	ImportRendererCallBack    string
	ImportRendererLocation    string
	ImportRendererMarkup      string
	ImportRendererWindow      string
	ImportRendererEvent       string
	ImportRendererDisplay     string
	ImportRendererProofs      string
	ImportRendererApplication string

	// renderer framework

	ImportRendererFramework string
}

// GetImports returns the go import paths.
func GetImports() *Imports {
	return &Imports{

		// domain data

		ImportDomainDataFilepaths: importDomainDataFilepaths,
		ImportDomainDataLogLevels: importDomainDataLogLevels,
		ImportDomainDataSettings:  importDomainDataSettings,

		// lpc

		ImportDomainLPC:              importDomainLPC,
		ImportDomainLPCMessage:       importDomainLPCMessage,
		ImportRendererLPC:            importRendererLPC,
		ImportMainProcessLPC:         importMainProcessLPC,
		ImportMainProcessLPCDispatch: importMainProcessLPCDispatch,

		// main process

		ImportMainProcessHomes: importMainProcessHomes,

		// renderer

		ImportRenderer:          importRenderer,
		ImportRendererPanels:    importRendererPanels,
		ImportRendererViewTools: importRendererViewTools,
		ImportRendererPaneling:  importRendererPaneling,

		ImportRendererSpawnPack:   importRendererSpawnPack,
		ImportRendererSpawnPanels: importRendererSpawnPanels,

		// store

		ImportDomainStore:        importDomainStore,
		ImportDomainStoreRecord:  importDomainStoreRecord,
		ImportDomainStoreStorer:  importDomainStoreStorer,
		ImportDomainStoreStoring: importDomainStoreStoring,

		// v 13

		ImportRendererDOM:         importRendererDOM,
		ImportRendererCallBack:    importRendererCallBack,
		ImportRendererLocation:    importRendererLocation,
		ImportRendererMarkup:      importRendererMarkup,
		ImportRendererWindow:      importRendererWindow,
		ImportRendererEvent:       importRendererEvent,
		ImportRendererDisplay:     importRendererDisplay,
		ImportRendererProofs:      importRendererProofs,
		ImportRendererApplication: importRendererApplication,

		// renderer framework

		ImportRendererFramework: importRendererFramework,
	}
}

// ApplicationPathsI is a test
type ApplicationPathsI interface {
	GetPaths() *Paths
	Copy(dest, src string) error
	WriteFile(fpath string, data []byte) error
	GetDMode() os.FileMode
	GetFMode() os.FileMode
	GetFileNames() *FileNames
	GetFolderNames() *FolderNames
}

// ApplicationPaths is the paths needed for kick.
type ApplicationPaths struct {
	appName string

	inputFolderRendererPath    string
	inputFolderMainProcessPath string

	outputFolderRendererPath    string
	outputFolderMainProcessPath string

	// FMode is the applications mode for files.
	FMode os.FileMode

	// DMode is the applications mode for folders.
	DMode os.FileMode

	initerr error

	paths Paths
}

// Initialize initializes the paths.
func (ap *ApplicationPaths) Initialize(pwd, outputFolder, appname string) {
	ap.FMode = os.FileMode(0666)
	ap.DMode = os.FileMode(0775)
	ap.appName = appname

	ap.initializeOutput(pwd, outputFolder, appname)

	// import paths
	ap.paths.ImportDomainDataFilepaths = importDomainDataFilepaths
	ap.paths.ImportDomainDataLogLevels = importDomainDataLogLevels
	ap.paths.ImportDomainDataSettings = importDomainDataSettings

	// lpc

	ap.paths.ImportDomainLPC = importDomainLPC
	ap.paths.ImportDomainLPCMessage = importDomainLPCMessage
	ap.paths.ImportRendererLPC = importRendererLPC
	ap.paths.ImportMainProcessLPC = importMainProcessLPC
	ap.paths.ImportMainProcessLPCDispatch = importMainProcessLPCDispatch

	ap.paths.ImportMainProcessHomes = importMainProcessHomes
	ap.paths.ImportMainProcessLPCDispatch = importMainProcessLPCDispatch

	ap.paths.ImportRenderer = importRenderer
	ap.paths.ImportRendererPanels = importRendererPanels
	ap.paths.ImportRendererViewTools = importRendererViewTools
	ap.paths.ImportRendererPaneling = importRendererPaneling

	ap.paths.ImportRendererSpawnPack = importRendererSpawnPack
	ap.paths.ImportRendererSpawnPanels = importRendererSpawnPanels

	ap.paths.ImportDomainStore = importDomainStore
	ap.paths.ImportDomainStoreRecord = importDomainStoreRecord
	ap.paths.ImportDomainStoreStorer = importDomainStoreStorer
	ap.paths.ImportDomainStoreStoring = importDomainStoreStoring

	// v 13

	ap.paths.ImportRendererDOM = importRendererDOM
	ap.paths.ImportRendererCallBack = importRendererCallBack
	ap.paths.ImportRendererLocation = importRendererLocation
	ap.paths.ImportRendererMarkup = importRendererMarkup
	ap.paths.ImportRendererWindow = importRendererWindow
	ap.paths.ImportRendererEvent = importRendererEvent
	ap.paths.ImportRendererDisplay = importRendererDisplay
	ap.paths.ImportRendererProofs = importRendererProofs
	ap.paths.ImportRendererApplication = importRendererApplication

	// framework

	ap.paths.ImportRendererFramework = importRendererFramework
}

// GetDMode returns the file mode for directories.
func (ap *ApplicationPaths) GetDMode() os.FileMode {
	return ap.DMode
}

// GetFMode returns the file mode for files.
func (ap *ApplicationPaths) GetFMode() os.FileMode {
	return ap.FMode
}

// GetPaths returns a struct of all the paths.
func (ap *ApplicationPaths) GetPaths() *Paths {
	paths := ap.paths
	return &paths
}

// GetFileNames returns a struct of all the file names.
func (ap *ApplicationPaths) GetFileNames() *FileNames {
	return GetFileNames()
}

// GetFolderNames returns the folder names.
func (ap *ApplicationPaths) GetFolderNames() (fnames *FolderNames) {
	fnames = GetFolderNames()
	fnames.SitePack = strings.ToLower(ap.appName) + "sitepack"
	return
}

// Paths are the folder paths
type Paths struct {
	Output                 string
	OutputDotKickwasm      string
	OutputDotKickwasmYAML  string
	OutputDotKickwasmFlags string

	OutputDotKickstore string

	// output sitepack
	OutputSitePack string

	// output domain

	OutputDomain string

	OutputDomainData          string
	OutputDomainDataFilepaths string
	OutputDomainDataLogLevels string
	OutputDomainDataSettings  string

	OutputDomainStore        string
	OutputDomainStoreStoring string
	OutputDomainStoreStorer  string
	OutputDomainStoreRecord  string

	// output domain lpc

	OutputDomainLPC        string
	OutputDomainLPCMessage string

	// output main process

	OutputMainProcess            string
	OutputMainProcessLPCDispatch string

	// output main process lpc

	OutputMainProcessLPC string

	// output renderer

	OutputRenderer               string
	OutputRendererCSS            string
	OutputRendererMyCSS          string
	OutputRendererTemplates      string
	OutputRendererSpawnTemplates string
	OutputRendererPanels         string
	OutputRendererSpawns         string
	OutputRendererSite           string
	OutputRendererViewTools      string
	OutputRendererPaneling       string
	OutputRendererSpawnPack      string

	// output renderer lpc

	OutputRendererLPC string

	// v 13

	OutputRendererDOM         string
	OutputRendererCallBack    string
	OutputRendererLocation    string
	OutputRendererMarkup      string
	OutputRendererWindow      string
	OutputRendererEvent       string
	OutputRendererDisplay     string
	OutputRendererProofs      string
	OutputRendererApplication string

	// framework

	OutputRendererFramework string

	// import domain

	ImportDomainDataFilepaths string
	ImportDomainDataLogLevels string
	ImportDomainDataSettings  string

	// import domain lpc

	ImportDomainLPC        string
	ImportDomainLPCMessage string

	// import main process

	ImportMainProcessHomes string

	// import main process lpc

	ImportMainProcessLPC         string
	ImportMainProcessLPCDispatch string

	// import renderer

	ImportRenderer          string
	ImportRendererPanels    string
	ImportRendererViewTools string
	ImportRendererPaneling  string

	ImportRendererSpawnPack   string
	ImportRendererSpawnPanels string

	// import renderer lpc

	ImportRendererLPC string

	// import store

	ImportDomainStore        string
	ImportDomainStoreRecord  string
	ImportDomainStoreStorer  string
	ImportDomainStoreStoring string

	// v 13

	ImportRendererDOM         string
	ImportRendererCallBack    string
	ImportRendererLocation    string
	ImportRendererMarkup      string
	ImportRendererWindow      string
	ImportRendererEvent       string
	ImportRendererDisplay     string
	ImportRendererProofs      string
	ImportRendererApplication string

	// renderer framework

	ImportRendererFramework string
}

// initializeOutput defines the output paths
func (ap *ApplicationPaths) initializeOutput(pwd, outputFolder, appname string) {
	buildingInCurrentFolder := (len(outputFolder) == 0 && filepath.Base(pwd) == appname)
	fileNames := GetFileNames()
	folderNames := ap.GetFolderNames()

	// set and create the output folder and sub folders.
	// fix the output folder.
	if buildingInCurrentFolder {
		ap.paths.Output = pwd
		ap.paths.OutputSitePack = filepath.Join(filepath.Dir(pwd), folderNames.SitePack)
	} else {
		ap.paths.Output = filepath.Join(pwd, outputFolder, appname)
		ap.paths.OutputSitePack = filepath.Join(pwd, outputFolder, folderNames.SitePack)
	}
	// output .kickwasm folder and sub folders
	ap.paths.OutputDotKickwasm = filepath.Join(ap.paths.Output, folderNames.DotKickwasm)
	ap.paths.OutputDotKickwasmYAML = filepath.Join(ap.paths.OutputDotKickwasm, folderNames.YAML)
	ap.paths.OutputDotKickwasmFlags = filepath.Join(ap.paths.OutputDotKickwasm, fileNames.FlagDotYAML)

	// output .kickstore folder and sub folder
	ap.paths.OutputDotKickstore = filepath.Join(ap.paths.Output, folderNames.DotKickstore)

	// output domain folder and sub folders.
	ap.paths.OutputDomain = filepath.Join(ap.paths.Output, folderNames.Domain)
	ap.paths.OutputDomainData = filepath.Join(ap.paths.OutputDomain, folderNames.Data)
	ap.paths.OutputDomainDataFilepaths = filepath.Join(ap.paths.OutputDomainData, folderNames.FilePaths)
	ap.paths.OutputDomainDataLogLevels = filepath.Join(ap.paths.OutputDomainData, folderNames.LogLevels)
	ap.paths.OutputDomainDataSettings = filepath.Join(ap.paths.OutputDomainData, folderNames.Settings)
	ap.paths.OutputDomainStore = filepath.Join(ap.paths.OutputDomain, folderNames.Store)
	ap.paths.OutputDomainStoreStoring = filepath.Join(ap.paths.OutputDomainStore, folderNames.Storing)
	ap.paths.OutputDomainStoreStorer = filepath.Join(ap.paths.OutputDomainStore, folderNames.Storer)
	ap.paths.OutputDomainStoreRecord = filepath.Join(ap.paths.OutputDomainStore, folderNames.Record)

	// output renderer folder and sub folders.
	ap.paths.OutputRenderer = filepath.Join(ap.paths.Output, folderNames.Renderer)
	ap.paths.OutputRendererPanels = filepath.Join(ap.paths.OutputRenderer, folderNames.Panels)
	ap.paths.OutputRendererSpawns = filepath.Join(ap.paths.OutputRenderer, folderNames.SpawnPanels)
	ap.paths.OutputRendererPaneling = filepath.Join(ap.paths.OutputRenderer, folderNames.Paneling)

	// output renderer site folders
	ap.paths.OutputRendererSite = filepath.Join(ap.paths.Output, folderNames.RendererSite)
	ap.paths.OutputRendererCSS = filepath.Join(ap.paths.OutputRendererSite, folderNames.CSS)
	ap.paths.OutputRendererMyCSS = filepath.Join(ap.paths.OutputRendererSite, folderNames.MyCSS)
	ap.paths.OutputRendererTemplates = filepath.Join(ap.paths.OutputRendererSite, folderNames.Templates)
	ap.paths.OutputRendererSpawnTemplates = filepath.Join(ap.paths.OutputRendererSite, folderNames.SpawnTemplates)

	// v 13

	ap.paths.OutputRendererDisplay = filepath.Join(ap.paths.OutputRenderer, folderNames.Display)
	ap.paths.OutputRendererDOM = filepath.Join(ap.paths.OutputRenderer, folderNames.DOM)
	ap.paths.OutputRendererMarkup = filepath.Join(ap.paths.OutputRenderer, folderNames.Markup)
	ap.paths.OutputRendererEvent = filepath.Join(ap.paths.OutputRenderer, folderNames.Event)
	ap.paths.OutputRendererWindow = filepath.Join(ap.paths.OutputRenderer, folderNames.Window)
	ap.paths.OutputRendererApplication = filepath.Join(ap.paths.OutputRenderer, folderNames.Application)

	// output renderer framework

	ap.paths.OutputRendererFramework = filepath.Join(ap.paths.OutputRenderer, folderNames.Framework)
	ap.paths.OutputRendererViewTools = filepath.Join(ap.paths.OutputRendererFramework, folderNames.ViewTools)
	ap.paths.OutputRendererSpawnPack = filepath.Join(ap.paths.OutputRendererFramework, folderNames.SpawnPack)
	ap.paths.OutputRendererCallBack = filepath.Join(ap.paths.OutputRendererFramework, folderNames.CallBack)
	ap.paths.OutputRendererLocation = filepath.Join(ap.paths.OutputRendererFramework, folderNames.Location)
	ap.paths.OutputRendererProofs = filepath.Join(ap.paths.OutputRendererFramework, folderNames.Proofs)
	// ap.paths.OutputRendererLPC is below

	// output mainprocess folder and sub folders.
	ap.paths.OutputMainProcess = filepath.Join(ap.paths.Output, folderNames.MainProcess)

	// output lpc
	ap.paths.OutputDomainLPC = filepath.Join(ap.paths.OutputDomain, folderNames.LPC)
	ap.paths.OutputDomainLPCMessage = filepath.Join(ap.paths.OutputDomainLPC, folderNames.Message)
	ap.paths.OutputRendererLPC = filepath.Join(ap.paths.OutputRendererFramework, folderNames.LPC)
	ap.paths.OutputMainProcessLPC = filepath.Join(ap.paths.OutputMainProcess, folderNames.LPC)
	ap.paths.OutputMainProcessLPCDispatch = filepath.Join(ap.paths.OutputMainProcessLPC, folderNames.Dispatch)
}

// MakeOutput creates the output paths
func (ap *ApplicationPaths) MakeOutput() (err error) {
	// output .kickwasm folder and sub folders
	if err = os.MkdirAll(ap.paths.OutputDotKickwasmYAML, ap.DMode); err != nil {
		return
	}
	// output .kickstore folder
	if err = os.MkdirAll(ap.paths.OutputDotKickstore, ap.DMode); err != nil {
		return
	}
	// output domain data
	if err = os.MkdirAll(ap.paths.OutputDomainDataFilepaths, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainDataLogLevels, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainDataSettings, ap.DMode); err != nil {
		return
	}
	// output domain store
	if err = os.MkdirAll(ap.paths.OutputDomainStoreRecord, ap.DMode); err != nil {
		return
	}
	if err = os.Mkdir(ap.paths.OutputDomainStoreStorer, ap.DMode); err != nil {
		return
	}
	if err = os.Mkdir(ap.paths.OutputDomainStoreStoring, ap.DMode); err != nil {
		return
	}
	// output domain lpc folder and subfolders
	if err = os.MkdirAll(ap.paths.OutputDomainLPCMessage, ap.DMode); err != nil {
		return
	}

	// output mainprocess folder and sub folders.
	if err = os.MkdirAll(ap.paths.OutputMainProcessLPCDispatch, ap.DMode); err != nil {
		return
	}
	// output renderer folder and sub folders.
	if err = os.MkdirAll(ap.paths.OutputRendererCSS, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererMyCSS, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererPanels, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererPaneling, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputSitePack, ap.DMode); err != nil {
		return
	}

	// v 13
	if err = os.MkdirAll(ap.paths.OutputRendererProofs, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererDisplay, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererDOM, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererMarkup, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererEvent, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererWindow, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererApplication, ap.DMode); err != nil {
		return
	}

	// renderer framework
	if err = os.MkdirAll(ap.paths.OutputRendererCallBack, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererLocation, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererSpawnPack, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererLPC, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererViewTools, ap.DMode); err != nil {
		return
	}

	return nil
}

// CreateTemplateFolder creates the renderer templates folder.
func (ap *ApplicationPaths) CreateTemplateFolder() error {
	return os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode)
}

// WriteFile writes a file.
func (ap *ApplicationPaths) WriteFile(fpath string, data []byte) (err error) {
	var ofile *os.File
	ofile, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, ap.FMode)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("WriteFile: opening file %s:", fpath))
		return
	}
	defer ofile.Close()
	if _, err = ofile.Write(data); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("WriteFile: writing to file %s", fpath))
	}
	return
}

// Copy copies a sources.FileMap[src] path to the dest path.
func (ap *ApplicationPaths) Copy(dest, src string) (err error) {
	var fsrc *os.File
	fsrc, err = os.Open(src)
	if err != nil {
		return
	}
	defer fsrc.Close()
	var fdest *os.File
	fdest, err = os.Create(dest)
	if err != nil {
		return
	}
	defer fdest.Close()
	l := 1024 * 32
	bb := make([]byte, l, l)
	for {
		var rcount int
		rcount, err = fsrc.Read(bb)
		if err != nil && err != io.EOF {
			return
		}
		if rcount == 0 {
			break
		}
		if _, err = fdest.Write(bb); err != nil {
			return
		}
	}
	return
}
