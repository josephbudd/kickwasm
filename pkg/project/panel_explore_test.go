package project

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExplore(t *testing.T) {
	builder := NewBuilder()
	builder.Name = "name"
	builder.Title = "title"
	var err error
	var okHomes []*Button
	okHomes, err = buildShortOkHomes()
	if err != nil {
		t.Fatal(err)
	}
	err = builder.BuildFromHomes(okHomes)
	if err != nil {
		t.Fatal(err)
	}
	// generate the html
	_, err = builder.ToHTML("masterid", false)
	if err != nil {
		t.Fatal(err)
	}
	testGenerateHomeEmptyPanelIDsMap(t, builder)
	testGenerateTabBarIDStartPanelIDMap(t, builder)
	testGenerateHomeButtonPanelGroups(t, builder)
	testgenerateHomePanelNameTemplateMap(t, builder)
	testGenerateHomeTemplatePanelName(t, builder)
	// deep
	var deepHomes []*Button
	if deepHomes, err = buildDeepHomes(); err != nil {
		t.Fatal(err)
	}
	builder2 := NewBuilder()
	if err = builder2.BuildFromHomes(deepHomes); err != nil {
		t.Fatal(err)
	}
	if _, err = builder2.ToHTML("masterid", false); err != nil {
		t.Fatal(err)
	}
	testBuilderGenerateHomeEmptyInsidePanelNamePathMap(t, builder2)
	testBuilderGenerateHomePanelNamePanelMap(t, builder2)

	var longOkHomes []*Button
	if longOkHomes, err = buildLongOkHomes(); err != nil {
		t.Fatal(err)
	}
	builder3 := NewBuilder()
	if err = builder3.BuildFromHomes(longOkHomes); err != nil {
		t.Fatal(err)
	}
	if _, err = builder3.ToHTML("masterid", false); err != nil {
		t.Fatal(err)
	}
	testGenerateButtonIDsPanelIDs(t, builder3)
	testGenerateTabIDsPanelIDs(t, builder3)
}

func testGenerateHomeButtonPanelGroups(t *testing.T, builder *Builder) {
	wantStr := `OneButton:
- istabbutton: false
  buttonname: OneButton
  buttonid: masterid-home-pad-OneButton
  panelnamesidmap:
  OnePanel:
    id: OnePanel
    name: OnePanel
    HasRealTabs: false
    note: p1 note
    markup: <p>Panel 1-1</p>
    HVScroll: false
    HTMLID: masterid-home-pad-OneButton-OnePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
    H3ID: ""
  TwoPanel:
    id: TwoPanel
    name: TwoPanel
    HasRealTabs: false
    note: p2 note
    markup: <p>Panel 2-1</p>
    HVScroll: false
    HTMLID: masterid-home-pad-OneButton-TwoPanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
    H3ID: ""
TwoButton:
- istabbutton: false
  buttonname: TwoButton
  buttonid: masterid-home-pad-TwoButton
  panelnamesidmap:
  FourPanel:
    id: FourPanel
    name: FourPanel
    HasRealTabs: false
    note: p4 note
    markup: <p>Panel 4-1</p>
    HVScroll: false
    HTMLID: masterid-home-pad-TwoButton-FourPanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
    H3ID: ""
  ThreePanel:
    id: ThreePanel
    name: ThreePanel
    HasRealTabs: false
    note: p3 note
    markup: <p>Panel 3-1</p>
    HVScroll: false
    HTMLID: masterid-home-pad-TwoButton-ThreePanel
    TabBarHTMLID: ""
    UnderTabBarHTMLID: ""
    H3ID: ""
`
	var wants map[string][]*ButtonPanelGroup
	if err := yaml.Unmarshal([]byte(wantStr), &wants); err != nil {
		t.Fatal(err)
	}
	results := builder.GenerateHomeButtonPanelGroups()
	//bb, _ := yaml.Marshal(results)
	//t.Fatal(string(bb))
	for home, wbpg := range wants {
		rbpg, ok := results[home]
		if !ok {
			t.Errorf(`GenerateHomeButtonPanelGroups home %s is missing`, home)
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

func testGenerateTabBarIDStartPanelIDMap(t *testing.T, builder *Builder) {
	wants := map[string]string{}
	results := builder.GenerateTabBarIDStartPanelIDMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf(`builder.GenerateTabBarIDStartPanelIDMap() generated %#v\n\nwant: %#v`, results, wants)
	}
}

func testGenerateHomeEmptyPanelIDsMap(t *testing.T, builder *Builder) {
	wants := map[string][]string{"OneButton": []string{"masterid-home-pad-OneButton-OnePanel", "masterid-home-pad-OneButton-TwoPanel"}, "TwoButton": []string{"masterid-home-pad-TwoButton-ThreePanel", "masterid-home-pad-TwoButton-FourPanel"}}

	results := builder.GenerateHomeEmptyPanelIDsMap()
	if ok := reflect.DeepEqual(results, wants); !ok {
		t.Fatalf("builder.GenerateHomeEmptyPanelIDsMap() generated %#v\n\nwant: %#v", results, wants)
	}
}

func testgenerateHomePanelNameTemplateMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "a",
			want: map[string]map[string]string{"OneButton": map[string]string{"OnePanel": "\n<!--\n\nPanel name: \"OnePanel\"\n\nPanel note: p1 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-TwoPanel-inner-user-content\n  * Name: TwoPanel\n  * Note: p2 note\n\n-->\n\n<p>Panel 1-1</p>\n", "TwoPanel": "\n<!--\n\nPanel name: \"TwoPanel\"\n\nPanel note: p2 note\n\nThis is one of a group of 2 panels displayed when the \"One\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-OneButton-OnePanel-inner-user-content\n  * Name: OnePanel\n  * Note: p1 note\n\n-->\n\n<p>Panel 2-1</p>\n"}, "TwoButton": map[string]string{"FourPanel": "\n<!--\n\nPanel name: \"FourPanel\"\n\nPanel note: p4 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-ThreePanel-inner-user-content\n  * Name: ThreePanel\n  * Note: p3 note\n\n-->\n\n<p>Panel 4-1</p>\n", "ThreePanel": "\n<!--\n\nPanel name: \"ThreePanel\"\n\nPanel note: p3 note\n\nThis is one of a group of 2 panels displayed when the \"Two\" button is clicked.\n\nThis panel is just 1 in a group of 2 panels.\nYour other panel in this group is\n\n  * The content panel <div #masterid-home-pad-TwoButton-FourPanel-inner-user-content\n  * Name: FourPanel\n  * Note: p4 note\n\n-->\n\n<p>Panel 3-1</p>\n"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateHomePanelNameTemplateMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateHomePanelNameTemplateMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testGenerateHomeTemplatePanelName(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "wtf",
			want: map[string][]string{"OneButton": []string{"OnePanel", "TwoPanel"}, "TwoButton": []string{"ThreePanel", "FourPanel"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateHomeTemplatePanelName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateHomeTemplatePanelName() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilderGenerateHomeEmptyInsidePanelNamePathMap(t *testing.T, builder *Builder) {
	tests := []struct {
		name string
		want map[string]map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: map[string]map[string][]string{"OneButton": map[string][]string{"OneOnePanel": []string{"OneButton", "OnePanel", "OneOneButton"}, "TwoOnePanel": []string{"OneButton", "TwoPanel", "TwoOneButton"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.GenerateHomeEmptyInsidePanelNamePathMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.GenerateHomeEmptyInsidePanelNamePathMap() = %#v, want %v", got, tt.want)
			}
		})
	}
}

func testBuilderGenerateHomePanelNamePanelMap(t *testing.T, builder *Builder) {
	wantStr := `
OneButton:
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
	got := builder.GenerateHomePanelNamePanelMap()
	// t.Fatalf("%#v", got)
	for wHome, wPanelNamePanelMap := range want {
		gPanelNamePanelMap, ok := got[wHome]
		if !ok {
			t.Errorf("home %q not found in result.", wHome)
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
	wantTabs := map[string]TabSpawnPanelIDs{
		"OneTab":  TabSpawnPanelIDs{Spawn: false, PanelIDs: []string{"OnlyPanel"}},
		"TwoTab":  TabSpawnPanelIDs{Spawn: true, PanelIDs: []string{"LastTabPanel"}},
		"ZeroTab": TabSpawnPanelIDs{Spawn: true, PanelIDs: []string{"ZeroTabFirstPanel", "ZeroTabSecondPanel"}},
	}

	t.Run("testGenerateTabIDsPanelIDs", func(t *testing.T) {
		if gotTabs := builder.GenerateTabIDsPanelIDs(); !reflect.DeepEqual(gotTabs, wantTabs) {
			fmt.Printf("Builder.GenerateTabIDsPanelIDs() = %#v, want %#v", gotTabs, wantTabs)
			t.Fatal("oej")
			// t.Errorf("Builder.GenerateTabIDsPanelIDs() = %#v, want %#v", gotTabs, wantTabs)
		}
	})
}
