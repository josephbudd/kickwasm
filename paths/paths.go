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

	ImportDomainInterfacesCallers = "/domain/interfaces/caller"
	ImportDomainInterfacesStorers = "/domain/interfaces/storer"

	// domain data

	ImportDomainDataFilepaths  = "/domain/data/filepaths"
	ImportDomainDataCallParams = "/domain/data/callParams"

	// domain types

	ImportDomainTypes = "/domain/types"

	// domain implementations

	ImportDomainImplementationsCalling     = "/domain/implementations/calling"
	ImportDomainImplementationsStoringBolt = "/domain/implementations/storing/boltstoring"

	// main process

	ImportMainProcessCallServer    = "/mainprocess/callserver"
	ImportMainProcessServices      = "/mainprocess/services"
	ImportMainProcessServicesAbout = "/mainprocess/services/about"

	// renderer

	ImportRendererCall      = "/renderer/call"
	ImportRendererPanels    = "/renderer/panels"
	ImportRendererViewTools = "/renderer/viewtools"
)

// ApplicationPathsI is a test
type ApplicationPathsI interface {
	GetPaths() *Paths
	CreateAboutFolders() error
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
	ap.paths.ImportDomainInterfacesCallers = ImportDomainInterfacesCallers
	ap.paths.ImportDomainInterfacesStorers = ImportDomainInterfacesStorers
	ap.paths.ImportDomainDataFilepaths = ImportDomainDataFilepaths
	ap.paths.ImportDomainDataCallParams = ImportDomainDataCallParams
	ap.paths.ImportDomainTypes = ImportDomainTypes
	ap.paths.ImportDomainImplementationsCalling = ImportDomainImplementationsCalling
	ap.paths.ImportDomainImplementationsStoringBolt = ImportDomainImplementationsStoringBolt

	ap.paths.ImportMainProcessServices = ImportMainProcessServices
	ap.paths.ImportMainProcessServicesAbout = ImportMainProcessServicesAbout
	ap.paths.ImportMainProcessCallServer = ImportMainProcessCallServer

	ap.paths.ImportRendererCall = ImportRendererCall
	ap.paths.ImportRendererPanels = ImportRendererPanels
	ap.paths.ImportRendererViewTools = ImportRendererViewTools
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

// Paths are the folder paths
type Paths struct {
	Output            string
	OutputDotKick     string
	OutputDotKickYAML string

	// output domain

	OutputDomain string

	OutputDomainInterfaces        string
	OutputDomainInterfacesCallers string
	OutputDomainInterfacesStorers string

	OutputDomainData           string
	OutputDomainDataFilepaths  string
	OutputDomainDataCallParams string

	OutputDomainImplementations            string
	OutputDomainImplementationsCalling     string
	OutputDomainImplementationsStoring     string
	OutputDomainImplementationsStoringBolt string

	OutputDomainTypes string

	// output main process

	OutputMainProcess              string
	OutputMainProcessCallServer    string
	OutputMainProcessServices      string
	OutputMainProcessServicesAbout string

	// output renderer

	OutputRenderer          string
	OutputRendererCSS       string
	OutputRendererTemplates string
	OutputRendererCall      string
	OutputRendererPanels    string
	OutputRendererViewTools string

	// import domain

	ImportDomainInterfacesCallers          string
	ImportDomainInterfacesStorers          string
	ImportDomainDataFilepaths              string
	ImportDomainDataCallParams             string
	ImportDomainTypes                      string
	ImportDomainImplementationsCalling     string
	ImportDomainImplementationsStoringBolt string

	// import main process

	ImportMainProcessCallServer    string
	ImportMainProcessServices      string
	ImportMainProcessServicesAbout string

	// import renderer

	ImportRendererCall      string
	ImportRendererPanels    string
	ImportRendererViewTools string
}

// initializeOutput defines the output paths
func (ap *ApplicationPaths) initializeOutput(pwd, outputFolder, appname string) {
	// set and create the output folder and sub folders.
	// fix the output folder.
	ap.paths.Output = filepath.Join(pwd, outputFolder, appname)
	// output .kick folder and sub folders
	ap.paths.OutputDotKick = filepath.Join(ap.paths.Output, ".kick")
	ap.paths.OutputDotKickYAML = filepath.Join(ap.paths.OutputDotKick, "yaml")
	// output domain folder and sub folders.
	ap.paths.OutputDomain = filepath.Join(ap.paths.Output, "domain")
	ap.paths.OutputDomainInterfaces = filepath.Join(ap.paths.OutputDomain, "interfaces")
	ap.paths.OutputDomainInterfacesCallers = filepath.Join(ap.paths.OutputDomainInterfaces, "caller")
	ap.paths.OutputDomainInterfacesStorers = filepath.Join(ap.paths.OutputDomainInterfaces, "storer")
	ap.paths.OutputDomainData = filepath.Join(ap.paths.OutputDomain, "data")
	ap.paths.OutputDomainDataFilepaths = filepath.Join(ap.paths.OutputDomainData, "filepaths")
	ap.paths.OutputDomainDataCallParams = filepath.Join(ap.paths.OutputDomainData, "callParams")
	ap.paths.OutputDomainImplementations = filepath.Join(ap.paths.OutputDomain, "implementations")
	ap.paths.OutputDomainImplementationsCalling = filepath.Join(ap.paths.OutputDomainImplementations, "calling")
	ap.paths.OutputDomainImplementationsStoring = filepath.Join(ap.paths.OutputDomainImplementations, "storing")
	ap.paths.OutputDomainImplementationsStoringBolt = filepath.Join(ap.paths.OutputDomainImplementationsStoring, "boltstoring")
	ap.paths.OutputDomainTypes = filepath.Join(ap.paths.OutputDomain, "types")
	// output renderer folder and sub folders.
	ap.paths.OutputRenderer = filepath.Join(ap.paths.Output, "renderer")
	ap.paths.OutputRendererCSS = filepath.Join(ap.paths.OutputRenderer, "css")
	ap.paths.OutputRendererTemplates = filepath.Join(ap.paths.OutputRenderer, "templates")
	ap.paths.OutputRendererCall = filepath.Join(ap.paths.OutputRenderer, "call")
	ap.paths.OutputRendererPanels = filepath.Join(ap.paths.OutputRenderer, "panels")
	ap.paths.OutputRendererViewTools = filepath.Join(ap.paths.OutputRenderer, "viewtools")
	// output mainprocess folder and sub folders.
	ap.paths.OutputMainProcess = filepath.Join(ap.paths.Output, "mainprocess")
	ap.paths.OutputMainProcessCallServer = filepath.Join(ap.paths.OutputMainProcess, "callserver")
	ap.paths.OutputMainProcessServices = filepath.Join(ap.paths.OutputMainProcess, "services")
	ap.paths.OutputMainProcessServicesAbout = filepath.Join(ap.paths.OutputMainProcessServices, "about")
}

// MakeOutput creates the output paths
func (ap *ApplicationPaths) MakeOutput() error {
	// output .kick folder and sub folders
	if err := os.MkdirAll(ap.paths.OutputDotKick, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputDotKickYAML, ap.DMode); err != nil {
		return err
	}
	// output domain interfaces
	if err := os.MkdirAll(ap.paths.OutputDomainInterfaces, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputDomainInterfacesCallers, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputDomainInterfacesStorers, ap.DMode); err != nil {
		return err
	}
	// output domain data
	if err := os.MkdirAll(ap.paths.OutputDomainDataFilepaths, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputDomainDataCallParams, ap.DMode); err != nil {
		return err
	}
	// output domain implementations
	if err := os.MkdirAll(ap.paths.OutputDomainImplementationsCalling, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputDomainImplementationsStoringBolt, ap.DMode); err != nil {
		return err
	}
	// output domain types
	if err := os.MkdirAll(ap.paths.OutputDomainTypes, ap.DMode); err != nil {
		return err
	}
	// output mainprocess folder and sub folders.
	if err := os.MkdirAll(ap.paths.OutputMainProcessCallServer, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessServices, ap.DMode); err != nil {
		return err
	}
	// output renderer folder and sub folders.
	if err := os.MkdirAll(ap.paths.OutputRendererCSS, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererCall, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererPanels, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererViewTools, ap.DMode); err != nil {
		return err
	}
	return nil
}

// CreateAboutFolders creates the generated source code folder for the mainprocess/about package.
func (ap *ApplicationPaths) CreateAboutFolders() error {
	return os.MkdirAll(ap.paths.OutputMainProcessServicesAbout, ap.DMode)
}

// CreateTemplateFolder creates the generated source code folder for the mainprocess/about package.
func (ap *ApplicationPaths) CreateTemplateFolder() error {
	return os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode)
}

// WriteFile writes a file.
func (ap *ApplicationPaths) WriteFile(fpath string, data []byte) error {
	ofile, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, ap.FMode)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("WriteFile: opening file %s", fpath))
	}
	defer ofile.Close()
	if _, err = ofile.Write(data); err != nil {
		return errors.Wrap(err, fmt.Sprintf("WriteFile: writing to file %s", fpath))
	}
	return nil
}

// Copy copies a sources.FileMap[src] path to the dest path.
func (ap *ApplicationPaths) Copy(dest, src string) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()
	fdest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fdest.Close()
	l := 1024 * 32
	bb := make([]byte, l, l)
	for {
		rcount, err := fsrc.Read(bb)
		if err != nil && err != io.EOF {
			return err
		}
		if rcount == 0 {
			break
		}
		if _, err := fdest.Write(bb); err != nil {
			return err
		}
	}
	return nil
}
