package paths

// FolderNames is the names of folders.
type FolderNames struct {
	BoltStoring     string
	CSS             string
	MyCSS           string
	Data            string
	Domain          string
	DotKickwasm     string
	DotKickstore    string
	FilePaths       string
	Implementations string
	Interfaces      string
	LogLevels       string
	MainProcess     string
	Paneling        string
	Panels          string
	Record          string
	Renderer        string
	RendererSite    string
	Settings        string
	Store           string
	Storer          string
	Storing         string
	Templates       string
	ViewTools       string
	WASM            string
	YAML            string

	SpawnPanels    string
	SpawnTemplates string
	SpawnPack      string

	LPC      string
	Dispatch string
	Message  string

	SitePack string

	// v 14 renderer

	DOM         string
	CallBack    string
	Event       string
	Location    string
	Markup      string
	Window      string
	Display     string
	Framework   string
	Proofs      string
	Application string
	API         string
}

// GetFolderNames returns the folder names.
func GetFolderNames() *FolderNames {
	return &FolderNames{
		BoltStoring:     "boltstoring",
		CSS:             "css",
		MyCSS:           "mycss",
		Data:            "data",
		Domain:          "domain",
		DotKickwasm:     ".kickwasm",
		DotKickstore:    ".kickstore",
		FilePaths:       "filepaths",
		Implementations: "implementations",
		Interfaces:      "interfaces",
		LogLevels:       "loglevels",
		MainProcess:     "mainprocess",
		Paneling:        "paneling",
		Panels:          "panels",
		Record:          "record",
		Renderer:        "rendererprocess",
		RendererSite:    "site",
		Settings:        "settings",
		Store:           "store",
		Storer:          "storer",
		Storing:         "storing",
		Templates:       "templates",
		ViewTools:       "viewtools",
		WASM:            "wasm",
		YAML:            "yaml",

		SpawnPanels:    "spawnPanels",
		SpawnTemplates: "spawnTemplates",
		SpawnPack:      "spawnpack",

		LPC:      "lpc",
		Dispatch: "dispatch",
		Message:  "message",

		SitePack: "sitepack",

		// v 14 renderer

		DOM:         "dom",
		CallBack:    "callback",
		Event:       "event",
		Location:    "location",
		Markup:      "markup",
		Window:      "window",
		Display:     "display",
		Framework:   "framework",
		Proofs:      "proofs",
		Application: "application",
		API:         "api",
	}
}
