// +build js, wasm

package provebuttonpanel

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/api/markup"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type panelGroup struct {
	proveButtonPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(group *panelGroup) defineMembers(): %w", err)
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-ProveButton-ProveButtonPanel"); panel == nil {
	err = fmt.Errorf("unable to find #mainMasterView-home-pad-ProveButton-ProveButtonPanel")
		return
    }
    group.proveButtonPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showProveButtonPanel shows the panel you named ProveButtonPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-ProveButton-ProveButtonPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
The panel for the prove button.
Run the tests from this panel.
Do not move the ProveButtonPanel from the ProveButton.

*/
func (group *panelGroup) showProveButtonPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.proveButtonPanel, force)
}
