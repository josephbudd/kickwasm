package message

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain"
	"github.com/josephbudd/kickwasm/pkg/mainprocess"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer"
)

// Manager manages lpc messages.
type Manager struct {
	messageFolder string
	importGitPath string
	appPaths      paths.ApplicationPathsI
	log           string
	init          string
	instructions  string
}

// NewManager constructs a new Message.
func NewManager(pwd string, appName string, importGitPath string) (messages *Manager) {
	appPaths := &paths.ApplicationPaths{}
	appPaths.Initialize(pwd, "", appName)
	paths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	logParts := strings.Split(fileNames.LogDotGo, ".")
	initParts := strings.Split(fileNames.InitDotGo, ".")
	instructionsParts := strings.Split(fileNames.InstructionsDotTXT, ".")
	return &Manager{
		messageFolder: paths.OutputDomainLPCMessage,
		importGitPath: importGitPath,
		appPaths:      appPaths,
		log:           logParts[0],
		init:          initParts[0],
		instructions:  instructionsParts[0],
	}
}

// List returns the filtered list of message names.
func (mngr *Manager) List() (defaults, normals []string, err error) {
	if defaults, err = mngr.defaultList(); err != nil {
		return
	}
	normals, err = mngr.filteredList()
	return
}

// Add add a new message.
func (mngr *Manager) Add(messagename string) (err error) {
	var messageName string
	if messageName, err = mngr.checkName(messagename); err != nil {
		return
	}
	fName := messageName + ".go"
	// Check the file name.
	var files []string
	if files, err = mngr.fileNames(); err != nil {
		return
	}
	for _, name := range files {
		if fName == name {
			errmsg := fmt.Sprintf("the message %q already exists", messageName)
			err = errors.New(errmsg)
			return
		}
	}
	// Create the domain/message file.
	if err = domain.CreateCustomLPC(mngr.appPaths, messageName); err != nil {
		return
	}
	// Create the mainprocess/dispatch/ message file.
	if err = mainprocess.CreateCustomDispatch(mngr.appPaths, mngr.importGitPath, messageName); err != nil {
		return
	}
	err = mngr.rebuild()
	return
}

// Del removes a message.
func (mngr *Manager) Del(messagename string) (err error) {

	errList := make([]string, 0, 2)
	defer func() {
		if len(errList) > 0 {
			errmsg := strings.Join(errList, "\n")
			err = errors.New(errmsg)
		}
	}()

	var messageName string
	if messageName, err = mngr.checkName(messagename); err != nil {
		return
	}
	// Check the file name.
	var files []string
	if files, err = mngr.filteredList(); err != nil {
		return
	}
	var found bool
	for _, f := range files {
		if f == messageName {
			found = true
		}
	}
	if !found {
		errmsg := fmt.Sprintf("the message %q does not exist", messageName)
		err = errors.New(errmsg)
		return
	}
	// Delete the domain/message file.
	if err = domain.DeleteCustomLPC(mngr.appPaths, messageName); err != nil {
		errList = append(errList, err.Error())
	}
	// Delete the mainprocess/dispatch/ message file.
	if err = mainprocess.DeleteCustomDispatch(mngr.appPaths, messageName); err != nil {
		errList = append(errList, err.Error())
	}
	err = mngr.rebuild()
	return
}

// rebuild rebuilds the main process dispatcher, instructions.
// rebuilds the domain instructions.
// rebuilds the renderer/lpc/channels.go.
func (mngr *Manager) rebuild() (err error) {
	var list []string
	if list, err = mngr.filteredList(); err != nil {
		return
	}
	lpcNames := make([]string, len(list))
	for i, fname := range list {
		parts := strings.Split(fname, ".")
		lpcNames[i] = parts[0]
	}
	if err = mainprocess.RebuildDispatchDotGo(mngr.appPaths, mngr.importGitPath, lpcNames); err != nil {
		return
	}
	if err = mainprocess.RebuildLPCChannels(mngr.appPaths, mngr.importGitPath, lpcNames); err != nil {
		return
	}
	if err = domain.RebuildDomainLPCInstructions(mngr.appPaths, mngr.importGitPath, lpcNames); err != nil {
		return
	}
	err = renderer.RebuildChannelsDotGo(mngr.appPaths, mngr.importGitPath, lpcNames)
	return
}

// defaultList returns the filtered list of default names.
func (mngr *Manager) defaultList() (messages []string, err error) {
	var fnames []string
	if fnames, err = mngr.fileNames(); err != nil {
		return
	}
	names := make([]string, 0, len(fnames))
	for _, fname := range fnames {
		parts := strings.Split(fname, ".")
		name := parts[0]
		if name == mngr.init || name == mngr.log {
			names = append(names, name)
		}
	}
	sorted := sort.StringSlice(names)
	sorted.Sort()
	messages = sorted
	return
}

// filteredList returns the filtered list of message names.
func (mngr *Manager) filteredList() (messages []string, err error) {
	var fnames []string
	if fnames, err = mngr.fileNames(); err != nil {
		return
	}
	names := make([]string, 0, len(fnames))
	for _, fname := range fnames {
		parts := strings.Split(fname, ".")
		name := parts[0]
		if name != mngr.init && name != mngr.log && name != mngr.instructions {
			names = append(names, name)
		}
	}
	sorted := sort.StringSlice(names)
	sorted.Sort()
	messages = sorted
	return
}

// fileNames returns the names of every file in the domain/lpc/message/ folder.
func (mngr *Manager) fileNames() (messages []string, err error) {

	var dir *os.File
	defer func() {
		cerr := dir.Close()
		if err == nil {
			err = cerr
		}
	}()

	// open the message dir.
	if dir, err = os.Open(mngr.messageFolder); err != nil {
		return
	}
	messages, err = dir.Readdirnames(-1)
	return
}

func (mngr *Manager) checkName(messagename string) (messageName string, err error) {
	if messageName = cases.CamelCase(messagename); messageName != messagename {
		msg := fmt.Sprintf("Message names are CamelCased so %q should be %q", messagename, messageName)
		err = errors.New(msg)
	}
	if messageName == mngr.log || messageName == mngr.init {
		err = errors.New(messageName + " is one of the default message names")
	}
	return
}
