package paths

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

	// v 14

	importRendererAPIDOM         = "/rendererprocess/api/dom"
	importRendererAPIMarkup      = "/rendererprocess/api/markup"
	importRendererAPIWindow      = "/rendererprocess/api/window"
	importRendererAPIEvent       = "/rendererprocess/api/event"
	importRendererAPIDisplay     = "/rendererprocess/api/display"
	importRendererAPIApplication = "/rendererprocess/api/application"

	// renderer framework

	importRendererFramework          = "/rendererprocess/framework"
	importRendererFrameworkCallBack  = "/rendererprocess/framework/callback"
	importRendererFrameworkLocation  = "/rendererprocess/framework/location"
	importRendererFrameworkLPC       = "/rendererprocess/framework/lpc"
	importRendererFrameworkSpawnPack = "/rendererprocess/framework/spawnpack"
	importRendererFrameworkViewTools = "/rendererprocess/framework/viewtools"
	importRendererFrameworkProofs    = "/rendererprocess/framework/proofs"

	// v 16

	importRendererAPIJSValue = "/rendererprocess/api/jsvalue"
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
	ImportRendererFrameworkLPC   string
	ImportMainProcessLPC         string
	ImportMainProcessLPCDispatch string

	// main process

	ImportMainProcessHomes string

	// renderer

	ImportRenderer                   string
	ImportRendererPanels             string
	ImportRendererFrameworkViewTools string
	ImportRendererPaneling           string

	ImportRendererFrameworkSpawnPack string
	ImportRendererSpawnPanels        string

	// store

	ImportDomainStore        string
	ImportDomainStoreRecord  string
	ImportDomainStoreStorer  string
	ImportDomainStoreStoring string

	// v 14

	ImportRendererFrameworkCallBack string
	ImportRendererFrameworkLocation string

	ImportRendererAPIDOM          string
	ImportRendererAPIMarkup       string
	ImportRendererAPIWindow       string
	ImportRendererAPIEvent        string
	ImportRendererAPIDisplay      string
	ImportRendererFrameworkProofs string
	ImportRendererAPIApplication  string

	// renderer framework

	ImportRendererFramework string

	// v 16

	ImportRendererAPIJSValue string
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
		ImportRendererFrameworkLPC:   importRendererFrameworkLPC,
		ImportMainProcessLPC:         importMainProcessLPC,
		ImportMainProcessLPCDispatch: importMainProcessLPCDispatch,

		// main process

		ImportMainProcessHomes: importMainProcessHomes,

		// renderer

		ImportRenderer:                   importRenderer,
		ImportRendererPanels:             importRendererPanels,
		ImportRendererFrameworkViewTools: importRendererFrameworkViewTools,
		ImportRendererPaneling:           importRendererPaneling,

		ImportRendererFrameworkSpawnPack: importRendererFrameworkSpawnPack,
		ImportRendererSpawnPanels:        importRendererSpawnPanels,

		// store

		ImportDomainStore:        importDomainStore,
		ImportDomainStoreRecord:  importDomainStoreRecord,
		ImportDomainStoreStorer:  importDomainStoreStorer,
		ImportDomainStoreStoring: importDomainStoreStoring,

		// v 14

		ImportRendererFrameworkCallBack: importRendererFrameworkCallBack,
		ImportRendererFrameworkLocation: importRendererFrameworkLocation,

		ImportRendererAPIDOM:          importRendererAPIDOM,
		ImportRendererAPIMarkup:       importRendererAPIMarkup,
		ImportRendererAPIWindow:       importRendererAPIWindow,
		ImportRendererAPIEvent:        importRendererAPIEvent,
		ImportRendererAPIDisplay:      importRendererAPIDisplay,
		ImportRendererFrameworkProofs: importRendererFrameworkProofs,
		ImportRendererAPIApplication:  importRendererAPIApplication,

		// renderer framework

		ImportRendererFramework: importRendererFramework,

		// v 16

		ImportRendererAPIJSValue: importRendererAPIJSValue,
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
	ap.paths.ImportRendererFrameworkLPC = importRendererFrameworkLPC
	ap.paths.ImportMainProcessLPC = importMainProcessLPC
	ap.paths.ImportMainProcessLPCDispatch = importMainProcessLPCDispatch

	ap.paths.ImportMainProcessHomes = importMainProcessHomes
	ap.paths.ImportMainProcessLPCDispatch = importMainProcessLPCDispatch

	ap.paths.ImportRenderer = importRenderer
	ap.paths.ImportRendererPanels = importRendererPanels
	ap.paths.ImportRendererFrameworkViewTools = importRendererFrameworkViewTools
	ap.paths.ImportRendererPaneling = importRendererPaneling

	ap.paths.ImportRendererFrameworkSpawnPack = importRendererFrameworkSpawnPack
	ap.paths.ImportRendererSpawnPanels = importRendererSpawnPanels

	ap.paths.ImportDomainStore = importDomainStore
	ap.paths.ImportDomainStoreRecord = importDomainStoreRecord
	ap.paths.ImportDomainStoreStorer = importDomainStoreStorer
	ap.paths.ImportDomainStoreStoring = importDomainStoreStoring

	// v 14

	ap.paths.ImportRendererAPIDOM = importRendererAPIDOM
	ap.paths.ImportRendererFrameworkCallBack = importRendererFrameworkCallBack
	ap.paths.ImportRendererFrameworkLocation = importRendererFrameworkLocation
	ap.paths.ImportRendererAPIMarkup = importRendererAPIMarkup
	ap.paths.ImportRendererAPIWindow = importRendererAPIWindow
	ap.paths.ImportRendererAPIEvent = importRendererAPIEvent
	ap.paths.ImportRendererAPIDisplay = importRendererAPIDisplay
	ap.paths.ImportRendererFrameworkProofs = importRendererFrameworkProofs
	ap.paths.ImportRendererAPIApplication = importRendererAPIApplication

	// framework

	ap.paths.ImportRendererFramework = importRendererFramework

	// v 16

	ap.paths.ImportRendererAPIJSValue = importRendererAPIJSValue
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

	OutputRenderer                   string
	OutputRendererCSS                string
	OutputRendererMyCSS              string
	OutputRendererTemplates          string
	OutputRendererSpawnTemplates     string
	OutputRendererPanels             string
	OutputRendererSpawns             string
	OutputRendererSite               string
	OutputRendererFrameworkViewTools string
	OutputRendererPaneling           string
	OutputRendererFrameworkSpawnPack string

	// output renderer lpc

	OutputRendererFrameworkLPC string

	// v 14

	OutputRendererAPI               string
	OutputRendererDOM               string
	OutputRendererFrameworkCallBack string
	OutputRendererFrameworkLocation string
	OutputRendererMarkup            string
	OutputRendererWindow            string
	OutputRendererEvent             string
	OutputRendererDisplay           string
	OutputRendererFrameworkProofs   string
	OutputRendererApplication       string

	// framework

	OutputRendererFramework string

	// v 16

	OutputRendererAPIJSValue string

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

	ImportRenderer                   string
	ImportRendererPanels             string
	ImportRendererFrameworkViewTools string
	ImportRendererPaneling           string

	ImportRendererFrameworkSpawnPack string
	ImportRendererSpawnPanels        string

	// import renderer lpc

	ImportRendererFrameworkLPC string

	// import store

	ImportDomainStore        string
	ImportDomainStoreRecord  string
	ImportDomainStoreStorer  string
	ImportDomainStoreStoring string

	// v 14

	ImportRendererAPIDOM            string
	ImportRendererFrameworkCallBack string
	ImportRendererFrameworkLocation string
	ImportRendererAPIMarkup         string
	ImportRendererAPIWindow         string
	ImportRendererAPIEvent          string
	ImportRendererAPIDisplay        string
	ImportRendererFrameworkProofs   string
	ImportRendererAPIApplication    string

	// renderer framework

	ImportRendererFramework string

	// v 16

	ImportRendererAPIJSValue string
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

	// v 14

	// output renderer api

	ap.paths.OutputRendererAPI = filepath.Join(ap.paths.OutputRenderer, folderNames.API)
	ap.paths.OutputRendererDisplay = filepath.Join(ap.paths.OutputRendererAPI, folderNames.Display)
	ap.paths.OutputRendererDOM = filepath.Join(ap.paths.OutputRendererAPI, folderNames.DOM)
	ap.paths.OutputRendererMarkup = filepath.Join(ap.paths.OutputRendererAPI, folderNames.Markup)
	ap.paths.OutputRendererEvent = filepath.Join(ap.paths.OutputRendererAPI, folderNames.Event)
	ap.paths.OutputRendererWindow = filepath.Join(ap.paths.OutputRendererAPI, folderNames.Window)
	ap.paths.OutputRendererApplication = filepath.Join(ap.paths.OutputRendererAPI, folderNames.Application)

	// output renderer framework

	ap.paths.OutputRendererFramework = filepath.Join(ap.paths.OutputRenderer, folderNames.Framework)
	ap.paths.OutputRendererFrameworkViewTools = filepath.Join(ap.paths.OutputRendererFramework, folderNames.ViewTools)
	ap.paths.OutputRendererFrameworkSpawnPack = filepath.Join(ap.paths.OutputRendererFramework, folderNames.SpawnPack)
	ap.paths.OutputRendererFrameworkCallBack = filepath.Join(ap.paths.OutputRendererFramework, folderNames.CallBack)
	ap.paths.OutputRendererFrameworkLocation = filepath.Join(ap.paths.OutputRendererFramework, folderNames.Location)
	ap.paths.OutputRendererFrameworkProofs = filepath.Join(ap.paths.OutputRendererFramework, folderNames.Proofs)
	// ap.paths.OutputRendererFrameworkLPC is below

	// output mainprocess folder and sub folders.
	ap.paths.OutputMainProcess = filepath.Join(ap.paths.Output, folderNames.MainProcess)

	// output lpc
	ap.paths.OutputDomainLPC = filepath.Join(ap.paths.OutputDomain, folderNames.LPC)
	ap.paths.OutputDomainLPCMessage = filepath.Join(ap.paths.OutputDomainLPC, folderNames.Message)
	ap.paths.OutputRendererFrameworkLPC = filepath.Join(ap.paths.OutputRendererFramework, folderNames.LPC)
	ap.paths.OutputMainProcessLPC = filepath.Join(ap.paths.OutputMainProcess, folderNames.LPC)
	ap.paths.OutputMainProcessLPCDispatch = filepath.Join(ap.paths.OutputMainProcessLPC, folderNames.Dispatch)

	// version 16

	ap.paths.OutputRendererAPIJSValue = filepath.Join(ap.paths.OutputRendererAPI, folderNames.JSValue)
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

	// v 14
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkProofs, ap.DMode); err != nil {
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
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkCallBack, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkLocation, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkSpawnPack, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkLPC, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererFrameworkViewTools, ap.DMode); err != nil {
		return
	}

	// v 16

	if err = os.MkdirAll(ap.paths.OutputRendererAPIJSValue, ap.DMode); err != nil {
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
		err = fmt.Errorf("WriteFile: opening file %s: %w", fpath, err)
		return
	}
	defer ofile.Close()
	if _, err = ofile.Write(data); err != nil {
		err = fmt.Errorf("WriteFile: writing to file %s: %w", fpath, err)
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
