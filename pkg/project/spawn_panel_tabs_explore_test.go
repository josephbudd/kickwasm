package project

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

const spawnYamlString = `
name: TabsButton
label: Tabs
heading: Tabs
cc: Tabs
panels:
- id: TabsButtonTabBarPanel
  name: TabsButtonTabBarPanel
  tabs:
  - name: FirstTab
	label: First Tab
	heading: ""
	panels:
	- id: CreatePanel
	  name: CreatePanel
	  HasRealTabs: false
	  note: Button to create a new hello world.
	  markup: |
		<p>
		  <button id="newHelloWorldButton">New Hello World</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	spawn: false
  - name: SecondTab
	label: Second Tab
	heading: ""
	panels:
	- id: NotReadyTemplatePanel
	  name: NotReadyTemplatePanel
	  HasRealTabs: false
	  note: Click the ready button to switch to the hello world panel.
	  markup: |
		<p>Are you ready? {{.SpawnID}}!</p>
		<p>
			<button id="readySpawnButton{{.SpawnID}}">Ready</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	- id: HelloWorldTemplatePanel
	  name: HelloWorldTemplatePanel
	  HasRealTabs: false
	  note: Yet another "hello world".
	  markup: |
		<p id="p{{.SpawnID}}">Hello World {{.SpawnID}}!</p>
		<p>
			<button id="closeSpawnButton{{.SpawnID}}">Close</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	spawn: true
  - name: FirstOtherTab
	label: First OtherTab
	heading: ""
	panels:
	- id: CreateOtherPanel
	  name: CreateOtherPanel
	  HasRealTabs: false
	  note: Button to create a new hello world.
	  markup: |
		<p>
		  <button id="newHelloWorldButton">New Hello World</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	spawn: false
  - name: SecondOtherTab
	label: Second OtherTab
	heading: ""
	panels:
	- id: NotReadyTemplateOtherPanel
	  name: NotReadyTemplateOtherPanel
	  HasRealTabs: false
	  note: Click the ready button to switch to the hello world panel.
	  markup: |
		<p>Are you ready? {{.SpawnID}}!</p>
		<p>
			<button id="readySpawnButton{{.SpawnID}}">Ready</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	- id: HelloWorldTemplateOtherPanel
	  name: HelloWorldTemplateOtherPanel
	  HasRealTabs: false
	  note: Yet another "hello world".
	  markup: |
		<p id="p{{.SpawnID}}">Hello World {{.SpawnID}}!</p>
		<p>
			<button id="closeSpawnButton{{.SpawnID}}">Close</button>
		</p>
	  HVScroll: false
	  HTMLID: ""
	  TabBarHTMLID: ""
	  UnderTabBarHTMLID: ""
	  H3ID: ""
	spawn: true
  HasRealTabs: false
  note: ""
  HVScroll: false
  HTMLID: ""
  TabBarHTMLID: ""
  UnderTabBarHTMLID: ""
  H3ID: ""
`

func TestSpawn(t *testing.T) {
	builder := NewBuilder()
	builder.Name = "name"
	builder.Title = "title"
	var home *Button
	var err error
	if home, err = getHome("testyaml/spawn/home.yaml"); err != nil {
		t.Fatal(err)
	}
	spawnHomes := []*Button{home}
	if err = builder.BuildFromHomes(spawnHomes); err != nil {
		t.Fatal(err)
	}
	var bb []byte
	if bb, err = yaml.Marshal(*builder.Homes[0]); err != nil {
		t.Fatal(err)
	}
	bbstr := string(bb)
	if bbstr != spawnYamlString {
		lbb := len(bbstr)
		lw := len(spawnYamlString)
		if lbb < lw {
			for i := 0; i < lbb; i++ {
				j := i + 1
				if bbstr[i:j] != spawnYamlString[i:j] {
					t.Fatalf("bbstr[%d:%d] %q != spawnYamlString[%d:%d] %q", i, j, bbstr[i:j], i, j, spawnYamlString[i:j])
				}
			}
		}
	}
	testGenerateHomeSpawnTabEmptyInsidePanelNamePathMap(t, builder)
	testGenerateHomeSpawnTabButtonPanelGroups(t, builder)
}

func testGenerateHomeSpawnTabEmptyInsidePanelNamePathMap(t *testing.T, builder *Builder) {
	wants := map[string]map[string][]string{
		"TabsButton": map[string][]string{
			"HelloWorldTemplatePanel":      []string{"TabsButton", "TabsButtonTabBarPanel", "SecondTab", "HelloWorldTemplatePanel"},
			"NotReadyTemplatePanel":        []string{"TabsButton", "TabsButtonTabBarPanel", "SecondTab", "NotReadyTemplatePanel"},
			"HelloWorldTemplateOtherPanel": []string{"TabsButton", "TabsButtonTabBarPanel", "SecondOtherTab", "HelloWorldTemplateOtherPanel"},
			"NotReadyTemplateOtherPanel":   []string{"TabsButton", "TabsButtonTabBarPanel", "SecondOtherTab", "NotReadyTemplateOtherPanel"},
		},
	}
	var results map[string]map[string][]string
	results = builder.GenerateHomeSpawnTabEmptyInsidePanelNamePathMap()
	var found bool
	var rHome map[string][]string
	if rHome, found = results["TabsButton"]; !found {
		t.Fatalf("results is %#v", results)
	}
	var wHome map[string][]string
	wHome = wants["TabsButton"]
	if len(rHome) != len(wHome) {
		t.Logf("len(rHome) is %d, len(wHome) is %d", len(rHome), len(wHome))
		t.Fatalf("results is %#v", results)
	}
	var pName string
	var wFolders []string
	var rFolders []string
	for pName, wFolders = range wHome {
		if rFolders, found = rHome[pName]; !found {
			t.Logf("panel name %q not found in result.", pName)
			t.Fatalf("results is %#v", results)
		}
		wPath := filepath.Join(wFolders...)
		rPath := filepath.Join(rFolders...)
		if wPath != rPath {
			t.Logf("panel name %q: want %q got %q", pName, wPath, rPath)
			t.Fatalf("results is %#v", results)
		}
	}
}

func testGenerateHomeSpawnTabButtonPanelGroups(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string][]*TabPanelGroup
	}{
		{
			name: "first test",
			want: map[string][]*TabPanelGroup{
				"TabsButton": []*TabPanelGroup{
					&TabPanelGroup{
						TabBarID:     "",
						TabBarName:   "TabsButtonTabBarPanel",
						TabName:      "SecondTab",
						TabID:        "",
						TabLabel:     "Second Tab",
						PanelHeading: "",
						PanelNamesIDMap: map[string]*Panel{
							"HelloWorldTemplatePanel": &Panel{
								ID:                "HelloWorldTemplatePanel",
								Name:              "HelloWorldTemplatePanel",
								Tabs:              []*Tab(nil),
								HasRealTabs:       false,
								Buttons:           []*Button(nil),
								Note:              "Yet another \"hello world\".",
								Markup:            "<p id=\"p{{.SpawnID}}\">Hello World {{.SpawnID}}!</p>\n<p>\n    <button id=\"closeSpawnButton{{.SpawnID}}\">Close</button>\n</p>\n",
								HVScroll:          false,
								HTMLID:            "",
								TabBarHTMLID:      "",
								UnderTabBarHTMLID: "",
								H3ID:              "",
								Level:             0x0,
								Template:          "",
							},
							"NotReadyTemplatePanel": &Panel{
								ID:          "NotReadyTemplatePanel",
								Name:        "NotReadyTemplatePanel",
								Tabs:        []*Tab(nil),
								HasRealTabs: false,
								Buttons:     []*Button(nil),
								Note:        "Click the ready button to switch to the hello world panel.",
								Markup:      "<p>Are you ready? {{.SpawnID}}!</p>\n<p>\n    <button id=\"readySpawnButton{{.SpawnID}}\">Ready</button>\n</p>\n",
								HVScroll:    false, HTMLID: "",
								TabBarHTMLID:      "",
								UnderTabBarHTMLID: "",
								H3ID:              "",
								Level:             0x0,
								Template:          "",
							},
						},
					},
					&TabPanelGroup{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := builder.GenerateHomeSpawnTabButtonPanelGroups()
			if err := checkGenerateHomeSpawnTabButtonPanelGroups(tt.want, got); err != nil {
				t.Error("check error is " + err.Error())
			}
		})
	}
}

func checkGenerateHomeSpawnTabButtonPanelGroups(want, got map[string][]*TabPanelGroup) (err error) {
	var found bool
	var wantTabName string
	var gotTabPanelGroups []*TabPanelGroup
	var gotTabPanelGroup *TabPanelGroup
	var wantTabPanelGroups []*TabPanelGroup
	var wantTabPanelGroup *TabPanelGroup
	var i int
	for wantTabName, wantTabPanelGroups = range want {
		if gotTabPanelGroups, found = got[wantTabName]; !found {
			emsg := fmt.Sprintf("Tab name %q not found in got.", wantTabName)
			err = errors.New(emsg)
			return
		}
		if len(gotTabPanelGroups) != len(wantTabPanelGroups) {
			emsg := fmt.Sprintf("len(gotTabPanelGroups) %d != len(wantTabPanelGroups) %d.", len(gotTabPanelGroups), len(wantTabPanelGroups))
			err = errors.New(emsg)
			return
		}
		// match panel groups.
		found = false
		for i, wantTabPanelGroup = range wantTabPanelGroups {
			for _, gotTabPanelGroup = range gotTabPanelGroups {
				if wantTabPanelGroup.TabName == gotTabPanelGroup.TabName {
					found = true
					break
				}
			}
			if !found {
				emsg := fmt.Sprintf("i is %d: wantTabPanelGroup.TabName %q != found in gotTabPanelGroups.\n\n%#v\n\n", i, wantTabPanelGroup.TabName, *gotTabPanelGroup)
				err = errors.New(emsg)
				return
			}
			var wantPanelName string
			for wantPanelName = range wantTabPanelGroup.PanelNamesIDMap {
				if _, found = gotTabPanelGroup.PanelNamesIDMap[wantPanelName]; !found {
					emsg := fmt.Sprintf("gotTabPanelGroup.PanelNamesIDMap[%q] not found", wantPanelName)
					err = errors.New(emsg)
					return
				}
			}
		}
	}
	return
}
