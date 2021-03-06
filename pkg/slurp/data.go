package slurp

// Slurper slurps an application yaml file and checks it.
type Slurper struct {
	panelNames   map[string]string   // name : source path
	tabNames     map[string]string   // name : source path
	buttonNames  map[string]string   // name : source path
	buttonIDs    map[string][]string // panel name : button names
	tabIDs       map[string][]string // panel name : tab names
	maxLevel     int
	CurrentLevel int
	panelFiles   []string
}

// NewSlurper constructs a new slurper.
func NewSlurper() *Slurper {
	return &Slurper{
		panelNames:   make(map[string]string, 100),
		tabNames:     make(map[string]string, 100),
		buttonNames:  make(map[string]string, 100),
		buttonIDs:    make(map[string][]string, 100),
		tabIDs:       make(map[string][]string, 100),
		maxLevel:     5,
		CurrentLevel: 2,
		panelFiles:   make([]string, 0, 5),
	}
}

// ApplicationInfo is info defining an application.
type ApplicationInfo struct {
	SourcePath string        `yaml:"sourcePath"`
	Title      string        `yaml:"title"`
	ImportPath string        `yaml:"importPath"`
	Homes      []*ButtonInfo `yaml:"buttons"`
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
	Heading    string       `yaml:"heading"`
	PanelFiles []string     `yaml:"panelFiles,omitempty"`
	Panels     []*PanelInfo `yaml:"panels,omitempty"` // "-"
	Spawn      bool         `yaml:"spawn,omitempty"`
}

// PanelInfo is info defining a panel.
type PanelInfo struct {
	SourcePath string `yaml:"sourcePath"`
	Level      int    `yaml:"level"`
	ID         string `yaml:"id"`
	Name       string `yaml:"name"`

	Buttons  []*ButtonInfo `yaml:"buttons"`
	Tabs     []*TabInfo    `yaml:"tabs"`
	Markup   string        `yaml:"markup,omitempty"`
	Note     string        `yaml:"note"`
	HVScroll bool          `yaml:"hvscroll"`
}
