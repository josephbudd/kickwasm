package fix

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/proofs/common"
	"github.com/josephbudd/kickwasm/tools/script"
)

const (
	controllerDotGo = `// +build js, wasm

package provebuttonpanel

import (
	"fmt"
	"github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/prove"
)
		
/*

	Panel name: ProveButtonPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	// example:

	import "syscall/js"
	import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/api/markup"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

	*/
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %w", err)
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	// Define the customer name text input GUI controller.
	if controller.addCustomerName = document.ElementByID("addCustomerName"); controller.addCustomerName == nil {
		err = fmt.Errorf("unable to find #addCustomerName")
		return
	}

	// Define the submit button GUI controller.
	if controller.addCustomerSubmit = document.ElementByID("addCustomerSubmit"); controller.addCustomerSubmit == nil {
		err = fmt.Errorf("unable to find #addCustomerSubmit")
		return
	}
	// Handle the submit button's onclick event.
	controller.addCustomerSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/store/record"
import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/api/event"
import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/api/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	// See rendererprocess/api/event/event.go.
	// The event.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.JSTarget
	//   event := e.JSEvent
	// You must use the javascript event e.JSEvent, as a js.Value.
	// However, you can use the target as a *markup.Element
	//   target := document.NewElementFromJSValue(e.JSTarget)

	name := strings.TrimSpace(controller.addCustomerName.Value())
	if len(name) == 0 {
		display.Error("Customer Name is required.")
		return
	}
	r := &record.Customer{
		Name: name,
	}
	controller.messenger.AddCustomer(r)
	return
}

*/

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/

	if err := prove.Pass(); err != nil {
		controller.messenger.LogFail(err)
	} else {
		controller.messenger.logPass()
	}
}
`

	messengerDotGo = `// +build js, wasm

package provebuttonpanel

import (
	"strings"

	"github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/data/loglevels"
	"github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/lpc/message"
)

/*

	Panel name: ProveButtonPanel

*/

// panelMessenger communicates with the main process via an asynchrounous connection.
type panelMessenger struct {
	group      *panelGroup
	presenter  *panelPresenter
	controller *panelController

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1.1: Declare your panelMessenger members.

	// example:

	state uint64

	*/
}

/* NOTE TO DEVELOPER. Step 2 of 4.

// 2.1: Define your funcs which send a message to the main process.
// 2.2: Define your funcs which receive a message from the main process.

// example:

import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/store/record"
import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/lpc/message"
import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/rendererprocess/api/display"

// Add Customer.

func (messenger *panelMessenger) addCustomer(r *record.Customer) {
	msg := &message.AddCustomerRendererToMainProcess{
		UniqueID: messenger.uniqueID,
		Record:   record,
	}
	sendCh <- msg
}

func (messenger *panelMessenger) addCustomerRX(msg *message.AddCustomerMainProcessToRenderer) {
	if msg.UniqueID == messenger.uniqueID {
		if msg.Error {
			display.Error(msg.ErrorMessage)
			return
		}
		// no errors
		display.Success("Customer Added.")
	}
}

*/

func (messenger *panelMessenger) LogFail(err error) {
	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelError,
		Message: "Failed %[1]s test.\n"+strings.ReplaceAll(err.Error(), "\n", "<br/>"),
	}
	sendCh <- msg
}

func (messenger *panelMessenger) logPass() {
	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "%[2]s Passed test.",
	}
	sendCh <- msg
}

func (messenger *panelMessenger) rxLog(msg *message.LogMainProcessToRenderer) {
	quitCh <- struct{}{}
}

// dispatchMessages dispatches LPC messages from the main process.
// It stops when it receives on the eoj channel.
func (messenger *panelMessenger) dispatchMessages() {
	go func() {
		for {
			select {
			case <-eojCh:
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

				/* NOTE TO DEVELOPER. Step 3 of 4.

				// 3.1:   Remove the default clause below.
				// 3.2.a: Add a case for each of the messages
				//          that you are expecting from the main process.
				// 3.2.b: In that case statement, pass the message to your message receiver func.

				// example:

				import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/lpc/message"

				case *message.AddCustomerMainProcessToRenderer:
					messenger.addCustomerRX(msg)

				*/

				case *message.LogMainProcessToRenderer:
					messenger.rxLog(msg)
				}
			}
		}
	}()

	return
}

// initialSends sends the first messages to the main process.
func (messenger *panelMessenger) initialSends() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4.1: Send messages to the main process right when the app starts.

	// example:

	import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/data/loglevels"
	import "github.com/josephbudd/kickwasm/proofs/%[1]s/%[1]stest/domain/lpc/message"

	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	sendCh <- msg

	*/
}
`
)

var (
	controlerRelPath   = filepath.Join("rendererprocess", "panels", "ProveButton", "ProveButtonPanel", "Controller.go")
	messengerRelPath   = filepath.Join("rendererprocess", "panels", "ProveButton", "ProveButtonPanel", "Messenger.go")
	proveRelFolderPath = filepath.Join("rendererprocess", "prove")
)

func yamlCode(appName, formatString string) (contents string) {
	contents = fmt.Sprintf(formatString, appName)
	return
}

func goCode(appName, desc, formatString string) (contents string) {
	contents = fmt.Sprintf(formatString, appName, desc)
	return
}

// Refactor does the refactoring.
func Refactor(appName, description, sourceCodeFolderPath, kickwasmDotYAML, rekickwasmDotYAML, proveDotGo string, undo, testing bool) (err error) {

	executable := common.Executable("." + string(os.PathSeparator) + filepath.Base(sourceCodeFolderPath))
	defer func() {
		if !testing {
			// remove the source code folders.
			os.RemoveAll(sourceCodeFolderPath + "sitepack")
			os.RemoveAll(sourceCodeFolderPath)
		}
		// Remove the executable.
		rmdir := common.Executable(filepath.Dir(sourceCodeFolderPath))
		rmf := common.Executable(filepath.Base(rmdir))
		rmfpath := filepath.Join(rmdir, rmf)
		log.Printf(rmfpath)
		os.Remove(rmfpath)
	}()

	// Make the source code folder.
	if common.PathFound(sourceCodeFolderPath) {
		if err = os.RemoveAll(sourceCodeFolderPath); err != nil {
			return
		}
	}
	if err = common.MkDir(sourceCodeFolderPath); err != nil {
		return
	}
	// Write the kickwasm.yaml file and build the framework.
	path := filepath.Join(sourceCodeFolderPath, common.KickwasmDotYAML)
	if err = common.Write(path, yamlCode(appName, kickwasmDotYAML)); err != nil {
		return
	}
	var dump []byte
	if testing {
		log.Printf("script.RunDump(nil, %q, \"kickwasm\")\n", sourceCodeFolderPath)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "kickwasm"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		return
	}
	if testing {
		log.Printf("script.RunDump(nil, %q, \"kickbuild\", \"-rp\", \"-mp\")\n", sourceCodeFolderPath)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "kickbuild", "-rp", "-mp"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		return
	}
	// Add prove.go which will verify change correctness after the refactoring.
	path = filepath.Join(sourceCodeFolderPath, proveRelFolderPath)
	if err = common.MkDir(path); err != nil {
		return
	}
	path = filepath.Join(path, "prove.go")
	if err = common.Write(path, goCode(appName, description, proveDotGo)); err != nil {
		return
	}
	// Edit the controler which will run prove.Pass() after the refactoring.
	path = filepath.Join(sourceCodeFolderPath, controlerRelPath)
	if err = common.Write(path, goCode(appName, description, controllerDotGo)); err != nil {
		return
	}
	// Edit the messenger which will send the fatal to the main process.
	path = filepath.Join(sourceCodeFolderPath, messengerRelPath)
	if err = common.Write(path, goCode(appName, description, messengerDotGo)); err != nil {
		return
	}

	// Refactor the framework.
	// Initialize the refactoring.
	if testing {
		log.Println(`script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-i", "-dne")`)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-i", "-dne"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		return
	}
	// Edit ./rekickwasm/edit/yaml/kickwasm.yaml
	path = common.RekickwasmDotYAMLEditPath(sourceCodeFolderPath)
	if err = common.Write(path, yamlCode(appName, rekickwasmDotYAML)); err != nil {
		return
	}
	// Refactor.
	if testing {
		log.Println(`script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-R")`)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-R"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		return
	}
	if undo {
		if testing {
			log.Println(`script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-u")`)
		}
		if dump, err = script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-u"); err != nil {
			return
		}
	}
	// Clear
	if testing {
		log.Println(`script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-x")`)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "rekickwasm", "-x"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		return
	}
	// Build the proofs verifying that the refactor is correct.
	if testing {
		log.Println(`script.RunDump(nil, sourceCodeFolderPath, "kickbuild", "-rp", "-mp")`)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, "kickbuild", "-rp", "-mp"); err != nil {
		if testing {
			log.Println(string(dump))
		}
		if err == nil && len(dump) > 0 {
			err = fmt.Errorf(appName + " kickbuild error")
		}
		return
	}
	// Run the proofs verifying that the refactor is correct.
	// Make the executable file name.
	if testing {
		log.Println(sourceCodeFolderPath)
	}
	if testing {
		log.Printf(`script.RunDump(nil, %q, %q)`, sourceCodeFolderPath, executable)
	}
	if dump, err = script.RunDump(nil, sourceCodeFolderPath, executable); err != nil {
		if err == nil && len(dump) > 0 {
			err = fmt.Errorf(appName + " run error")
		}
	}
	log.Println(string(dump))
	return
}
