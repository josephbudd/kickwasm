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
	ImportData                        = "/mainprocess/data"
	ImportMainProcessRepositoriesBolt = "/mainprocess/repositories/bolt"
	ImportMainProcessServices         = "/mainprocess/services"
	ImportMainProcessDataRecords      = "/mainprocess/data/records"
	ImportMainProcessBehaviorRepoi    = "/mainprocess/behavior/repos"
	ImportTransports                  = "/mainprocess/transports"
)

// ApplicationPathsI is a test
type ApplicationPathsI interface {
	GetPaths() *Paths
	CreateAboutFolders() error
	CreateTemplateFolder() error
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
	Output                      string
	OutputDotKick               string
	OutputDotKickYAML           string
	OutputRenderer              string
	OutputRendererWASMCaller    string
	OutputRendererTemplates     string
	OutputRendererWASMViewTools string
	OutputRendererWASM          string
	OutputRendererWASMPanels    string
	OutputRendererCSS           string

	OutputMainProcess                 string
	OutputMainProcessData             string
	OutputMainProcessRepositories     string
	OutputMainProcessRepositoriesBolt string
	OutputMainProcessServices         string
	OutputMainProcessServicesAbout    string
	OutputMainProcessTransports       string
	OutputMainProcessTransportsCalls  string

	OutputMainProcessTransportsCallServer string
	OutputMainProcessDataFilePaths        string
	OutputMainProcessDataRecords          string
	OutputMainProcessBehavior             string
	OutputMainProcessBehaviorRepoI        string

	ImportMainProcessBehaviorRepoi        string
	ImportMainProcessDataFilePaths        string
	ImportMainProcessDataRecords          string
	ImportMainProcessRepositoriesBolt     string
	ImportMainProcessServices             string
	ImportMainProcessServicesAbout        string
	ImportMainProcessTransportsCalls      string
	ImportMainProcessTransportsCallServer string

	ImportRendererWASMCall      string
	ImportRendererWASMPanels    string
	ImportRendererWASMViewTools string
}

// initializeOutput defines the output paths
func (ap *ApplicationPaths) initializeOutput(pwd, outputFolder, appname string) {
	// set and create the output folder and sub folders.
	// fix the output folder.
	ap.paths.Output = filepath.Join(pwd, outputFolder, appname)
	// output .kick folder and sub folders
	ap.paths.OutputDotKick = filepath.Join(ap.paths.Output, ".kick")
	ap.paths.OutputDotKickYAML = filepath.Join(ap.paths.OutputDotKick, "yaml")
	// output renderer folder and sub folders.
	ap.paths.OutputRenderer = filepath.Join(ap.paths.Output, "renderer")
	ap.paths.OutputRendererCSS = filepath.Join(ap.paths.OutputRenderer, "css")
	ap.paths.OutputRendererTemplates = filepath.Join(ap.paths.OutputRenderer, "templates")
	ap.paths.OutputRendererWASM = filepath.Join(ap.paths.OutputRenderer, "wasm")
	ap.paths.OutputRendererWASMCaller = filepath.Join(ap.paths.OutputRendererWASM, "transports/call")
	ap.paths.OutputRendererWASMPanels = filepath.Join(ap.paths.OutputRendererWASM, "panels")
	ap.paths.OutputRendererWASMViewTools = filepath.Join(ap.paths.OutputRendererWASM, "viewtools")
	// output mainprocess folder and sub folders.
	ap.paths.OutputMainProcess = filepath.Join(ap.paths.Output, "mainprocess")
	ap.paths.OutputMainProcessBehavior = filepath.Join(ap.paths.OutputMainProcess, "behavior")
	ap.paths.OutputMainProcessBehaviorRepoI = filepath.Join(ap.paths.OutputMainProcessBehavior, "repoi")
	ap.paths.OutputMainProcessData = filepath.Join(ap.paths.OutputMainProcess, "data")
	ap.paths.OutputMainProcessDataFilePaths = filepath.Join(ap.paths.OutputMainProcessData, "filepaths")
	ap.paths.OutputMainProcessDataRecords = filepath.Join(ap.paths.OutputMainProcessData, "records")
	ap.paths.OutputMainProcessRepositories = filepath.Join(ap.paths.OutputMainProcess, "repositories")
	ap.paths.OutputMainProcessRepositoriesBolt = filepath.Join(ap.paths.OutputMainProcessRepositories, "bolt")
	ap.paths.OutputMainProcessServices = filepath.Join(ap.paths.OutputMainProcess, "services")
	ap.paths.OutputMainProcessServicesAbout = filepath.Join(ap.paths.OutputMainProcessServices, "about")
	ap.paths.OutputMainProcessTransports = filepath.Join(ap.paths.OutputMainProcess, "transports")
	ap.paths.OutputMainProcessTransportsCalls = filepath.Join(ap.paths.OutputMainProcessTransports, "calls")
	ap.paths.OutputMainProcessTransportsCallServer = filepath.Join(ap.paths.OutputMainProcessTransports, "callserver")

	// import paths
	ap.paths.ImportMainProcessBehaviorRepoi = "/mainprocess/behavior/repoi"
	ap.paths.ImportMainProcessDataFilePaths = "/mainprocess/data/filepaths"
	ap.paths.ImportMainProcessDataRecords = "/mainprocess/data/records"
	ap.paths.ImportMainProcessRepositoriesBolt = "/mainprocess/repositories/bolt"
	ap.paths.ImportMainProcessServices = "/mainprocess/services"
	ap.paths.ImportMainProcessServicesAbout = "/mainprocess/services/about"
	ap.paths.ImportMainProcessTransportsCalls = "/mainprocess/transports/calls"
	ap.paths.ImportMainProcessTransportsCallServer = "/mainprocess/transports/callserver"

	ap.paths.ImportRendererWASMCall = "/renderer/wasm/transports/call"
	ap.paths.ImportRendererWASMViewTools = "/renderer/wasm/viewtools"
	ap.paths.ImportRendererWASMPanels = "/renderer/wasm/panels"

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
	// output renderer folder and sub folders.
	if err := os.MkdirAll(ap.paths.OutputRenderer, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererWASMCaller, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererTemplates, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererWASMViewTools, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererWASMPanels, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputRendererCSS, ap.DMode); err != nil {
		return err
	}
	// output mainprocess folder and sub folders.
	if err := os.MkdirAll(ap.paths.OutputMainProcessTransportsCalls, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessTransportsCallServer, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessDataFilePaths, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessDataRecords, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessRepositoriesBolt, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessBehaviorRepoI, ap.DMode); err != nil {
		return err
	}
	if err := os.MkdirAll(ap.paths.OutputMainProcessServices, ap.DMode); err != nil {
		return err
	}
	return nil
}

// CreateAboutFolders creates the generated source code folder for the mainprocess/about package.
func (ap *ApplicationPaths) CreateAboutFolders() error {
	return os.MkdirAll(ap.paths.OutputMainProcessServicesAbout, ap.DMode)
}

// CreateTemplateFolder creates a template folder.
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
