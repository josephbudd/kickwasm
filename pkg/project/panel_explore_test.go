package project

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExplore(t *testing.T) {
	builder := NewBuilder()
	builder.Name = "name"
	builder.Title = "title"
	var err error
	var okServices []*Service
	okServices, err = buildShortOkServices()
	if err != nil {
		t.Fatal(err)
	}
	err = builder.BuildFromServices(okServices)
	if err != nil {
		t.Fatal(err)
	}
	// ok service names
	testGenerateServiceNames(t, builder)
	// generate the html
	_, err = builder.ToHTML("masterid", false)
	if err != nil {
		t.Fatal(err)
	}
	testGenerateServiceEmptyPanelIDsMap(t, builder)
	testGenerateServiceEmptyInsidePanelIDsMap(t, builder)
	testGenerateTabBarLevelStartPanelMap(t, builder)
	testGenerateServiceButtonPanelGroups(t, builder)
	testgenerateServicePanelNameTemplateMap(t, builder)
	testGenerateServiceTemplatePanelName(t, builder)
	// deep
	var deepServices []*Service
	if deepServices, err = buildDeepServices(); err != nil {
		t.Fatal(err)
	}
	builder2 := NewBuilder()
	if err = builder2.BuildFromServices(deepServices); err != nil {
		t.Fatal(err)
	}
	if _, err = builder2.ToHTML("masterid", false); err != nil {
		t.Fatal(err)
	}
	testBuilderGenerateServiceEmptyInsidePanelNamePathMap(t, builder2)
	testBuilderGenerateServicePanelNamePanelMap(t, builder2)

	var longOkServices []*Service
	if longOkServices, err = buildLongOkServices(); err != nil {
		t.Fatal(err)
	}
	builder3 := NewBuilder()
	if err = builder3.BuildFromServices(longOkServices); err != nil {
		t.Fatal(err)
	}
	testGenerateButtonIDsPanelIDs(t, builder3)
	testGenerateTabIDsPanelIDs(t, builder3)
}

func testGenerateServiceButtonPanelGroups(t *testing.T, builder *Builder) {
	wantStr := `
Service1:
  - buttonname: OneButton
    buttonid: masterid-home-pad-OneButton
    panelnamesidmap:
      OnePanel:
        id: OnePanel
        name: OnePanel
        note: p1 note
        markup: <p>Panel 1-1</p>
        HTMLID: masterid-home-pad-OneButton-OnePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
      TwoPanel:
        id: TwoPanel
        name: TwoPanel
        note: p2 note
        markup: <p>Panel 2-1</p>
        HTMLID: masterid-home-pad-OneButton-TwoPanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
Service2:
  - buttonname: TwoButton
    buttonid: masterid-home-pad-TwoButton
    panelnamesidmap:
      FourPanel:
        id: fourPanel
        name: FourPanel
        note: p4 note
        markup: <p>Panel 4-1</p>
        HTMLID: masterid-home-pad-TwoButton-fourPanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
      ThreePanel:
        id: threePanel
        name: ThreePanel
        note: p3 note
        markup: <p>Panel 3-1</p>
        HTMLID: masterid-home-pad-TwoButton-threePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""`
	var wants map[string][]*ButtonPanelGroup
	if err := yaml.Unmarshal([]byte(wantStr), &wants); err != nil {
		t.Fatal(err)
	}
	results := builder.GenerateServiceButtonPanelGroups()
	//bb, _ := yaml.Marshal(results)
	//t.Fatal(string(bb))
	for service, wbpg := range wants {
		rbpg, ok := results[service]
		if !ok {
			t.Errorf(`GenerateServiceButtonPanelGroups service %s is missing`, service)
		} else {
			for i, wbp := range wbpg {
				rbp := rbpg[i]
				if rbp.ButtonName != wbp.ButtonName {
					t.Errorf(`rbpg[%d].ButtonName != wbp.ButtonName: got %s want %s`, i, rbp.ButtonName, wbp.ButtonName)
				}
				if rbpg[i].ButtonID != wbp.ButtonID {
					t.Errorf(`rbpg[%d].ButtonID != wbp.ButtonID: got %s want %s`, i, rbp.ButtonID, wbp.ButtonID)
				}
				for wpname := range wbp.PanelNamesIDMap {
					if _, ok := rbp.PanelNamesIDMap[wpname]; !ok {
						t.Errorf(`%s missing in rbpg[%d].PanelNamesIDMap`, wpname, i)
					}
				}
			}
		}
	}
}

func testGenerateTabBarLevelStartPanelMap(t *testing.T, builder *Builder) {
	wants := map[string]string{}
	results := builder.GenerateTabBarLevelStartPanelMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateTabBarLevelStartPanelMap() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateServiceEmptyInsidePanelIDsMap(t *testing.T, builder *Builder) {
	wants := map[string]map[string]string{
		"Service1": {
			"OnePanel": "masterid-home-pad-OneButton-OnePanel-inner-user-content",
			"TwoPanel": "masterid-home-pad-OneButton-TwoPanel-inner-user-content",
		},
		"Service2": {
			"FourPanel":  "masterid-home-pad-TwoButton-FourPanel-inner-user-content",
			"ThreePanel": "masterid-home-pad-TwoButton-ThreePanel-inner-user-content",
		},
	}
	results := builder.GenerateServiceEmptyInsidePanelIDsMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateServiceEmptyInsidePanelIDsMap() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateServiceEmptyPanelIDsMap(t *testing.T, builder *Builder) {
	wants := map[string][]string{
		"Service1": {
			"masterid-home-pad-OneButton-OnePanel",
			"masterid-home-pad-OneButton-TwoPanel",
		},
		"Service2": {
			"masterid-home-pad-TwoButton-ThreePanel",
			"masterid-home-pad-TwoButton-FourPanel",
		},
	}

	results := builder.GenerateServiceEmptyPanelIDsMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf("builder.GenerateServiceEmptyPanelIDsMap() generated %#v\n\nwant: %#v", results, wants)
	}
}

func testGenerateServiceNames(t *testing.T, builder *Builder) {
	correctNamesAnswer := []string{"Service1", "Service2"}
	serviceNames := builder.GenerateServiceNames()
	if len(serviceNames) != len(correctNamesAnswer) {
		t.Fatalf(`builder.GenerateServiceNames() len is %d not %d`, len(serviceNames), len(correctNamesAnswer))
	}
	for i, name := range serviceNames {
		if correctNamesAnswer[i] != name {
			t.Errorf(`builder.GenerateServiceNames() [%d] != %q its %q`, i, correctNamesAnswer[i], name)
		}
	}
}

func testgenerateServicePanelNameTemplateMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "a",
			want: map[string]map[string]string{
				"Service1": {
					"OnePanel": "\n<!--\n\nPanel name: \"OnePanel\"\n\nPanel note: p1 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-TwoPanel-inner-user-content\n  * Name: TwoPanel\n  * Note: p2 note\n\n-->\n\n<p>Panel 1-1</p>\n",
					"TwoPanel": "\n<!--\n\nPanel name: \"TwoPanel\"\n\nPanel note: p2 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-OnePanel-inner-user-content\n  * Name: OnePanel\n  * Note: p1 note\n\n-->\n\n<p>Panel 2-1</p>\n",
				},
				"Service2": {
					"ThreePanel": "\n<!--\n\nPanel name: \"ThreePanel\"\n\nPanel note: p3 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-FourPanel-inner-user-content\n  * Name: FourPanel\n  * Note: p4 note\n\n-->\n\n<p>Panel 3-1</p>\n",
					"FourPanel":  "\n<!--\n\nPanel name: \"FourPanel\"\n\nPanel note: p4 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-ThreePanel-inner-user-content\n  * Name: ThreePanel\n  * Note: p3 note\n\n-->\n\n<p>Panel 4-1</p>\n",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServicePanelNameTemplateMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServicePanelNameTemplateMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testGenerateServiceTemplatePanelName(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "wtf",
			want: map[string][]string{
				"Service2": {
					"ThreePanel",
					"FourPanel",
				},
				"Service1": {
					"OnePanel",
					"TwoPanel",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServiceTemplatePanelName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServiceTemplatePanelName() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilderGenerateServiceEmptyInsidePanelNamePathMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: map[string]map[string][]string{
				"Service1": {
					"OneOnePanel": {"OneButton", "OnePanel", "OneOneButton"},
					"TwoOnePanel": {"OneButton", "TwoPanel", "TwoOneButton"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateServiceEmptyInsidePanelNamePathMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateServiceEmptyInsidePanelNamePathMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilderGenerateServicePanelNamePanelMap(t *testing.T, builder *Builder) {
	wantStr := `
Service1:
  OneOnePanel:
    id: OneOnePanel
    name: OneOnePanel
    note: ""
    markup: <p>One One Panel</p>
    HTMLID: masterid-home-pad-OneButton-OnePanel-OneOneButton-OneOnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  OnePanel:
    id: OnePanel
    name: OnePanel
    buttons:
    - id: OneOneButton
      label: One One
      heading: One One
      cc: One One
      panels:
      - id: OneOnePanel
        name: OneOnePanel
        note: ""
        markup: <p>One One Panel</p>
        HTMLID: masterid-home-pad-OneButton-OnePanel-OneOneButton-OneOnePanel
        TabBarHTMLID: ""
        UnderTabBarHTMLID: ""
    note: p1 note
    HTMLID: masterid-home-pad-OneButton-OnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  TwoOnePanel:
    id: TwoOnePanel
    name: TwoOnePanel
    note: ""
    markup: <p>Two One Panel</p>
    HTMLID: masterid-home-pad-OneButton-TwoPanel-TwoOneButton-TwoOnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
  TwoPanel:
    id: TwoPanel
    name: TwoPanel
    buttons:
      - id: TwoOneButton
        label: Two One
        heading: Two One
        cc: Two One
        panels:
        - id: TwoOnePanel
          name: TwoOnePanel
          note: ""
          markup: <p>Two One Panel</p>
          HTMLID: masterid-home-pad-OneButton-TwoPanel-TwoOneButton-TwoOnePanel
          TabBarHTMLID: ""
          UnderTabBarHTMLID: ""
    note: p2 note
    HTMLID: masterid-home-pad-OneButton-TwoPanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""`
	var want map[string]map[string]*Panel
	if err := yaml.Unmarshal([]byte(wantStr), &want); err != nil {
		t.Fatal(err)
	}
	got := builder.GenerateServicePanelNamePanelMap()
	for wService, wPanelNamePanelMap := range want {
		gPanelNamePanelMap, ok := got[wService]
		if !ok {
			t.Errorf("service %q not found in result.", wService)
			return
		}
		for wName, wPanel := range wPanelNamePanelMap {
			gPanel, ok := gPanelNamePanelMap[wName]
			if !ok {
				t.Errorf("panel %q not found in result.", wName)
				return
			}
			if wPanel.ID != gPanel.ID {
				t.Errorf("panel.ID %q not found in result.", wPanel.ID)
				return
			}
		}

	}
}

func testGenerateButtonIDsPanelIDs(t *testing.T, builder *Builder) {
	wantButtons := map[string][]string{
		"FourButton":  {"OtherTabPanel"},
		"OneButton":   {"OnePanel", "TwoPanel"},
		"ThreeButton": {"TabPanel"},
		"TwoButton":   {"ThreePanel", "FourPanel"},
	}
	t.Run("testGenerateButtonIDsPanelIDs", func(t *testing.T) {
		if gotButtons := builder.GenerateButtonIDsPanelIDs(); !reflect.DeepEqual(gotButtons, wantButtons) {
			t.Errorf("Builder.GenerateButtonIDsPanelIDs() = %#v, want %#v", gotButtons, wantButtons)
		}
	})
}

func testGenerateTabIDsPanelIDs(t *testing.T, builder *Builder) {
	wantTabs := map[string][]string{
		"OtherTabPanel-OneTab": {"LastTabPanel"},
		"TabPanel-OneTab":      {"OnlyPanel"},
	}
	t.Run("testGenerateTabIDsPanelIDs", func(t *testing.T) {
		if gotTabs := builder.GenerateTabIDsPanelIDs(); !reflect.DeepEqual(gotTabs, wantTabs) {
			t.Errorf("Builder.GenerateTabIDsPanelIDs() = %#v, want %#v", gotTabs, wantTabs)
		}
	})
}
