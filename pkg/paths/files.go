package paths

// FileNames is the names of files.
type FileNames struct {
	AttributesDotGo  string
	BuildDotSH       string
	CallBackDotGo    string
	CallerDotGo      string
	CallServerDotGo  string
	ClassDotGo       string
	ClientDotGo      string
	ControlerDotGo   string
	ColorsDotCSS     string
	CreateGetDotGo   string
	DataDotGo        string
	DocumentDotGo    string
	ExampleDotTXT    string
	ExampleGoDotTXT  string
	EventsDotGo      string
	FavIconDotICO    string
	FlagDotYAML      string
	FormsDotGo       string
	HeadDotTMPL      string
	HelperDotGo      string
	HelpersDotGo     string
	HTTPDotYAML      string
	InnerDotGo       string
	KickwasmDotYAML  string // yamlFileName
	KickwasmDotYML   string // ymlFileName
	LockedDotGo      string
	LogDotGo         string
	MainDotCSS       string
	MainDotGo        string
	MainDotTMPL      string
	MapDotGo         string
	NoHelpDotGo      string
	NotJSDotGo       string
	PanelDotGo       string
	PanelsDotGo      string
	PanelGroupDotGo  string
	PanelMapDotGo    string
	ParentChildDotGo string
	PresenterDotGo   string
	RunDotGo         string
	ScrollDotGo      string
	ServeDotGo       string
	SizeDotGo        string
	StyleDotGo       string
	ViewToolsDotGo   string
	WasmExecJS       string
	WebSocketDotGo   string
}

// GetFileNames returns the file names.
func GetFileNames() *FileNames {
	return &FileNames{
		AttributesDotGo:  "attributes.go",
		BuildDotSH:       "build.sh",
		CallBackDotGo:    "callback.go",
		CallerDotGo:      "caller.go",
		CallServerDotGo:  "callserver.go",
		ClassDotGo:       "class.go",
		ClientDotGo:      "client.go",
		ColorsDotCSS:     "colors.css",
		ControlerDotGo:   "controler.go",
		CreateGetDotGo:   "createGet.go",
		DataDotGo:        "data.go",
		DocumentDotGo:    "document.go",
		ExampleDotTXT:    "example.txt",
		ExampleGoDotTXT:  "exampleGo.txt",
		EventsDotGo:      "events.go",
		FavIconDotICO:    "favicon.ico",
		FlagDotYAML:      "flags.yaml",
		FormsDotGo:       "forms.go",
		HeadDotTMPL:      "head.tmpl",
		HelperDotGo:      "helper.go",
		HelpersDotGo:     "helpers.go",
		HTTPDotYAML:      "http.yaml",
		InnerDotGo:       "inner.go",
		KickwasmDotYAML:  "kickwasm.yaml",
		KickwasmDotYML:   "kickwasm.yml",
		LockedDotGo:      "locked.go",
		LogDotGo:         "log.go",
		MainDotCSS:       "main.css",
		MainDotGo:        "main.go",
		MainDotTMPL:      "main.tmpl",
		MapDotGo:         "map.go",
		NoHelpDotGo:      "noHelp.go",
		NotJSDotGo:       "notJS.go",
		PanelDotGo:       "panel.go",
		PanelsDotGo:      "panels.go",
		PanelGroupDotGo:  "panelGroup.go",
		PanelMapDotGo:    "panelMap.go",
		ParentChildDotGo: "parentChild.go",
		PresenterDotGo:   "presenter.go",
		RunDotGo:         "run.go",
		ScrollDotGo:      "scroll.go",
		ServeDotGo:       "serve.go",
		SizeDotGo:        "size.go",
		StyleDotGo:       "style.go",
		ViewToolsDotGo:   "viewtools.go",
		WasmExecJS:       "wasm_exec.js",
		WebSocketDotGo:   "websocket.go",
	}
}
