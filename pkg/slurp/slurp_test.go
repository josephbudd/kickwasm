package slurp

import (
	"reflect"
	"testing"

	"github.com/josephbudd/kickwasm/pkg/project"
	yaml "gopkg.in/yaml.v2"
)

func Test_constructButton(t *testing.T) {
	type args struct {
		panel *project.Panel
		b     *ButtonInfo
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantButtons []*project.Button
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				panel: &project.Panel{
					Buttons: make([]*project.Button, 0, 5),
				},
				b: &ButtonInfo{
					ID:      "id",
					Label:   "Label",
					Heading: "Heading",
					CC:      "CC",
				},
			},
			wantErr: false,
			wantButtons: []*project.Button{
				{
					ID:               "id",
					Label:            "Label",
					Heading:          "Heading",
					Location:         "CC",
					Panels:           []*project.Panel{},
					HTMLID:           "",
					PanelHTMLID:      "",
					PanelInnerHTMLID: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := constructButton(tt.args.panel, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("constructButton() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if ok := reflect.DeepEqual(tt.args.panel.Buttons, tt.wantButtons); !ok {
			t.Errorf(`reflect.DeepEqual(tt.args.panel.Buttons, tt.wantButtons) its: %#v`, tt.args.panel.Buttons)
		}
	}
}

func Test_constructTab(t *testing.T) {
	type args struct {
		panel *project.Panel
		t     *TabInfo
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantTabs []*project.Tab
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				panel: &project.Panel{
					Buttons: make([]*project.Button, 0, 5),
				},
				t: &TabInfo{
					ID:    "id",
					Label: "Label",
				},
			},
			wantErr: false,
			wantTabs: []*project.Tab{
				{
					ID:     "id",
					Label:  "Label",
					Panels: []*project.Panel{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := constructTab(tt.args.panel, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("constructTab() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if ok := reflect.DeepEqual(tt.args.panel.Tabs, tt.wantTabs); !ok {
			t.Errorf(`reflect.DeepEqual(tt.args.panel.Tabs, tt.wantTabs) its: %#v`, tt.args.panel.Buttons)
		}
	}
}

func Test_constructTabPanel(t *testing.T) {
	type args struct {
		tab *project.Tab
		pi  *PanelInfo
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantPanels []*project.Panel
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			args: args{
				tab: &project.Tab{
					ID:     "id",
					Label:  "Label",
					Panels: []*project.Panel{},
				},
				pi: &PanelInfo{
					ID:     "id",
					Name:   "name",
					Note:   "note",
					Markup: "markup",
				},
			},
			wantErr: false,
			wantPanels: []*project.Panel{
				{
					ID:      "id",
					Name:    "name",
					Note:    "note",
					Markup:  "markup",
					Tabs:    make([]*project.Tab, 0, 5),
					Buttons: make([]*project.Button, 0, 5),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := constructTabPanel(tt.args.tab, tt.args.pi); (err != nil) != tt.wantErr {
				t.Errorf("constructTabPanel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if ok := reflect.DeepEqual(tt.args.tab.Panels, tt.wantPanels); !ok {
			t.Errorf(`reflect.DeepEqual(tt.args.tab.Panels, tt.wantPanels) its: %#v`, tt.args.tab.Panels[0])
		}
	}
}

func Test_constructButtonPanel(t *testing.T) {
	type args struct {
		button *project.Button
		pi     *PanelInfo
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantPanels []*project.Panel
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			args: args{
				button: &project.Button{
					ID:     "id",
					Label:  "label",
					Panels: []*project.Panel{},
				},
				pi: &PanelInfo{
					ID:     "id",
					Name:   "name",
					Note:   "note",
					Markup: "markup",
				},
			},
			wantErr: false,
			wantPanels: []*project.Panel{
				{
					ID:      "id",
					Name:    "name",
					Note:    "note",
					Markup:  "markup",
					Tabs:    make([]*project.Tab, 0, 5),
					Buttons: make([]*project.Button, 0, 5),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := constructButtonPanel(tt.args.button, tt.args.pi); (err != nil) != tt.wantErr {
				t.Errorf("constructButtonPanel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if ok := reflect.DeepEqual(tt.args.button.Panels, tt.wantPanels); !ok {
			t.Errorf(`reflect.DeepEqual(tt.args.button.Panels, tt.wantPanels) its: %#v`, tt.args.button.Panels[0])
			t.Errorf(`%#v`, *tt.wantPanels[0])
		}
	}
}

func TestDo(t *testing.T) {
	type args struct {
		yamlPath string
	}
	tests := []struct {
		name           string
		args           args
		want           *project.Builder
		wantErr        bool
		wantErrMessage string
		wantAppInfo    *ApplicationInfo
	}{
		// TODO: Add test cases.
		{
			name: "fail empty tab",
			args: args{
				yamlPath: "testyaml/fails/empty_tab.yaml",
			},
			wantErr:        true,
			wantErrMessage: `a tab labeled "One" is missing panel files in testyaml/fails/empty_tab.yaml`,
		},
		{
			name: "fail duplicate home button name",
			args: args{
				yamlPath: "testyaml/fails/dup_service_button_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the home button name "NextButton" is used more than once`,
		},
		{
			name: "fail bad panel name",
			args: args{
				yamlPath: "testyaml/fails/bad_panel_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button panel name "PName" should end with the suffix "Panel" in testyaml/fails/bad_panel_name.yaml`,
		},
		{
			name: "fail bad button name",
			args: args{
				yamlPath: "testyaml/fails/bad_button_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the XXX home panel button name "Next" should end with the suffix "Button"`,
		},
		{
			name: "fail bad tab name",
			args: args{
				yamlPath: "testyaml/fails/bad_tab_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the XXX home panel: the panel named "PNamePanel" tab name "One" should end with the suffix "Tab"`,
		},

		{
			name: "fail duplicate panel name",
			args: args{
				yamlPath: "testyaml/fails/dup_panel_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button panel name "PNamePanel" used in testyaml/fails/dup_panel_name.yaml has already been used in testyaml/fails/dup_panel_name.yaml`,
		},
		{
			name: "fail duplicate button name",
			args: args{
				yamlPath: "testyaml/fails/dup_button_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button name "OneButton" used in testyaml/fails/dup_button_name.yaml has already been used in testyaml/fails/dup_button_name.yaml`,
		},
		{
			name: "fail duplicate tab name",
			args: args{
				yamlPath: "testyaml/fails/dup_tab_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the tab name "OneTab" used in testyaml/fails/dup_tab_name.yaml has already been used in testyaml/fails/dup_tab_name.yaml`,
		},
		{
			name: "fail missing panel name",
			args: args{
				yamlPath: "testyaml/fails/missing_panel_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button panel name "" should end with the suffix "Panel" in testyaml/fails/missing_panel_name.yaml`,
		},
		{
			name: "fail missing button name",
			args: args{
				yamlPath: "testyaml/fails/missing_button_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `in the home named "XXX", a button is missing a name in testyaml/fails/missing_button_name.yaml`,
		},
		{
			name: "panels",
			args: args{
				yamlPath: "testyaml/panels_test/panels.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button name "NextColorsButton" used in testyaml/panels_test/panels.yaml has already been used in testyaml/panels_test/panels.yaml`,
		},
		{
			name: "colors",
			args: args{
				yamlPath: "testyaml/deep_test/colors.yaml",
			},
			wantErr:        true, // levels are no longer limited.
			wantErrMessage: `the button name "ClickHereButton" used in testyaml/deep_test/colors.yaml has already been used in testyaml/deep_test/colors.yaml`,
		},
		{
			name: "pass",
			args: args{
				yamlPath: "testyaml/simple_test/app.yaml",
			},
			want:    nil,
			wantErr: false,
			wantAppInfo: &ApplicationInfo{
				Title:      "Test 1",
				ImportPath: "github.com/josephbudd/kickwasm/examples/test1",
				Homes: []*ButtonInfo{
					{
						ID:         "1Button",
						Label:      "Button 1",
						Heading:    "Button 1",
						CC:         "",
						PanelFiles: []string{"home1/panel1.yaml", "home1/panel2.yaml"},
						Panels:     []*PanelInfo{},
					},
					{
						ID:         "2Button",
						Label:      "Button 2",
						Heading:    "Button 2 Heading.",
						CC:         "",
						PanelFiles: []string{"home2/panel3.yaml", "home2/panel5.yaml"},
						Panels:     []*PanelInfo{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := NewSlurper()
			got, err := sl.Gulp(tt.args.yamlPath)
			if tt.wantErr && err != nil {
				if err.Error() != tt.wantErrMessage {
					t.Errorf("%s: got the wrong error message got: %s", tt.name, err.Error())
				}
			} else {
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: sl.Gulp(...) error = %v, wantErr %v", tt.name, err, tt.wantErr)
					return
				}
				if tt.wantAppInfo != nil {
					if len(got.Homes) != len(tt.wantAppInfo.Homes) {
						t.Error(`sl.Gulp(...) error len(got.Homes) != len(tt.wantAppInfo.Homes)`)
						t.Errorf("sl.Gulp(...) = %v, want %v", got, tt.want)
					}
					if got.Homes[0].ID != tt.wantAppInfo.Homes[0].ID {
						t.Error(`sl.Gulp(...) error got.Homes[0].ID != tt.wantAppInfo.Homes[0].ID`)
					}
					if got.Homes[1].ID != tt.wantAppInfo.Homes[1].ID {
						t.Error(`sl.Gulp(...) error got.Homes[1].ID != tt.wantAppInfo.Homes[1].ID`)
					}
					if got.Homes[0].Label != tt.wantAppInfo.Homes[0].Label {
						t.Error(`sl.Gulp(...) error got.Homes[0].Label != tt.wantAppInfo.Homes[0].Label`)
					}
					if got.Homes[1].Label != tt.wantAppInfo.Homes[1].Label {
						t.Error(`sl.Gulp(...) error got.Homes[1].Label != tt.wantAppInfo.Homes[1].Label`)
					}
				}
			}
		})
	}
}

func Test_slurpApplication(t *testing.T) {
	type args struct {
		fpath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				fpath: "testyaml/simple_test/app.yaml",
			},
			want: `
sourcePath: testyaml/simple_test/app.yaml
title: Test 1
importPath: github.com/josephbudd/kickwasm/examples/test1
homes:
- sourcePath: testyaml/simple_test/app.yaml
  name: Home1
  button:
    sourcePath: testyaml/simple_test/app.yaml
    name: OneButton
    label: Button 1
    heading: Button 1 Heading.
    cc: Button 1
    panelFiles:
    - home1/panel1.yaml
    - home1/panel2.yaml
    panels:
    - sourcePath: testyaml/simple_test/home1/panel1.yaml
      level: 1
      id: ""
      name: OnePanel
      buttons: []
      tabs: []
      markup: <p>Panel 1-1</p>
      note: p1 note
      HVScroll: false
    - sourcePath: testyaml/simple_test/home1/panel2.yaml
      level: 1
      id: ""
      name: TwoPanel
      buttons: []
      tabs: []
      markup: <p>Panel 2-1</p>
      note: p2 note
      HVScroll: false
- sourcePath: testyaml/simple_test/app.yaml
  name: Home2
  button:
    sourcePath: testyaml/simple_test/app.yaml
    name: TwoButton
    label: Button 2
    heading: Button 2 Heading.
    cc: Button 2
    panelFiles:
    - home2/panel3.yaml
    - home2/panel5.yaml
    panels:
    - sourcePath: testyaml/simple_test/home2/panel3.yaml
      level: 1
      id: ""
      name: ThreePanel
      buttons:
      - sourcePath: testyaml/simple_test/home2/panel3.yaml
        name: Panel31Button
        label: Panel 3 Button 1
        heading: Panel 3 Button 1 Heading.
        cc: Panel 3 Button 1
        panelFiles:
        - panel3-button1-panels/panel1.yaml
        - panel3-button1-panels/panel2.yaml
        panels:
        - sourcePath: testyaml/simple_test/home2/panel3-button1-panels/panel1.yaml
          level: 2
          id: ""
          name: Panel3Button1OnePanel
          buttons: []
          tabs: []
          markup: Panel3Button1Panel1 Markup
          note: Panel3Button1Panel1 Note
          HVScroll: false
        - sourcePath: testyaml/simple_test/home2/panel3-button1-panels/panel2.yaml
          level: 2
          id: ""
          name: Panel3Button1TwoPanel
          buttons: []
          tabs: []
          markup: Panel3Button1Panel2 Markup
          note: Panel3Button1Panel2 Note
          HVScroll: false
      - sourcePath: testyaml/simple_test/home2/panel3.yaml
        name: Panel32Button
        label: Panel 3 Button 2
        heading: Panel 3 Button 2 Heading.
        cc: Panel 3 Button 2
        panelFiles:
        - panel3-button2-panels/panel1.yaml
        - panel3-button2-panels/panel2.yaml
        panels:
        - sourcePath: testyaml/simple_test/home2/panel3-button2-panels/panel1.yaml
          level: 2
          id: ""
          name: Panel3Button2OnePanel
          buttons: []
          tabs: []
          markup: Panel3Button2Panel1 Markup
          note: Panel3Button2Panel1 Note
          HVScroll: false
        - sourcePath: testyaml/simple_test/home2/panel3-button2-panels/panel2.yaml
          level: 2
          id: ""
          name: Panel3Button2TwoPanel
          buttons: []
          tabs: []
          markup: Panel3Button2Panel2 Markup
          note: Panel3Button2Panel2 Note
          HVScroll: false
      tabs: []
      note: p3 note
      HVScroll: false
    - sourcePath: testyaml/simple_test/home2/panel5.yaml
      level: 1
      id: ""
      name: FivePanel
      buttons: []
      tabs:
      - sourcePath: testyaml/simple_test/home2/panel5.yaml
        name: P51Tab
        label: P5T1 Label
        heading: P5T1 Label
        panelFiles:
        - panel5-tab1-panels/panel1.yaml
        - panel5-tab1-panels/panel2.yaml
        panels:
        - sourcePath: testyaml/simple_test/home2/panel5-tab1-panels/panel1.yaml
          level: 2
          id: ""
          name: Panel5Tab1OnePanel
          buttons: []
          tabs: []
          markup: Panel5Tab1Panel1 Markup
          note: Panel5Tab1Panel1 Note
          HVScroll: false
        - sourcePath: testyaml/simple_test/home2/panel5-tab1-panels/panel2.yaml
          level: 2
          id: ""
          name: Panel5Tab1TwoPanel
          buttons: []
          tabs: []
          markup: Panel5Tab1Panel2 Markup
          note: Panel5Tab1Panel2 Note
          HVScroll: false
      - sourcePath: testyaml/simple_test/home2/panel5.yaml
        name: P52Tab
        label: P5T2 Label
        heading: P5T2 Label
        panelFiles:
        - panel5-tab2-panels/panel1.yaml
        - panel5-tab2-panels/panel2.yaml
        panels:
        - sourcePath: testyaml/simple_test/home2/panel5-tab2-panels/panel1.yaml
          level: 2
          id: ""
          name: Panel5Tab2OnePanel
          buttons: []
          tabs: []
          markup: Panel5Tab2Panel1 Markup
          note: Panel5Tab2Panel1 Note
          HVScroll: false
        - sourcePath: testyaml/simple_test/home2/panel5-tab2-panels/panel2.yaml
          level: 2
          id: ""
          name: Panel5Tab2TwoPanel
          buttons: []
          tabs: []
          markup: Panel5Tab2Panel2 Markup
          note: Panel5Tab2Panel2 Note
          HVScroll: false
      note: p5 note
      HVScroll: false
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := NewSlurper()
			got, err := sl.slurpApplication(tt.args.fpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("sl.slurpApplication() got error %#v, want error %v", err, tt.wantErr)
				return
			}
			want := &ApplicationInfo{}
			if err = yaml.Unmarshal([]byte(tt.want), want); err != nil {
				t.Fatal("yaml.Unmarshal([]byte(tt.want), want) err is " + err.Error())
			}
			if !checkApplicaitonInfo(got, want) {
				gotbb, err := yaml.Marshal(got)
				if err != nil {
					t.Fatal(err)
				}
				t.Errorf("sl.slurpApplication(): !reflect.DeepEqual(want, got): got %s\n\nwant:\n%s", gotbb, tt.want)
			}
		})
	}
}

func checkApplicaitonInfo(control, test *ApplicationInfo) (equal bool) {
	if control.SourcePath != test.SourcePath {
		return
	}
	if control.Title != test.Title {
		return
	}
	if control.ImportPath != test.ImportPath {
		return
	}
	if len(control.Homes) != len(test.Homes) {
		return
	}
	for i, cHomeButton := range control.Homes {
		tHomeButton := test.Homes[i]
		if cHomeButton.SourcePath != tHomeButton.SourcePath {
			return
		}
		if cHomeButton.Label != tHomeButton.Label {
			return
		}
		if cHomeButton.ID != tHomeButton.ID {
			return
		}
		if cHomeButton.Heading != tHomeButton.Heading {
			return
		}
		if cHomeButton.CC != tHomeButton.CC {
			return
		}
		if len(cHomeButton.PanelFiles) != len(tHomeButton.PanelFiles) {
			return
		}
		for i, cFile := range cHomeButton.PanelFiles {
			if cFile != tHomeButton.PanelFiles[i] {
				return
			}
		}
		if len(cHomeButton.Panels) != len(tHomeButton.Panels) {
			return
		}
		for i, cPanel := range cHomeButton.Panels {
			if equal = checkPanelInfo(cPanel, tHomeButton.Panels[i]); !equal {
				return
			}
		}
	}
	equal = true
	return
}

func checkPanelInfo(control, test *PanelInfo) (equal bool) {
	if control.SourcePath != test.SourcePath {
		return
	}
	if control.Level != test.Level {
		return
	}
	if control.ID != test.ID {
		return
	}
	if control.Name != test.Name {
		return
	}
	if control.Markup != test.Markup {
		return
	}
	if control.Note != test.Note {
		return
	}
	if control.HVScroll != test.HVScroll {
		return
	}
	if control.Note != test.Note {
		return
	}
	if len(control.Buttons) != len(test.Buttons) {
		return
	}
	for i, cButton := range control.Buttons {
		tButton := test.Buttons[i]
		if equal = checkButtonInfo(cButton, tButton); !equal {
			return
		}
	}
	if len(control.Tabs) != len(test.Tabs) {
		return
	}
	for i, cTab := range control.Tabs {
		tTab := test.Tabs[i]
		if equal = checkTabInfo(cTab, tTab); !equal {
			return
		}
	}
	equal = true
	return
}

func checkButtonInfo(control, test *ButtonInfo) (equal bool) {
	if control.SourcePath != test.SourcePath {
		return
	}
	if control.Label != test.Label {
		return
	}
	if control.ID != test.ID {
		return
	}
	if control.Heading != test.Heading {
		return
	}
	if control.CC != test.CC {
		return
	}
	if len(control.PanelFiles) != len(test.PanelFiles) {
		return
	}
	for i, cFile := range control.PanelFiles {
		if cFile != test.PanelFiles[i] {
			return
		}
	}
	if len(control.Panels) != len(test.Panels) {
		return
	}
	for i, cPanel := range control.Panels {
		if equal = checkPanelInfo(cPanel, test.Panels[i]); !equal {
			return
		}
	}
	equal = true
	return
}

func checkTabInfo(control, test *TabInfo) (equal bool) {
	if control.SourcePath != test.SourcePath {
		return
	}
	if control.Label != test.Label {
		return
	}
	if control.ID != test.ID {
		return
	}
	if control.Heading != test.Heading {
		return
	}
	if control.Spawn != test.Spawn {
		return
	}
	if len(control.PanelFiles) != len(test.PanelFiles) {
		return
	}
	for i, cFile := range control.PanelFiles {
		if cFile != test.PanelFiles[i] {
			return
		}
	}
	if len(control.Panels) != len(test.Panels) {
		return
	}
	for i, cPanel := range control.Panels {
		if equal = checkPanelInfo(cPanel, test.Panels[i]); !equal {
			return
		}
	}
	equal = true
	return
}
