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
		unSpawningCh: make(chan struct{}),
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
		unSpawningCh: make(chan struct{}),
	}

	/* NOTE TO DEVELOPER. Step 1 of 2.

	// param spawnData interface{}.

	// If you have your own type spawnData defined
	//   then you can use it here.
	//
	// Example:
	//
	// * Let's say that I define my spawn data types
	//     in my own new folder at renderer/spawndata/
	//     with the following definition.
	//   type JoinedChatRoomSpawnData struct {
	// 	     ServerName   string // Use for the panel heading.
	// 	     ChannelName  string // Use for the tab label and panel heading.
	// 	     ConnectionID string // The caller needs this.
	//   }
	// * Let's say that this panel is an IRC chat room panel
	//     where the user can read the chat room conversation log and write into
	//     the conversation.
	// * Let's say that the main process joined the user into chat room
	//     and now it is sending the information back to the renderer process
	//     so that it can open a chat room panel to allow the user to
	//     read and send chat room text.
	// * So when I spawned this panel I put the chat room information
	//     from the main process into a *spawndata.JoinedChatRoomSpawnData.
	// * Below is how I could use the *spawndata.JoinedChatRoomSpawnData here
	//     in this constructor as I build this panel package.
	
	// import "{{.ApplicationGitPath}}{{.ImportRenderer}}/spawndata"

	switch spawnData := spawnData.(type) {
	case *spawndata.JoinedChatRoomSpawnData:
		caller.ircConnectionID = spawnData.ConnectionID
		presenter.serverName = spawnData.ServerName
		presenter.channelName = spawnData.ChannelName
	}

	*/

	/* NOTE TO DEVELOPER. Step 2 of 2.

	// var help.

	// This package's var help in Data.go is a pointer to the renderer/paneling.Help.
	// If you redefined paneling.Help in renderer/paneling/Helping.go,
	//   then you may need to use it here.
	// Set any controller, presenter or caller members that you added.
	// Below is an example of me using help to set the caller's state.
	//
	// Example:

	caller.state = help.GetStateAdd()

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
