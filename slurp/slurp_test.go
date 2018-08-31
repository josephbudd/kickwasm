package slurp

import (
	"reflect"
	"testing"

	"github.com/josephbudd/kickwasm/tap"
	yaml "gopkg.in/yaml.v2"
)

func Test_constructButton(t *testing.T) {
	type args struct {
		panel *tap.Panel
		b     *ButtonInfo
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantButtons []*tap.Button
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				panel: &tap.Panel{
					Buttons: make([]*tap.Button, 0, 5),
				},
				b: &ButtonInfo{
					ID:      "id",
					Label:   "Label",
					Heading: "Heading",
					CC:      "CC",
				},
			},
			wantErr: false,
			wantButtons: []*tap.Button{
				&tap.Button{
					ID:               "id",
					Label:            "Label",
					Heading:          "Heading",
					Location:         "CC",
					Panels:           []*tap.Panel{},
					Generated:        false,
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
		panel *tap.Panel
		t     *TabInfo
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantTabs []*tap.Tab
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				panel: &tap.Panel{
					Buttons: make([]*tap.Button, 0, 5),
				},
				t: &TabInfo{
					ID:    "id",
					Label: "Label",
				},
			},
			wantErr: false,
			wantTabs: []*tap.Tab{
				&tap.Tab{
					ID:     "id",
					Label:  "Label",
					Panels: []*tap.Panel{},
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
		tab *tap.Tab
		pi  *PanelInfo
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantPanels []*tap.Panel
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			args: args{
				tab: &tap.Tab{
					ID:     "id",
					Label:  "Label",
					Panels: []*tap.Panel{},
				},
				pi: &PanelInfo{
					ID:     "id",
					Name:   "name",
					Note:   "note",
					Markup: "markup",
				},
			},
			wantErr: false,
			wantPanels: []*tap.Panel{
				{
					ID:      "id",
					Name:    "name",
					Note:    "note",
					Markup:  "markup",
					Tabs:    make([]*tap.Tab, 0, 5),
					Buttons: make([]*tap.Button, 0, 5),
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
		button *tap.Button
		pi     *PanelInfo
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantPanels []*tap.Panel
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			args: args{
				button: &tap.Button{
					ID:     "id",
					Label:  "label",
					Panels: []*tap.Panel{},
				},
				pi: &PanelInfo{
					ID:     "id",
					Name:   "name",
					Note:   "note",
					Markup: "markup",
				},
			},
			wantErr: false,
			wantPanels: []*tap.Panel{
				{
					ID:      "id",
					Name:    "name",
					Note:    "note",
					Markup:  "markup",
					Tabs:    make([]*tap.Tab, 0, 5),
					Buttons: make([]*tap.Button, 0, 5),
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
		want           *tap.Builder
		wantErr        bool
		wantErrMessage string
		wantAppInfo    *ApplicationInfo
	}{
		// TODO: Add test cases.
		{
			name: "fail bad panel name",
			args: args{
				yamlPath: "testyaml/fails/bad_panel_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the button panel name "PNamePanel" used in testyaml/fails/dup_panel_name.yaml has already been used in testyaml/fails/dup_panel_name.yaml`,
		},
		{
			name: "fail bad button name",
			args: args{
				yamlPath: "testyaml/fails/bad_button_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the XXX service panel: the button panel named "PNamePanel" has more than one button named "OneButton"`,
		},
		{
			name: "fail bad tab name",
			args: args{
				yamlPath: "testyaml/fails/bad_tab_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the XXX service panel: the tab panel named "PNamePanel" has more than one tab named "OneTab"`,
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
			wantErrMessage: `the XXX service panel: the button panel named "PNamePanel" has more than one button named "OneButton"`,
		},
		{
			name: "fail duplicate tab name",
			args: args{
				yamlPath: "testyaml/fails/dup_tab_name.yaml",
			},
			wantErr:        true,
			wantErrMessage: `the XXX service panel: the tab panel named "PNamePanel" has more than one tab named "OneTab"`,
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
			wantErrMessage: "a button is missing a name in testyaml/fails/missing_button_name.yaml",
		},
		{
			name: "panels",
			args: args{
				yamlPath: "testyaml/panels_test/panels.yaml",
			},
		},
		{
			name: "colors",
			args: args{
				yamlPath: "testyaml/deep_test/colors.yaml",
			},
			wantErr:        false, // levels are no longer limited.
			wantErrMessage: `the panel named "LevelSixPanel" is too deep to have buttons in testyaml/deep_test/level2/level3/level4/level5/level6/panel.yaml`,
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
				Repos:      []string{"Test1"},
				Services: []*ServiceInfo{
					&ServiceInfo{
						Name: "Service1",
						Button: &ButtonInfo{
							ID:         "1Button",
							Label:      "Button 1",
							Heading:    "Button 1 Heading.",
							CC:         "",
							PanelFiles: []string{"service1/panel1.yaml", "service1/panel2.yaml"},
							Panels:     []*PanelInfo{},
						},
					},
					&ServiceInfo{
						Name: "Service2",
						Button: &ButtonInfo{
							ID:         "2Button",
							Label:      "Button 2",
							Heading:    "Button 2 Heading.",
							CC:         "",
							PanelFiles: []string{"service2/panel3.yaml", "service2/panel5.yaml"},
							Panels:     []*PanelInfo{},
						},
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
					if len(got.Services) != len(tt.wantAppInfo.Services) {
						t.Error(`sl.Gulp(...) error len(got.Services) != len(tt.wantAppInfo.Services)`)
						t.Errorf("sl.Gulp(...) = %v, want %v", got, tt.want)
					}
					if got.Services[0].Name != tt.wantAppInfo.Services[0].Name {
						t.Error(`sl.Gulp(...) error got.Services[0].Name != tt.wantAppInfo.Services[0].Name`)
					}
					if got.Services[1].Name != tt.wantAppInfo.Services[1].Name {
						t.Error(`sl.Gulp(...) error got.Services[1].Name != tt.wantAppInfo.Services[1].Name`)
					}
					if got.Services[0].Button.Label != tt.wantAppInfo.Services[0].Button.Label {
						t.Error(`sl.Gulp(...) error got.Services[0].Button.Label != tt.wantAppInfo.Services[0].Button.Label`)
					}
					if got.Services[1].Button.Label != tt.wantAppInfo.Services[1].Button.Label {
						t.Error(`sl.Gulp(...) error got.Services[1].Button.Label != tt.wantAppInfo.Services[1].Button.Label`)
					}
				}
				/*if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("sl.Gulp(...) = %v, want %v", got, tt.want)
				}
				bb, err := yaml.Marshal(got)
				if err != nil {
					t.Fatal(err)
				}
				t.Error(string(bb))
				*/
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
		want    *ApplicationInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				fpath: "testyaml/simple_test/app.yaml",
			},
			want: &ApplicationInfo{
				SourcePath: "testyaml/simple_test/app.yaml",
				Title:      "Test 1",
				ImportPath: "github.com/josephbudd/kickwasm/examples/test1",
				Repos:      []string{"Test1"},
				Services: []*ServiceInfo{
					&ServiceInfo{
						SourcePath: "testyaml/simple_test/app.yaml",
						Name:       "Service1",
						Button: &ButtonInfo{
							SourcePath: "testyaml/simple_test/app.yaml",
							ID:         "",
							Label:      "Button 1",
							Heading:    "Button 1 Heading.",
							CC:         "Button 1 Heading.",
							PanelFiles: []string{
								"service1/panel1.yaml",
								"service1/panel2.yaml",
							},
							Panels: []*PanelInfo{
								&PanelInfo{
									SourcePath: "testyaml/simple_test/service1/panel1.yaml",
									Level:      2,
									ID:         "",
									Name:       "OnePanel",
									Note:       "p1 note",
									Buttons:    []*ButtonInfo(nil),
									Tabs:       []*TabInfo(nil),
									Markup:     "<p>Panel 1-1</p>",
								},
								&PanelInfo{
									SourcePath: "testyaml/simple_test/service1/panel2.yaml",
									Level:      2,
									ID:         "",
									Name:       "TwoPanel",
									Note:       "p2 note",
									Buttons:    []*ButtonInfo(nil),
									Tabs:       []*TabInfo(nil),
									Markup:     "<p>Panel 2-1</p>",
								},
							},
						},
					},
					&ServiceInfo{
						SourcePath: "testyaml/simple_test/app.yaml",
						Name:       "Service2",
						Button: &ButtonInfo{
							SourcePath: "testyaml/simple_test/app.yaml",
							ID:         "",
							Label:      "Button 2",
							Heading:    "Button 2 Heading.",
							CC:         "Button 2 Heading.",
							PanelFiles: []string{
								"service2/panel3.yaml",
								"service2/panel5.yaml",
							},
							Panels: []*PanelInfo{
								&PanelInfo{
									SourcePath: "testyaml/simple_test/service2/panel3.yaml",
									Level:      2,
									ID:         "",
									Name:       "ThreePanel",
									Note:       "p3 note",
									Buttons: []*ButtonInfo{
										&ButtonInfo{
											SourcePath: "testyaml/simple_test/service2/panel3.yaml",
											ID:         "",
											Label:      "Panel 3 Button 1",
											Heading:    "Panel 3 Button 1 Heading.",
											CC:         "Panel 3 Button 1 Heading.",
											PanelFiles: []string{
												"panel3-button1-panels/panel1.yaml",
												"panel3-button1-panels/panel2.yaml",
											},
											Panels: []*PanelInfo{
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel3-button1-panels/panel1.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel3Button1OnePanel",
													Note:       "Panel3Button1Panel1 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel3Button1Panel1 Markup",
												},
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel3-button1-panels/panel2.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel3Button1TwoPanel",
													Note:       "Panel3Button1Panel2 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel3Button1Panel2 Markup",
												},
											},
										},
										&ButtonInfo{
											SourcePath: "testyaml/simple_test/service2/panel3.yaml",
											ID:         "",
											Label:      "Panel 3 Button 2",
											Heading:    "Panel 3 Button 2 Heading.",
											CC:         "Panel 3 Button 2 Heading.",
											PanelFiles: []string{
												"panel3-button2-panels/panel1.yaml",
												"panel3-button2-panels/panel2.yaml",
											},
											Panels: []*PanelInfo{
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel3-button2-panels/panel1.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel3Button2OnePanel",
													Note:       "Panel3Button2Panel1 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel3Button2Panel1 Markup",
												},
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel3-button2-panels/panel2.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel3Button2TwoPanel",
													Note:       "Panel3Button2Panel2 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel3Button2Panel2 Markup",
												},
											},
										},
									},
									Tabs:   []*TabInfo(nil),
									Markup: "",
								},
								&PanelInfo{
									SourcePath: "testyaml/simple_test/service2/panel5.yaml",
									Level:      2,
									ID:         "",
									Name:       "FivePanel",
									Note:       "p5 note",
									Buttons:    []*ButtonInfo{},
									Tabs: []*TabInfo{
										&TabInfo{
											SourcePath: "testyaml/simple_test/service2/panel5.yaml",
											ID:         "",
											Label:      "P5T1 Label",
											PanelFiles: []string{
												"panel5-tab1-panels/panel1.yaml",
												"panel5-tab1-panels/panel2.yaml",
											},
											Panels: []*PanelInfo{
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel5-tab1-panels/panel1.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel5Tab1OnePanel",
													Note:       "Panel5Tab1Panel1 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel5Tab1Panel1 Markup",
												},
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel5-tab1-panels/panel2.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel5Tab1TwoPanel",
													Note:       "Panel5Tab1Panel2 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel5Tab1Panel2 Markup",
												},
											},
										},
										&TabInfo{
											SourcePath: "testyaml/simple_test/service2/panel5.yaml",
											ID:         "",
											Label:      "P5T2 Label",
											PanelFiles: []string{
												"panel5-tab2-panels/panel1.yaml",
												"panel5-tab2-panels/panel2.yaml",
											},
											Panels: []*PanelInfo{
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel5-tab2-panels/panel1.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel5Tab2OnePanel",
													Note:       "Panel5Tab2Panel1 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel5Tab2Panel1 Markup",
												},
												&PanelInfo{
													SourcePath: "testyaml/simple_test/service2/panel5-tab2-panels/panel2.yaml",
													Level:      3,
													ID:         "",
													Name:       "Panel5Tab2TwoPanel",
													Note:       "Panel5Tab2Panel2 Note",
													Buttons:    []*ButtonInfo{},
													Tabs:       []*TabInfo(nil),
													Markup:     "Panel5Tab2Panel2 Markup",
												},
											},
										},
									},
									Markup: "",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := NewSlurper()
			got, err := sl.slurpApplication(tt.args.fpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("sl.slurpApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				gotbb, err := yaml.Marshal(got)
				if err != nil {
					t.Fatal(err)
				}
				wantbb, err := yaml.Marshal(tt.want)
				if err != nil {
					t.Fatal(err)
				}
				if string(gotbb) != string(wantbb) {
					t.Errorf("not equal")
					t.Errorf("sl.slurpApplication() = %s\n\nwant:\n%s", gotbb, wantbb)
				}
			}
			/*
				s := got.Services[0]
				t.Errorf("%#v", *s.Button.Panels[1].LPCCallIns[0])
				t.Errorf("%#v", *s.Button.Panels[1].LPCCallIns[1])
			*/
		})
	}
}
