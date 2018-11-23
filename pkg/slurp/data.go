package slurp

// Slurper slurps an application yaml file and checks it.
type Slurper struct {
	panelNames   map[string]string   // name : source path
	buttonIDs    map[string][]string // panel name : button names
	tabIDs       map[string][]string // panel name : tab names
	maxLevel     int
	CurrentLevel int
	panelFiles   []string
}

// NewSlurper constructs a new slurper.
func NewSlurper() *Slurper {
	return &Slurper{
		panelNames:   make(map[string]string),
		buttonIDs:    make(map[string][]string),
		tabIDs:       make(map[string][]string),
		maxLevel:     5,
		CurrentLevel: 2,
		panelFiles:   make([]string, 0, 5),
	}
}

// ApplicationInfo is info defining an application.
type ApplicationInfo struct {
	SourcePath string         `yaml:"sourcePath"`
	Title      string         `yaml:"title"`
	ImportPath string         `yaml:"importPath"`
	Stores     []string       `yaml:"stores"`
	Services   []*ServiceInfo `yaml:"services"`
}

// ServiceInfo is info defining a service.
type ServiceInfo struct {
	SourcePath string      `yaml:"sourcePath"`
	Name       string      `yaml:"name"`
	Button     *ButtonInfo `yaml:"button"`
}

// ButtonInfo is info defining a button.
type ButtonInfo struct {
	SourcePath string       `yaml:"sourcePath"`
	ID         string       `yaml:"name"`
	Label      string       `yaml:"label"`
	Heading    string       `yaml:"heading"`
	CC         string       `yaml:"cc"`
	PanelFiles []string     `yaml:"panelFiles,omitempty"`
	Panels     []*PanelInfo `yaml:"panels,omitempty"` // "-"
}

// TabInfo is info defining a tab.
type TabInfo struct {
	SourcePath string       `yaml:"sourcePath"`
	ID         string       `yaml:"name"`
	Label      string       `yaml:"label"`
	PanelFiles []string     `yaml:"panelFiles,omitempty"`
	Panels     []*PanelInfo `yaml:"panels,omitempty"` // "-"
}

// PanelInfo is info defining a panel.
type PanelInfo struct {
	SourcePath string `yaml:"sourcePath"`
	Level      int    `yaml:"level"`
	ID         string `yaml:"id"`
	Name       string `yaml:"name"`
	Note       string `yaml:"note"`

	Buttons []*ButtonInfo `yaml:"buttons"`
	Tabs    []*TabInfo    `yaml:"tabs"`
	Markup  string        `yaml:"markup,omitempty"`
}
