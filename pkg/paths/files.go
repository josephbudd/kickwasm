package paths

// FileNames is the names of files.
type FileNames struct {
	AppDotWASM         string
	BuildDotSH         string
	BuildPackDotSH     string
	ClassDotGo         string
	ColorsDotCSS       string
	CreateGetDotGo     string
	DocumentDotGo      string
	ExampleDotTXT      string
	ExampleGoDotTXT    string
	EventsDotGo        string
	FavIconDotICO      string
	FilePathsDotGo     string
	FlagDotYAML        string
	FormsDotGo         string
	GroupsDotGo        string
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
	PanelsDotGo        string
	PanelMapDotGo      string
	ChildParentDotGo   string
	RecordsDotGo       string
	ResizeDotGo        string
	ResizeSliderDotGo  string
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

	// v 13

	AttributesDotGo string
	CallBackDotGo   string
	CheckedDotGo    string
	DOMDotGo        string
	ElementDotGo    string
	EventDotGo      string
	FocusBlurDotGo  string
	HostPortDotGo   string
	// LCDataDotGo
	// MarkupDotGo
	MiscDotGo        string
	ScrollDotGo      string
	TextHTMLDotGo    string
	ValueDotGo       string
	WindowDotGo      string
	CloserDotGo      string
	HideShowDotGo    string
	IDDotGo          string
	DisplayGo        string
	PrintDotGo       string
	ProofsDotGo      string
	ApplicationDotGo string
}

// GetFileNames returns the file names.
func GetFileNames() *FileNames {
	return &FileNames{
		AppDotWASM:         "app.wasm",
		BuildDotSH:         "build.sh",
		BuildPackDotSH:     "buildPack.sh",
		ClassDotGo:         "class.go",
		ColorsDotCSS:       "colors.css",
		CreateGetDotGo:     "createGet.go",
		DocumentDotGo:      "document.go",
		ExampleDotTXT:      "example.txt",
		ExampleGoDotTXT:    "exampleGo.txt",
		EventsDotGo:        "events.go",
		FavIconDotICO:      "favicon.ico",
		FilePathsDotGo:     "filepaths.go",
		FlagDotYAML:        "flags.yaml",
		FormsDotGo:         "forms.go",
		GroupsDotGo:        "groups.go",
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
		PanelsDotGo:        "panels.go",
		PanelMapDotGo:      "panelMap.go",
		RecordsDotGo:       "Records.go",
		ResizeDotGo:        "resize.go",
		ResizeSliderDotGo:  "resizeSlider.go",
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

		// v 13
		AttributesDotGo:  "attributes.go",
		CallBackDotGo:    "callback.go",
		CheckedDotGo:     "checked.go",
		ChildParentDotGo: "childParent.go",
		DOMDotGo:         "dom.go",
		ElementDotGo:     "element.go",
		EventDotGo:       "event.go",
		FocusBlurDotGo:   "focusblur.go",
		HostPortDotGo:    "hostport.go",
		// LCDataDotGo:
		// MarkupDotGo:
		MiscDotGo:        "misc.go",
		ScrollDotGo:      "scroll.go",
		TextHTMLDotGo:    "texthtml.go",
		ValueDotGo:       "value.go",
		WindowDotGo:      "window.go",
		CloserDotGo:      "closer.go",
		HideShowDotGo:    "hideshow.go",
		IDDotGo:          "id.go",
		DisplayGo:        "display.go",
		PrintDotGo:       "print.go",
		ProofsDotGo:      "proofs.go",
		ApplicationDotGo: "Application.go",
	}
}
