package paths

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Import paths
const (
	// domain interfaces

	importDomainInterfacesCallers = "/domain/interfaces/caller"
	importDomainInterfacesStorers = "/domain/interfaces/storer"

	// domain data

	importDomainDataFilepaths  = "/domain/data/filepaths"
	importDomainDataCallIDs    = "/domain/data/callids"
	importDomainDataCallParams = "/domain/data/callParams"
	importDomainDataLogLevels  = "/domain/data/loglevels"
	importDomainDataSettings   = "/domain/data/settings"

	// domain types

	importDomainTypes = "/domain/types"

	// domain implementations

	importDomainImplementationsCalling     = "/domain/implementations/calling"
	importDomainImplementationsStoringBolt = "/domain/implementations/storing/boltstoring"

	// main process

	importMainProcessCalls      = "/mainprocess/calls"
	importMainProcessCallServer = "/mainprocess/callserver"
	importMainProcessServices   = "/mainprocess/services"

	// renderer

	importRendererCallClient                 = "/renderer/callClient"
	importRendererCalls                      = "/renderer/calls"
	importRendererPanels                     = "/renderer/panels"
	importRendererInterfacesPanelHelper      = "/renderer/interfaces/panelHelper"
	importRendererImplementationsPanelHelper = "/renderer/implementations/panelHelping"

	// renderer site

	importRendererViewTools = "/renderer/viewtools"
	importRendererNotJS     = "/renderer/notjs"
)

// Imports is the import paths.
type Imports struct {
	ImportDomainInterfacesCallers string
	ImportDomainInterfacesStorers string

	// domain data

	ImportDomainDataFilepaths  string
	ImportDomainDataCallIDs    string
	ImportDomainDataCallParams string
	ImportDomainDataLogLevels  string
	ImportDomainDataSettings   string

	// domain types

	ImportDomainTypes string

	// domain implementations

	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	// main process

	ImportMainProcessCalls      string
	ImportMainProcessCallServer string
	ImportMainProcessServices   string

	// renderer

	ImportRendererCallClient                 string
	ImportRendererCalls                      string
	ImportRendererPanels                     string
	ImportRendererViewTools                  string
	ImportRendererNotJS                      string
	ImportRendererInterfacesPanelHelper      string
	ImportRendererImplementationsPanelHelper string
}

// GetImports returns the go import paths.
func GetImports() *Imports {
	return &Imports{
		ImportDomainInterfacesCallers: importDomainInterfacesCallers,
		ImportDomainInterfacesStorers: importDomainInterfacesStorers,

		// domain data

		ImportDomainDataFilepaths:  importDomainDataFilepaths,
		ImportDomainDataCallIDs:    importDomainDataCallIDs,
		ImportDomainDataCallParams: importDomainDataCallParams,
		ImportDomainDataLogLevels:  importDomainDataLogLevels,
		ImportDomainDataSettings:   importDomainDataSettings,

		// domain types

		ImportDomainTypes: importDomainTypes,

		// domain implementations

		ImportDomainImplementationsCalling:     importDomainImplementationsCalling,
		ImportDomainImplementationsStoringBolt: importDomainImplementationsStoringBolt,

		// main process

		ImportMainProcessCalls:      importMainProcessCalls,
		ImportMainProcessCallServer: importMainProcessCallServer,
		ImportMainProcessServices:   importMainProcessServices,

		// renderer

		ImportRendererCallClient:                 importRendererCallClient,
		ImportRendererCalls:                      importRendererCalls,
		ImportRendererPanels:                     importRendererPanels,
		ImportRendererViewTools:                  importRendererViewTools,
		ImportRendererNotJS:                      importRendererNotJS,
		ImportRendererInterfacesPanelHelper:      importRendererInterfacesPanelHelper,
		ImportRendererImplementationsPanelHelper: importRendererImplementationsPanelHelper,
	}
}

// ApplicationPathsI is a test
type ApplicationPathsI interface {
	GetPaths() *Paths
	Copy(dest, src string) error
	WriteFile(fpath string, data []byte) error
	GetDMode() os.FileMode
	GetFMode() os.FileMode
}

// ApplicationPaths is the paths needed for kick.
type ApplicationPaths struct {
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

	ap.initializeOutput(pwd, outputFolder, appname)

	// import paths
	ap.paths.ImportDomainInterfacesCallers = importDomainInterfacesCallers
	ap.paths.ImportDomainInterfacesStorers = importDomainInterfacesStorers
	ap.paths.ImportDomainDataFilepaths = importDomainDataFilepaths
	ap.paths.ImportDomainDataCallIDs = importDomainDataCallIDs
	ap.paths.ImportDomainDataCallParams = importDomainDataCallParams
	ap.paths.ImportDomainDataLogLevels = importDomainDataLogLevels
	ap.paths.ImportDomainDataSettings = importDomainDataSettings
	ap.paths.ImportDomainTypes = importDomainTypes
	ap.paths.ImportDomainImplementationsCalling = importDomainImplementationsCalling
	ap.paths.ImportDomainImplementationsStoringBolt = importDomainImplementationsStoringBolt

	ap.paths.ImportMainProcessServices = importMainProcessServices
	ap.paths.ImportMainProcessCalls = importMainProcessCalls
	ap.paths.ImportMainProcessCallServer = importMainProcessCallServer

	ap.paths.ImportRendererCallClient = importRendererCallClient
	ap.paths.ImportRendererCalls = importRendererCalls
	ap.paths.ImportRendererPanels = importRendererPanels
	ap.paths.ImportRendererViewTools = importRendererViewTools
	ap.paths.ImportRendererNotJS = importRendererNotJS
	ap.paths.ImportRendererInterfacesPanelHelper = importRendererInterfacesPanelHelper
	ap.paths.ImportRendererImplementationsPanelHelper = importRendererImplementationsPanelHelper
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
func (ap *ApplicationPaths) GetFolderNames() *FolderNames {
	return GetFolderNames()
}

// Paths are the folder paths
type Paths struct {
	Output                 string
	OutputDotKickwasm      string
	OutputDotKickwasmYAML  string
	OutputDotKickwasmFlags string

	// output domain

	OutputDomain string

	OutputDomainInterfaces        string
	OutputDomainInterfacesCallers string
	OutputDomainInterfacesStorers string

	OutputDomainData          string
	OutputDomainDataFilepaths string
	OutputDomainDataCallIDs   string
	OutputDomainDataLogLevels string
	OutputDomainDataSettings  string

	OutputDomainImplementations            string
	OutputDomainImplementationsCalling     string
	OutputDomainImplementationsStoring     string
	OutputDomainImplementationsStoringBolt string

	OutputDomainTypes string

	// output main process

	OutputMainProcess           string
	OutputMainProcessCalls      string
	OutputMainProcessCallServer string
	OutputMainProcessServices   string

	// output renderer

	OutputRenderer                           string
	OutputRendererCSS                        string
	OutputRendererTemplates                  string
	OutputRendererCallClient                 string
	OutputRendererCalls                      string
	OutputRendererPanels                     string
	OutputRendererSite                       string
	OutputRendererViewTools                  string
	OutputRendererNotJS                      string
	OutputRendererInterfaces                 string
	OutputRendererInterfacesPanelHelper      string
	OutputRendererImplementations            string
	OutputRendererImplementationsPanelHelper string

	// import domain

	ImportDomainInterfacesCallers          string
	ImportDomainInterfacesStorers          string
	ImportDomainDataFilepaths              string
	ImportDomainDataCallIDs                string
	ImportDomainDataCallParams             string
	ImportDomainDataLogLevels              string
	ImportDomainDataSettings               string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	// import main process

	ImportMainProcessCalls      string
	ImportMainProcessCallServer string
	ImportMainProcessServices   string

	// import renderer

	ImportRendererCallClient                 string
	ImportRendererCalls                      string
	ImportRendererPanels                     string
	ImportRendererViewTools                  string
	ImportRendererNotJS                      string
	ImportRendererInterfacesPanelHelper      string
	ImportRendererImplementationsPanelHelper string
}

// initializeOutput defines the output paths
func (ap *ApplicationPaths) initializeOutput(pwd, outputFolder, appname string) {
	fileNames := GetFileNames()
	folderNames := GetFolderNames()
	// set and create the output folder and sub folders.
	// fix the output folder.
	ap.paths.Output = filepath.Join(pwd, outputFolder, appname)
	// output .kickwasm folder and sub folders
	ap.paths.OutputDotKickwasm = filepath.Join(ap.paths.Output, folderNames.DotKickwasm)
	ap.paths.OutputDotKickwasmYAML = filepath.Join(ap.paths.OutputDotKickwasm, folderNames.YAML)
	ap.paths.OutputDotKickwasmFlags = filepath.Join(ap.paths.OutputDotKickwasm, fileNames.FlagDotYAML)
	// output domain folder and sub folders.
	ap.paths.OutputDomain = filepath.Join(ap.paths.Output, folderNames.Domain)
	ap.paths.OutputDomainInterfaces = filepath.Join(ap.paths.OutputDomain, folderNames.Interfaces)
	ap.paths.OutputDomainInterfacesCallers = filepath.Join(ap.paths.OutputDomainInterfaces, folderNames.Caller)
	ap.paths.OutputDomainInterfacesStorers = filepath.Join(ap.paths.OutputDomainInterfaces, folderNames.Storer)
	ap.paths.OutputDomainData = filepath.Join(ap.paths.OutputDomain, folderNames.Data)
	ap.paths.OutputDomainDataFilepaths = filepath.Join(ap.paths.OutputDomainData, folderNames.FilePaths)
	ap.paths.OutputDomainDataCallIDs = filepath.Join(ap.paths.OutputDomainData, folderNames.CallIDs)
	ap.paths.OutputDomainDataLogLevels = filepath.Join(ap.paths.OutputDomainData, folderNames.LogLevels)
	ap.paths.OutputDomainDataSettings = filepath.Join(ap.paths.OutputDomainData, folderNames.Settings)
	ap.paths.OutputDomainImplementations = filepath.Join(ap.paths.OutputDomain, folderNames.Implementations)
	ap.paths.OutputDomainImplementationsCalling = filepath.Join(ap.paths.OutputDomainImplementations, folderNames.Calling)
	ap.paths.OutputDomainImplementationsStoring = filepath.Join(ap.paths.OutputDomainImplementations, folderNames.Storing)
	ap.paths.OutputDomainImplementationsStoringBolt = filepath.Join(ap.paths.OutputDomainImplementationsStoring, folderNames.BoltStoring)
	ap.paths.OutputDomainTypes = filepath.Join(ap.paths.OutputDomain, folderNames.Types)
	// output renderer folder and sub folders.
	ap.paths.OutputRenderer = filepath.Join(ap.paths.Output, folderNames.Renderer)
	ap.paths.OutputRendererCallClient = filepath.Join(ap.paths.OutputRenderer, folderNames.CallClient)
	ap.paths.OutputRendererCalls = filepath.Join(ap.paths.OutputRenderer, folderNames.Calls)
	ap.paths.OutputRendererPanels = filepath.Join(ap.paths.OutputRenderer, folderNames.Panels)
	ap.paths.OutputRendererInterfaces = filepath.Join(ap.paths.OutputRenderer, folderNames.Interfaces)
	ap.paths.OutputRendererInterfacesPanelHelper = filepath.Join(ap.paths.OutputRendererInterfaces, folderNames.PanelHelper)
	ap.paths.OutputRendererImplementations = filepath.Join(ap.paths.OutputRenderer, folderNames.Implementations)
	ap.paths.OutputRendererImplementationsPanelHelper = filepath.Join(ap.paths.OutputRendererImplementations, folderNames.PanelHelping)
	ap.paths.OutputRendererViewTools = filepath.Join(ap.paths.OutputRenderer, folderNames.ViewTools)
	ap.paths.OutputRendererNotJS = filepath.Join(ap.paths.OutputRenderer, folderNames.NotJS)

	// renderer site folders
	ap.paths.OutputRendererSite = filepath.Join(ap.paths.Output, folderNames.RendererSite)
	ap.paths.OutputRendererCSS = filepath.Join(ap.paths.OutputRendererSite, folderNames.CSS)
	ap.paths.OutputRendererTemplates = filepath.Join(ap.paths.OutputRendererSite, folderNames.Templates)

	// output mainprocess folder and sub folders.
	ap.paths.OutputMainProcess = filepath.Join(ap.paths.Output, folderNames.MainProcess)
	ap.paths.OutputMainProcessCalls = filepath.Join(ap.paths.OutputMainProcess, folderNames.Calls)
	ap.paths.OutputMainProcessCallServer = filepath.Join(ap.paths.OutputMainProcess, folderNames.CallServer)
	ap.paths.OutputMainProcessServices = filepath.Join(ap.paths.OutputMainProcess, folderNames.Services)
}

// MakeOutput creates the output paths
func (ap *ApplicationPaths) MakeOutput() (err error) {
	// output .kickwasm folder and sub folders
	if err = os.MkdirAll(ap.paths.OutputDotKickwasm, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDotKickwasmYAML, ap.DMode); err != nil {
		return
	}
	// output domain interfaces
	if err = os.MkdirAll(ap.paths.OutputDomainInterfaces, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainInterfacesCallers, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainInterfacesStorers, ap.DMode); err != nil {
		return
	}
	// output domain data
	if err = os.MkdirAll(ap.paths.OutputDomainDataFilepaths, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainDataCallIDs, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainDataLogLevels, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainDataSettings, ap.DMode); err != nil {
		return
	}
	// output domain implementations
	if err = os.MkdirAll(ap.paths.OutputDomainImplementationsCalling, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputDomainImplementationsStoringBolt, ap.DMode); err != nil {
		return
	}
	// output domain types
	if err = os.MkdirAll(ap.paths.OutputDomainTypes, ap.DMode); err != nil {
		return
	}
	// output mainprocess folder and sub folders.
	if err = os.MkdirAll(ap.paths.OutputMainProcessCalls, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputMainProcessCallServer, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputMainProcessServices, ap.DMode); err != nil {
		return
	}
	// output renderer folder and sub folders.
	if err = os.MkdirAll(ap.paths.OutputRendererCSS, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererCallClient, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererCalls, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererPanels, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererViewTools, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererNotJS, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererInterfaces, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererInterfacesPanelHelper, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererImplementations, ap.DMode); err != nil {
		return
	}
	if err = os.MkdirAll(ap.paths.OutputRendererImplementationsPanelHelper, ap.DMode); err != nil {
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
		err = errors.Wrap(err, fmt.Sprintf("WriteFile: opening file %s", fpath))
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
