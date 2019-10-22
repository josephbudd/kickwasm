package paths

// FileNames is the names of files.
type FileNames struct {
	AppDotWASM         string
	AttributesDotGo    string
	BuildDotSH         string
	BuildPackDotSH     string
	CallBackDotGo      string
	ClassDotGo         string
	ColorsDotCSS       string
	CreateGetDotGo     string
	DocumentDotGo      string
	EventDotGo         string
	ExampleDotTXT      string
	ExampleGoDotTXT    string
	EventsDotGo        string
	FavIconDotICO      string
	FilePathsDotGo     string
	FlagDotYAML        string
	FormsDotGo         string
	HeadDotTMPL        string
	HelpersDotGo       string
	HelpingDotGo       string
	HTTPDotYAML        string
	InnerDotGo         string
	InitDotGo          string
	InstructionsDotTXT string
	LogLevelsDotGo     string
	KickwasmDotYAML    string // yamlFileName
	KickwasmDotYML     string // ymlFileName
	MainDotCSS         string
	MainDotGo          string
	UCMainDotGo        string
	MainDotTMPL        string
	MapDotGo           string
	MarkupDotGo        string
	ModalDotGo         string
	UserContentDotCSS  string
	NotJSDotGo         string
	PanelsDotGo        string
	PanelMapDotGo      string
	ParentChildDotGo   string
	RecordsDotGo       string
	ScrollDotGo        string
	ServeDotGo         string
	SettingsDotGo      string
	SizeDotGo          string
	StoresDotGo        string
	StoresDotYAML      string
	StyleDotGo         string
	ViewToolsDotGo     string
	WasmExecJS         string

	// lpc & dispatch
	ChannelsDotGo  string
	ClientDotGo    string
	DispatchDotGo  string
	LockedDotGo    string
	LogDotGo       string
	PayloadDotGo   string
	RunDotGo       string
	ServerDotGo    string
	WebSocketDotGo string

	// markup panels
	MessengerDotGo  string
	ControllerDotGo string
	DataDotGo       string
	LCDataDotGo     string
	PanelGroupDotGo string
	PanelDotGo      string
	PresenterDotGo  string

	// spawn tabs
	APIDotGo     string
	PrepareDotGo string
	SpawnDotGo   string

	// widgets
	WidgetsDotGo string

	// vscode
	VSCodeMPWorkSpaceJSON string
	VSCodeRPWorkSpaceJSON string
}

// GetFileNames returns the file names.
func GetFileNames() *FileNames {
	return &FileNames{
		AppDotWASM:         "app.wasm",
		AttributesDotGo:    "attributes.go",
		BuildDotSH:         "build.sh",
		BuildPackDotSH:     "buildPack.sh",
		CallBackDotGo:      "callback.go",
		ClassDotGo:         "class.go",
		ColorsDotCSS:       "colors.css",
		CreateGetDotGo:     "createGet.go",
		DocumentDotGo:      "document.go",
		EventDotGo:         "event.go",
		ExampleDotTXT:      "example.txt",
		ExampleGoDotTXT:    "exampleGo.txt",
		EventsDotGo:        "events.go",
		FavIconDotICO:      "favicon.ico",
		FilePathsDotGo:     "filepaths.go",
		FlagDotYAML:        "flags.yaml",
		FormsDotGo:         "forms.go",
		HeadDotTMPL:        "Head.tmpl",
		HelpersDotGo:       "helpers.go",
		HelpingDotGo:       "Helping.go",
		HTTPDotYAML:        "Http.yaml",
		InnerDotGo:         "inner.go",
		InitDotGo:          "Init.go",
		InstructionsDotTXT: "instructions.txt",
		LogLevelsDotGo:     "LogLevels.go",
		KickwasmDotYAML:    "kickwasm.yaml",
		KickwasmDotYML:     "kickwasm.yml",
		MainDotCSS:         "main.css",
		MainDotGo:          "main.go",
		UCMainDotGo:        "Main.go",
		MainDotTMPL:        "main.tmpl",
		MapDotGo:           "map.go",
		MarkupDotGo:        "markup.go",
		ModalDotGo:         "modal.go",
		UserContentDotCSS:  "Usercontent.css",
		NotJSDotGo:         "notJS.go",
		PanelsDotGo:        "panels.go",
		PanelMapDotGo:      "panelMap.go",
		ParentChildDotGo:   "parentChild.go",
		RecordsDotGo:       "Records.go",
		ScrollDotGo:        "scroll.go",
		ServeDotGo:         "Serve.go",
		SettingsDotGo:      "settings.go",
		SizeDotGo:          "size.go",
		StoresDotGo:        "stores.go",
		StoresDotYAML:      "stores.yaml",
		StyleDotGo:         "style.go",
		ViewToolsDotGo:     "viewtools.go",
		WasmExecJS:         "wasm_exec.js",

		// lpc & dispatch
		ChannelsDotGo:  "channels.go",
		ClientDotGo:    "client.go",
		DispatchDotGo:  "dispatch.go",
		LockedDotGo:    "locked.go",
		LogDotGo:       "Log.go",
		PayloadDotGo:   "payload.go",
		RunDotGo:       "run.go",
		ServerDotGo:    "server.go",
		WebSocketDotGo: "websocket.go",

		// markup panels
		MessengerDotGo:  "Messenger.go",
		ControllerDotGo: "Controller.go",
		DataDotGo:       "Data.go",
		LCDataDotGo:     "data.go",
		PanelGroupDotGo: "group.go",
		PanelDotGo:      "Panel.go",
		PresenterDotGo:  "Presenter.go",

		// spawn tabs
		APIDotGo:     "api.go",
		PrepareDotGo: "prepare.go",
		SpawnDotGo:   "spawn.go",

		// widgets
		WidgetsDotGo: "widgets.go",

		// vscode
		VSCodeMPWorkSpaceJSON: "mainprocess.code-workspace",
		VSCodeRPWorkSpaceJSON: "rendererprocess.code-workspace",
	}
}
