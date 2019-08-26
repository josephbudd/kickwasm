package project

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

const spawnYamlString = `name: Tabs
button:
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
	var service *Service
	var err error
	if service, err = getService("testyaml/spawn/service.yaml"); err != nil {
		t.Fatal(err)
	}
	spawnServices := []*Service{service}
	if err = builder.BuildFromServices(spawnServices); err != nil {
		t.Fatal(err)
	}
	var bb []byte
	if bb, err = yaml.Marshal(*builder.Services[0]); err != nil {
		t.Fatal(err)
	}
	if string(bb) != spawnYamlString {
		t.Fatal(string(bb) != spawnYamlString)
	}

	testGenerateServiceSpawnTabEmptyInsidePanelNamePathMap(t, builder)
	testGenerateServiceSpawnTabButtonPanelGroups(t, builder)
}

func testGenerateServiceSpawnTabEmptyInsidePanelNamePathMap(t *testing.T, builder *Builder) {
	wants := map[string]map[string][]string{
		"Tabs": map[string][]string{
			"HelloWorldTemplatePanel":      []string{"TabsButton", "TabsButtonTabBarPanel", "SecondTab", "HelloWorldTemplatePanel"},
			"NotReadyTemplatePanel":        []string{"TabsButton", "TabsButtonTabBarPanel", "SecondTab", "NotReadyTemplatePanel"},
			"HelloWorldTemplateOtherPanel": []string{"TabsButton", "TabsButtonTabBarPanel", "SecondOtherTab", "HelloWorldTemplateOtherPanel"},
			"NotReadyTemplateOtherPanel":   []string{"TabsButton", "TabsButtonTabBarPanel", "SecondOtherTab", "NotReadyTemplateOtherPanel"},
		},
	}
	var results map[string]map[string][]string
	results = builder.GenerateServiceSpawnTabEmptyInsidePanelNamePathMap()
	var found bool
	var rService map[string][]string
	if rService, found = results["Tabs"]; !found {
		t.Fatalf("results is %#v", results)
	}
	var wService map[string][]string
	wService = wants["Tabs"]
	if len(rService) != len(wService) {
		t.Logf("len(rService) is %d, len(wService) is %d", len(rService), len(wService))
		t.Fatalf("results is %#v", results)
	}
	var pName string
	var wFolders []string
	var rFolders []string
	for pName, wFolders = range wService {
		if rFolders, found = rService[pName]; !found {
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

func testGenerateServiceSpawnTabButtonPanelGroups(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string][]*TabPanelGroup
	}{
		{
			name: "first test",
			want: map[string][]*TabPanelGroup{
				"Tabs": []*TabPanelGroup{
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
			got := builder.GenerateServiceSpawnTabButtonPanelGroups()
			if err := checkGenerateServiceSpawnTabButtonPanelGroups(tt.want, got); err != nil {
				t.Error("check error is " + err.Error())
			}
		})
	}
}

func checkGenerateServiceSpawnTabButtonPanelGroups(want, got map[string][]*TabPanelGroup) (err error) {
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
