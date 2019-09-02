package templates

// SpawnTabPanelPanel is the genereric renderer spawn panel template.
const SpawnTabPanelPanel = `{{$Dot := .}}package {{call .PackageNameCase .PanelName}}

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*

	Panel name: {{.PanelName}}

*/

// spawnedPanel has a controller, presenter and caller.
type spawnedPanel struct {
	uniqueID    uint64
	tabButton   js.Value
	panelNameID map[string]string
	controller  *panelController
	presenter   *panelPresenter
	caller      *panelCaller
	group       *panelGroup
}

// newPanel constructs a new panel.
func newPanel(uniqueID uint64, tabButton js.Value, tabPanelHeader js.Value, panelNameID map[string]string, spawnData interface{}, unspawn func() error) (panel *spawnedPanel) {

	group := &panelGroup{
		uniqueID:    uniqueID,
		panelNameID: panelNameID,
	}
	controller := &panelController{
		group:    group,
		uniqueID: uniqueID,
		unspawn:  unspawn,
		eventCh:  make(chan viewtools.Event, 1024),
	}
	presenter := &panelPresenter{
		group:          group,
		uniqueID:       uniqueID,
		tabButton:      tabButton,
		tabPanelHeader: tabPanelHeader,
	}
	caller := &panelCaller{
		group:        group,
		uniqueID:     uniqueID,
		unspawn:      unspawn,
		UnSpawningCh: make(chan struct{}),
	}

	/* NOTE TO DEVELOPER. Step 1 of 1.

	// Use your custom spawnData.
	// If you have your own spawnData defined
	//   then you can use it here.
	// example:
	// Let's say that I define my spawn data types in renderer/spawndata/
	//   with the following definition.
	// type ChatRoomSpawnData struct {
	// 	   ServerName   string // Use for the panel heading.
	// 	   ChannelName  string // Use for the tab label and panel heading.
	// 	   ConnectionID string // The caller needs this.
	// }

	// import "{{.ApplicationGitPath}}{{.ImportRenderer}}/spawndata"

	switch spawnData := spawnData.(type) {
	case *spawndata.MySpawnData:
		caller.ircConnectionID = spawnData.ConnectionID
		presenter.serverName = spawnData.ServerName
		presenter.channelName = spawnData.ChannelName
	}

	*/

	panel = &spawnedPanel{
		uniqueID:   uniqueID,
		tabButton:  tabButton,
		controller: controller,
		presenter:  presenter,
		caller:     caller,
		group:      group,
	}

	controller.panel = panel
	controller.presenter = presenter
	controller.caller = caller
	presenter.controller = controller
	presenter.caller = caller
	caller.controller = controller
	caller.presenter = presenter

	return
}
`
