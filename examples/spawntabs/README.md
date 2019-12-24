# Learn from the spawntabs example

## Add (spawn) and Remove (unspawn) tabs in a tab bar

### Tabs introduction

#### A panel with tabs

In a kickwasm.yaml file, a panel with tabs is renderered with a horizontal tab bar. One of the tabs in the tab bar is up front and the others are in back. That front tab is attached to the top of it's visible panel as you would expect.

A panel with tabs **must have at least one tab that is not a spawned tab**.

##### A tab and it's markup panels

Each tab only has one or more markup panels. Each markup panel

* Has it's own HTML template file that you can edit.
* Has it's own go package where you must provide functionality.

You do not program tabs because the framework does. You only program the markup panels.

##### A spawned tab and it's markup panels

In a kickwasm.yaml file, a tab with **spawn: true** is a spawned tab. A spawned tab does not exist and is not renderered in the tab bar until it is spawned. A spawned tab no longer exists and disappears from the tab bar when it is unspawned.

Each spawned tab has a unique ID. It passes that unique ID to each markup panel and corresponding template file that is has.

1. Each markup panel's HTML template file uses the spawn tab's unique ID. You will use **{{.SpawnID}}** especially for HTML element IDs. Example ID: **input{{.SpawnID}}**. Each occurance of **{{.SpawnID}}** in the HTML template file is replaced with the spawned tab's unique ID when the template is renderered in the GUI.
1. Each markup panel's go package uses the spawn tab's unique ID. So you will use it when you give functionality to the markup panel's go package.

### Example code

The panel **TabsButtonTabBarPanel** has 2 tabs.

1. **FirstTab** is the normal tab. It's markup panel has a button that will cause a **SecondTab** to be spawned.
1. **SecondTab** is a spawned tab. It's markup panel has a button that will close the spawned tab.

```yaml

title: Spawn Tabs
importPath: github.com/josephbudd/kickwasm/examples/spawntabs
buttons:
  - name: TabsButton
    label: Tabs
    panels:
      - name: TabsButtonTabBarPanel
        tabs:
          - name: FirstTab
            label: First Tab
            panels:
              - name: CreatePanel
                note: Button to create a new hello world.
                markup: |
                  <p>
                    <button id="newHelloWorldButton">New Hello World</button>
                  </p>
          - name: SecondTab
            spawn: true
            label: Second Tab
            panels:
              - name: HelloWorldTemplatePanel
                note: Yet another "hello world".
                markup: |
                  <p id="p{{.SpawnID}}">Hello World {{.SpawnID}}!</p>
                  <p>
                      <button id="closeSpawnButton{{.SpawnID}}">Close</button>
                  </p>


```

So after I created the framework I decided I wanted to edit the HTML template file for the HelloWorldTemplatePanel. That's what is nice about having template files for the markup panels. The edited file is shown below.

```html


<!--

Panel name: "HelloWorldTemplatePanel"

Panel note: Yet another "hello world".

This panel is displayed when the "Second Tab" tab button is clicked.

This panel is the only panel in it's panel group.

-->

<p id="p{{.SpawnID}}">Hello World {{.SpawnID}}!</p>
<p id="p2{{.SpawnID}}"></p>
<p>
    <label for="input{{.SpawnID}}">Tab Label</label>
    <input type="text" id="input{{.SpawnID}}"/>
    <br/>
    <button id="setButton{{.SpawnID}}">Set Label</button>
</p>
<p>
    <button id="closeSpawnButton{{.SpawnID}}">Close</button>
</p>

```

#### Constructor Data

For each spawn tab to be unique it must have some data that make it's unique. So I created the package **spawndata** where I declare the type **SecondTab** in the my new **spawndata** package at **examples/spawntabs/rendererprocess/spawndata/spawndata.go**. The file is shown below.

```go

package spawndata

// SecondTab is just an example how to create data for a spawned tab.
type SecondTab struct {
  Message string
}

```

I will pass it to the tab's spawn func when ever I call it.

#### It begins with the FirstTab's CreatePanel

The FirstTab is the normal tab that spawns the spawned tabs. The FirstTab's CreatePanel is the panel that causes the spawned tabs to be spawned. The FirstTab's CreatePanel is so simple that I only have to add functionaliry to it's controller so that it handles the **button#newHelloWorldButton** click. The button click handler is what spawns the spawned tabs. It is shown below.

##### Here are some things to note in the controller.go file shown below

1. The button click handler is **func (controller \*panelController) handleClick(e event.Event) (nilReturn interface{})**
1. The button click handler spawns a new **SecondTab** with the call **_, err := secondtab.Spawn(tabLabel, panelHeading, data)**.
   * Every **SecondTab** markup panel will need to be constructed with a **\*spawndata.SecondTab** data struct. It is the 3rd param passed in the call to **secondtab.Spawn**.
   * The call ignores the first param returned which is the spawned tab's unspawn func. That is because the spawned tab's panel will unspawn the tab when the user click a button.

```go

// +build js, wasm

package createpanel

import (
  "errors"
  "fmt"

  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/event"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/markup"
  secondtab "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/spawndata"
)

/*

  Panel name: CreatePanel

*/

// panelController controls user input.
type panelController struct {
  group     *panelGroup
  presenter *panelPresenter
  messenger *panelMessenger

  /* NOTE TO DEVELOPER. Step 1 of 4.

  // Declare your panelController fields.

  // example:

  import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/markup"

  addCustomerName   *markup.Element
  addCustomerSubmit *markup.Element

  */

  newHelloWorldButton *markup.Element
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

  // Define the submit button and set it's handler.
  if controller.newHelloWorldButton = document.ElementByID("newHelloWorldButton"); controller.newHelloWorldButton == nil {
    err = errors.New("unable to find #newHelloWorldButton")
    return
  }
  // Handle the button's onclick event.
  controller.newHelloWorldButton.SetEventHandler(controller.handleClick, "click", false)

  return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/store/record"
import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/event"
import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
  // See renderer/event/event.go.
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

func (controller *panelController) handleClick(e event.Event) (nilReturn interface{}) {
  spawnCount++
  n := spawnCount
  tabLabel := fmt.Sprintf("Tab %d", n)
  panelHeading := fmt.Sprintf("Panel Heading %d", n)
  data := &spawndata.SecondTab{
    Message: fmt.Sprintf("Message %d", n),
  }
  if _, err := secondtab.Spawn(tabLabel, panelHeading, data); err != nil {
    display.Error(err.Error())
  }
  return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

  /* NOTE TO DEVELOPER. Step 4 of 4.

  // Make the initial calls.
  // I use this to start up widgets. For example a virtual list widget.

  // example:

  controller.customerSelectWidget.start()

  */
}

```

#### The spawned tab's markup panels

##### func newPanel in Panel.go

The spawned tab is named **SecondTab**. It only has one panel named **HelloWorldTemplatePanel**. The panel's go package is located at **examples/spawntabs/rendererprocess/spawnPanels/TabsButton/TabsButtonTabBarPanel/SecondTab/HelloWorldTemplatePanel**.

The framework programs the tab but the developer has to program the markup panels. So with the **HelloWorldTemplatePanel** we begin with the packages **Panel.go** file and it's **func newPanel**.

**func newPanel** builds the entire go package for it's spawned markup panel. In this case I need to use the funcs **spawnData** to set the presenter's message that the presenter will display. The code from **func newPanel** in **Panel.go** is shown below.

```go

  data := spawnData.(*spawndata.SecondTab)
  presenter.message = data.Message
  
```

##### The panel presenter

The panel presenter displays it's message field which was set by the **func newPanel**. The file **Presenter.go** is shown below.

```go

// +build js, wasm

package helloworldtemplatepanel

import (
  "fmt"
  "strings"

  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/dom"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/markup"
)

/*

  Panel name: HelloWorldTemplatePanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
  uniqueID       uint64
  document       *dom.DOM
  group          *panelGroup
  controller     *panelController
  messenger      *panelMessenger
  tabButton      *markup.Element
  tabPanelHeader *markup.Element

  /* NOTE TO DEVELOPER: Step 1 of 3.

  // Declare your panelPresenter members here.

  // example:

  addCustomerName *markup.Element

  */

  message string
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

  defer func() {
    if err != nil {
      err = fmt.Errorf("(presenter *panelPresenter) defineMembers(): %w", err)
    }
  }()

  /* NOTE TO DEVELOPER. Step 2 of 3.

  // Define your panelPresenter members.

  // example:

  import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/framework/viewtools"

  // my spawn template has a name input field and a submit button.
  // <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">

  var id string

  // Define the customer name input field.
  // Build it's id using the uniqueID.
  id = display.SpawnID("addCustomerName{{.SpawnID}}", presenter.uniqueID)
  if presenter.addCustomerName = presenter.document.ElementByID(id); presenter.addCustomerName == nil {
    err = fmt.Errorf("unable to find #" + id)
    return
  }

  */

  var id string
  var p2 *markup.Element

  // Define the 2nd paragraph.
  // Build it's id using the uniqueID.
  id = display.SpawnID("p2{{.SpawnID}}", presenter.uniqueID)
  if p2 = presenter.document.ElementByID(id); p2 == nil {
    err = fmt.Errorf("unable to find #" + id)
    return
  }
  // Display the message in the 2nd paragraph.
  p2.SetInnerText(presenter.message)

  return
}

// Tab button label.

func (presenter *panelPresenter) getTabLabel() (label string) {
  label = presenter.tabButton.InnerText()
  return
}

func (presenter *panelPresenter) setTabLabel(label string) {
  presenter.tabButton.SetInnerText(label)
}

// Tab panel heading.

func (presenter *panelPresenter) getTabPanelHeading() (heading string) {
  heading = presenter.tabPanelHeader.InnerText()
  return
}

func (presenter *panelPresenter) setTabPanelHeading(heading string) {
  heading = strings.TrimSpace(heading)
  if len(heading) == 0 {
    presenter.tabPanelHeader.Hide()
  } else {
    presenter.tabPanelHeader.Show()
  }
  presenter.tabPanelHeader.SetInnerText(heading)
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/store/record"

// displayCustomer displays the customer in the add customer form.
// This short example only uses the customer name field in the form.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
  presenter.addCustomerName.SetValue(record.Name)
}

*/

```

##### The panel controller

The panel controller handles the panel's **button#closeSpawnButton{{.SpawnID}}** click event. The actual handler is shown below.

```go

func (controller *panelController) handleClick(e event.Event) (nilReturn interface{}) {
  if err := controller.unspawn(); err != nil {
    display.Error(err.Error())
  }
  return
}

```

The panel controller also handle the panels **button#setButton{{.SpawnID}}** click event by reading what the user typed and using it to set the tab's label. The actual handler is shown below.

```go

func (controller *panelController) handleSetClick(e event.Event) (nilReturn interface{}) {
  var text string
  if text = strings.TrimSpace(controller.input.Value()); len(text) == 0 {
    display.Error("Enter some text for the tab label.")
    return
  }
  controller.presenter.setTabLabel(text)
  return
}

```

The entire **Controller.go** file is shown below.

```go

// +build js, wasm

package helloworldtemplatepanel

import (
  "errors"
  "fmt"
  "strings"

  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/dom"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/event"
  "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/markup"
)

/*

  Panel name: HelloWorldTemplatePanel

*/

// panelController controls user input.
type panelController struct {
  uniqueID  uint64
  document  *dom.DOM
  panel     *spawnedPanel
  group     *panelGroup
  presenter *panelPresenter
  messenger *panelMessenger
  unspawn   func() error

  /* NOTE TO DEVELOPER. Step 1 of 5.

  // Declare your panelController members.

  // example:

  // my spawn template has a name input field and a submit button.
  // <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
  // <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

  import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/markup"

  addCustomerName   *markup.Element
  addCustomerSubmit *markup.Element

  */

  //closeSpawnButton{{.SpawnID}}
  closeSpawnButton *markup.Element
  //input{{.SpawnID}}
  input *markup.Element
  // setButton{{.SpawnID}}
  setButton *markup.Element
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

  defer func() {
    if err != nil {
      err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %w", err)
    }
  }()

  /* NOTE TO DEVELOPER. Step 2 of 5.

  // Define each controller in the GUI by it's html element.
  // Handle each controller's events.

  // example:

  import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"

  var id string

  // Define the customer name input field.
  id = display.SpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
  if controller.addCustomerName = contoller.document.ElementByID(id); controller.addCustomerName == nil {
    err = fmt.Errorf("unable to find #" + id)
    return
  }

  // Define the submit button.
  id = display.SpawnID("addCustomerSubmit{{.SpawnID}}", controller.uniqueID)
  if controller.addCustomerSubmit = contoller.document.ElementByID(id); controller.addCustomerSubmit == nil {
    err = fmt.Errorf("unable to find #" + id)
    return
  }
  // Handle the submit button's onclick event.
  controller.addCustomerSubmit.SetEventHandler(controller.handleSubmit, "click", false)

  */

  var id string

  // Define the self close button and set it's handler.
  id = display.SpawnID("closeSpawnButton{{.SpawnID}}", controller.uniqueID)
  if controller.closeSpawnButton = controller.document.ElementByID(id); controller.closeSpawnButton == nil {
    err = errors.New("unable to find #" + id)
    return
  }
  // Handle the close button's onclick event.
  controller.closeSpawnButton.SetEventHandler(controller.handleClick, "click", false)

  // Define the label input.
  id = display.SpawnID("input{{.SpawnID}}", controller.uniqueID)
  if controller.input = controller.document.ElementByID(id); controller.input == nil {
    err = errors.New("unable to find #" + id)
    return
  }

  // Define the label set button and set it's handler.
  id = display.SpawnID("setButton{{.SpawnID}}", controller.uniqueID)
  if controller.setButton = controller.document.ElementByID(id); controller.setButton == nil {
    err = errors.New("unable to find #" + id)
    return
  }
  // Handle the set button's onclick event.
  controller.setButton.SetEventHandler(controller.handleSetClick, "click", false)

  return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

// example:

import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/store/record"
import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/event"
import "github.com/josephbudd/kickwasm/examples/spawntabs/rendererprocess/api/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
  // See renderer/event/event.go.
  // The event.Event funcs.
  //   e.PreventDefaultBehavior()
  //   e.StopCurrentPhasePropagation()
  //   e.StopAllPhasePropagation()
  //   target := e.JSTarget
  //   event := e.JSEvent
  // You must use the javascript event e.JSEvent, as a js.Value.
  // However, you can use the target as a *markup.Element
  //   target := controller.document.NewElementFromJSValue(e.JSTarget)

  name := strings.TrimSpace(controller.addCustomerName.Value())
  if len(name) == 0 {
    display.Error("Customer Name is required.")
    return
  }
  r := &record.CustomerRecord{
    Name: name,
  }
  controller.messenger.AddCustomer(r)
  return
}

*/

func (controller *panelController) handleClick(e event.Event) (nilReturn interface{}) {
  if err := controller.unspawn(); err != nil {
    display.Error(err.Error())
  }
  return
}

func (controller *panelController) handleSetClick(e event.Event) (nilReturn interface{}) {
  var text string
  if text = strings.TrimSpace(controller.input.Value()); len(text) == 0 {
    display.Error("Enter some text for the tab label.")
    return
  }
  controller.presenter.setTabLabel(text)
  return
}

func (controller *panelController) UnSpawning() {

  /* NOTE TO DEVELOPER. Step 4 of 5.

  // This func is called when this tab and it's panels are in the process of unspawning.
  // So if you have some cleaning up to do then do it now.
  //
  // For example if you have a widget that needs to be unspawned
  //   because maybe it has a go routine running that needs to be stopped
  //   then do it here.

  // example:

  controller.myWidget.UnSpawn()

  */
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

  /* NOTE TO DEVELOPER. Step 5 of 5.

  // Make the initial calls.
  // I use this to start up widgets. For example a virtual list widget.

  // example:

  controller.customerSelectWidget.start()

  */
}

```