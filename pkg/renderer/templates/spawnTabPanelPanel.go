package templates

// SpawnTabPanelPanel is the genereric renderer spawn panel template.
const SpawnTabPanelPanel = `{{$Dot := .}}// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"context"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIDOM}}"
	"{{.ApplicationGitPath}}{{.ImportRendererAPIMarkup}}"
)

/*

	Panel name: {{.PanelName}}

*/

// spawnedPanel has a controller, presenter and messenger.
type spawnedPanel struct {
	uniqueID    uint64
	tabButton   *markup.Element
	panelNameID map[string]string
	controller  *panelController
	presenter   *panelPresenter
	messenger   *panelMessenger
	group       *panelGroup
}

// newPanel constructs a new panel.
func newPanel(ctx context.Context, ctxCancel context.CancelFunc, uniqueID uint64, tabButton, tabPanelHeader *markup.Element, panelNameID map[string]string, spawnData interface{}) (panel *spawnedPanel) {

	document := dom.NewDOM(uniqueID)
	group := &panelGroup{
		uniqueID:    uniqueID,
		document:    document,
		panelNameID: panelNameID,
	}
	controller := &panelController{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		group:     group,
		uniqueID:  uniqueID,
		document:  document,
	}
	presenter := &panelPresenter{
		group:          group,
		uniqueID:       uniqueID,
		document:       document,
		tabButton:      tabButton,
		tabPanelHeader: tabPanelHeader,
	}
	messenger := &panelMessenger{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		group:     group,
		uniqueID:  uniqueID,
	}

	/* NOTE TO DEVELOPER. Step 1 of 2.

	// param spawnData interface{}.

	// If you have your own type spawnData defined
	//   then you can use it here.
	//
	// example:
	//
	// * Let's say that I define my spawn data types
	//     in my own new folder at renderer/spawndata/
	//     with the following definition.
	//   type JoinedChatRoomSpawnData struct {
	// 	     ServerName   string // Use for the panel heading.
	// 	     ChannelName  string // Use for the tab label and panel heading.
	// 	     ConnectionID string // The messenger needs this.
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

	import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/spawndata"

	data := spawnData.(*spawndata.JoinedChatRoomSpawnData)
	messenger.ircConnectionID = data.ConnectionID
	presenter.serverName = data.ServerName
	presenter.channelName = data.ChannelName

	*/

	/* NOTE TO DEVELOPER. Step 2 of 2.

	// var help.

	// This package's var help in Data.go is a pointer to the renderer/paneling.Help.
	// If you redefined paneling.Help in renderer/paneling/Helping.go,
	//   then you may need to use it here.
	// Set any controller, presenter or messenger members that you added.
	// Below is an example of me using help to set the messenger's state.
	//
	// Example:

	messenger.state = help.GetStateIRCChannel()

	*/

	panel = &spawnedPanel{
		uniqueID:   uniqueID,
		tabButton:  tabButton,
		controller: controller,
		presenter:  presenter,
		messenger:  messenger,
		group:      group,
	}

	controller.panel = panel
	controller.presenter = presenter
	controller.messenger = messenger
	presenter.controller = controller
	presenter.messenger = messenger
	messenger.controller = controller
	messenger.presenter = presenter

	return
}
`
