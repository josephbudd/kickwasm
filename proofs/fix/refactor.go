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
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %w", err)
		}
	}()

	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

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
}

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
	rendererProcessCtxCancel()
}

// dispatchMessages dispatches LPC messages from the main process.
// It stops when it receives on the eoj channel.
func (messenger *panelMessenger) dispatchMessages() {
	go func() {
		for {
			select {
			case <-rendererProcessCtx.Done():
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

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
		if err != nil {
			log.Println("Refactor error: ", err.Error())
		} else {
			log.Println("no error")
		}
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
