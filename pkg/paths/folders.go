package paths

// FolderNames is the names of folders.
type FolderNames struct {
	BoltStoring     string
	CallClient      string
	CallIDs         string
	Caller          string
	Calling         string
	Calls           string
	CallServer      string
	CSS             string
	Data            string
	Domain          string
	DotKickwasm     string
	FilePaths       string
	Implementations string
	Interfaces      string
	LogLevels       string
	MainProcess     string
	NotJS           string
	PanelHelper     string
	PanelHelping    string
	Panels          string
	Renderer        string
	RendererSite    string
	Services        string
	Settings        string
	Storer          string
	Storing         string
	Templates       string
	Types           string
	ViewTools       string
	YAML            string
}

// GetFolderNames returns the folder names.
func GetFolderNames() *FolderNames {
	return &FolderNames{
		BoltStoring:     "boltstoring",
		CallClient:      "callClient",
		CallIDs:         "callids",
		Caller:          "caller",
		Calling:         "calling",
		Calls:           "calls",
		CallServer:      "callserver",
		CSS:             "css",
		Data:            "data",
		Domain:          "domain",
		DotKickwasm:     ".kickwasm",
		FilePaths:       "filepaths",
		Implementations: "implementations",
		Interfaces:      "interfaces",
		LogLevels:       "loglevels",
		MainProcess:     "mainprocess",
		NotJS:           "notjs",
		PanelHelper:     "panelHelper",
		PanelHelping:    "panelHelping",
		Panels:          "panels",
		Renderer:        "renderer",
		RendererSite:    "site",
		Services:        "services",
		Settings:        "settings",
		Storer:          "storer",
		Storing:         "storing",
		Templates:       "templates",
		Types:           "types",
		ViewTools:       "viewtools",
		YAML:            "yaml",
	}
}
