package refactor

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

var (
	buttonPanelsTabPanelsPath                 string
	buttonPanelsTabPanelsScrollPath           string
	buttonPanelsTabPanelsSwapButtonsPath      string
	buttonPanelsTabPanelsSwapTabsPath         string
	buttonPanelsTabPanelsReverseHomesPath  string
	buttonPanelsTabPanelsSwapPanelsPath       string
	buttonPanelsTabPanelsAddRemoveHomePath string
	buttonPanelsTabPanelsAddRemovePanelsPath  string
)

func init() {
	wd, _ := os.Getwd()
	testDatalPath := filepath.Join(wd, "testData")

	buttonPanelsTabPanelsPath = filepath.Join(testDatalPath, "button_panels_tab_panels.yaml")
	buttonPanelsTabPanelsScrollPath = filepath.Join(testDatalPath, "button_panels_tab_panels_scroll.yaml")
	buttonPanelsTabPanelsSwapButtonsPath = filepath.Join(testDatalPath, "button_panels_tab_panels_swap_buttons.yaml")
	buttonPanelsTabPanelsSwapTabsPath = filepath.Join(testDatalPath, "button_panels_tab_panels_swap_tabs.yaml")
	buttonPanelsTabPanelsReverseHomesPath = filepath.Join(testDatalPath, "button_panels_tab_panels_reverse_homes.yaml")
	buttonPanelsTabPanelsSwapPanelsPath = filepath.Join(testDatalPath, "button_panels_tab_panels_swap_panels.yaml")
	buttonPanelsTabPanelsAddRemoveHomePath = filepath.Join(testDatalPath, "button_panels_tab_panels_add_remove_home.yaml")
	buttonPanelsTabPanelsAddRemovePanelsPath = filepath.Join(testDatalPath, "button_panels_tab_panels_add_remove_panels.yaml")
}

func TestDifHomePositions(t *testing.T) {
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsReverseHomesPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		changes  *project.Builder
		original *project.Builder
	}
	tests := []struct {
		name        string
		args        args
		wantChanged bool
	}{
		{
			name: "TestDifHomePositions: No changes.",
			args: args{
				changes:  originalBuilder,
				original: originalBuilder,
			},
			wantChanged: false,
		},
		{
			name: "TestDifHomePositions: Changes.",
			args: args{
				changes:  changesBuilder,
				original: originalBuilder,
			},
			wantChanged: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChanged := DifHomePositionsButtons(tt.args.changes, tt.args.original); gotChanged != tt.wantChanged {
				t.Errorf("DifHomePositionsButtons() = %v\n\nwant %v", gotChanged, tt.wantChanged)
			}
		})
	}
}

func TestSwapTabs(t *testing.T) {
	var (
		additions = map[string]SpawnPath{}
		removals  = map[string]SpawnPath{}
		moves     = map[string]MoveSpawnPath{
			"Home3Button1Panel1Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel1Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home3Button1Panel1Tab",
					HVScroll: false,
				},
			},
			"Home3Button1Panel2Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel2Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home3Button1Panel2Tab",
					HVScroll: false,
				},
			},
			"Home4Button1Panel1Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel1Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home4Button1Panel1Tab",
					HVScroll: false,
				},
			},
			"Home4Button1Panel2Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel2Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home4Button1Panel2Tab",
					HVScroll: false,
				},
			},
		}

		positionsChanged = true
		scrollChanged    = false
	)
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsSwapTabsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testDifPanelPaths("TestSwapTabs", changesBuilder, originalBuilder, removals, additions, moves, positionsChanged, scrollChanged, t)
}

func TestSwapButtons(t *testing.T) {
	var (
		additions = map[string]SpawnPath{}
		removals  = map[string]SpawnPath{}
		moves     = map[string]MoveSpawnPath{
			"Home1ButtonPanel1ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel1Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home1ButtonPanel1Button",
					HVScroll: false,
				},
			}, "Home1ButtonPanel2ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home1ButtonPanel2Button",
					HVScroll: false,
				},
			},
			"Home2Button1Panel1ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel1Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home2Button1Panel1Button",
					HVScroll: false,
				},
			},
			"Home2Button1Panel2ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel2Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home2Button1Panel2Button",
					HVScroll: false,
				},
			},
		}
		positionsChanged = true
		scrollChanged    = false
	)
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsSwapButtonsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testDifPanelPaths("TestSwapButtons", changesBuilder, originalBuilder, removals, additions, moves, positionsChanged, scrollChanged, t)
}

func TestSwapPanels(t *testing.T) {
	var (
		additions = map[string]SpawnPath{}
		removals  = map[string]SpawnPath{}
		moves     = map[string]MoveSpawnPath{
			"Home1ButtonPanel1ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel1Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
					HVScroll: false,
				},
			},
			"Home1ButtonPanel2ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel1Button",
					HVScroll: false,
				},
			},
			"Home2Button1Panel1ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel1Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel2Button",
					HVScroll: false,
				},
			},
			"Home2Button1Panel2ButtonPanel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel2Button",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home2Button/Home2Button1Panel/Home2Button1Panel1Button",
					HVScroll: false,
				},
			},
			"Home3Button1Panel1Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel1Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel2Tab",
					HVScroll: false,
				},
			},
			"Home3Button1Panel2Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel2Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home3Button/Home3Button1Panel/Home3Button1Panel1Tab",
					HVScroll: false,
				},
			},
			"Home4Button1Panel1Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel1Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel2Tab",
					HVScroll: false,
				},
			},
			"Home4Button1Panel2Tab1Panel": MoveSpawnPath{
				From: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel2Tab",
					HVScroll: false,
				},
				To: SpawnPath{
					Spawn:    false,
					Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel1Tab",
					HVScroll: false,
				},
			},
		}
		positionsChanged = true
		scrollChanged    = false
	)
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsSwapPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testDifPanelPaths("TestSwapPanels", changesBuilder, originalBuilder, removals, additions, moves, positionsChanged, scrollChanged, t)
}

func TestAddRemoveHome(t *testing.T) {
	var (
		additions = map[string]SpawnPath{
			"Home5Button1Panel1Tab1Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home5Button/Home5Button1Panel/Home5Button1Panel1Tab",
				HVScroll: false,
			},
			"Home5Button1Panel2Tab1Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home5Button/Home5Button1Panel/Home5Button1Panel2Tab",
				HVScroll: false,
			},
		}
		removals = map[string]SpawnPath{
			"Home4Button1Panel1Tab1Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel1Tab",
				HVScroll: false,
			},
			"Home4Button1Panel2Tab1Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home4Button/Home4Button1Panel/Home4Button1Panel2Tab",
				HVScroll: false,
			},
		}
		moves            = map[string]MoveSpawnPath{}
		positionsChanged = true
		scrollChanged    = false
	)
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsAddRemoveHomePath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testDifPanelPaths("TestAddRemoveHome", changesBuilder, originalBuilder, removals, additions, moves, positionsChanged, scrollChanged, t)
}

func TestAddRemovePanel(t *testing.T) {
	var (
		additions = map[string]SpawnPath{
			"Home1ButtonPanel2Button2Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
				HVScroll: false,
			},
			"Home1ButtonPanel2Button3Panel": SpawnPath{
				Spawn:    false,
				Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
				HVScroll: false,
			},
		}
		removals = map[string]SpawnPath{
			"Home1ButtonPanel2ButtonPanel": SpawnPath{
				Spawn:    false,
				Path:     "Home1Button/Home1Button1Panel/Home1ButtonPanel2Button",
				HVScroll: false,
			},
		}
		moves            = map[string]MoveSpawnPath{}
		positionsChanged = true
		scrollChanged    = false
	)
	// original
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	// changes
	sl = slurp.NewSlurper()
	changesBuilder, err := sl.Gulp(buttonPanelsTabPanelsAddRemovePanelsPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = changesBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testDifPanelPaths("TestAddRemovePanel", changesBuilder, originalBuilder, removals, additions, moves, positionsChanged, scrollChanged, t)
}

/* todo */

func TestBuildPanelNamePathMap(t *testing.T) {
	var (
		buildPanelNamePathMapWant = map[string]project.SpawnFolders{
			"Home1ButtonPanel1ButtonPanel": project.SpawnFolders{
				Position: 0,
				Spawn:    false,
				Folders:  []string{"Home1Button", "Home1Button1Panel", "Home1ButtonPanel1Button"},
				HVScroll: false,
			},
			"Home1ButtonPanel2ButtonPanel": project.SpawnFolders{
				Position: 1,
				Spawn:    false,
				Folders:  []string{"Home1Button", "Home1Button1Panel", "Home1ButtonPanel2Button"},
				HVScroll: false,
			},
			"Home2Button1Panel1ButtonPanel": project.SpawnFolders{
				Position: 1,
				Spawn:    false,
				Folders:  []string{"Home2Button", "Home2Button1Panel", "Home2Button1Panel1Button"},
				HVScroll: false,
			},
			"Home2Button1Panel2ButtonPanel": project.SpawnFolders{
				Position: 2,
				Spawn:    false,
				Folders:  []string{"Home2Button", "Home2Button1Panel", "Home2Button1Panel2Button"},
				HVScroll: false,
			},
			"Home3Button1Panel1Tab1Panel": project.SpawnFolders{
				Position: 2,
				Spawn:    false,
				Folders:  []string{"Home3Button", "Home3Button1Panel", "Home3Button1Panel1Tab"},
				HVScroll: false,
			},
			"Home3Button1Panel2Tab1Panel": project.SpawnFolders{
				Position: 3,
				Spawn:    false,
				Folders:  []string{"Home3Button", "Home3Button1Panel", "Home3Button1Panel2Tab"},
				HVScroll: false,
			},
			"Home4Button1Panel1Tab1Panel": project.SpawnFolders{
				Position: 3,
				Spawn:    false,
				Folders:  []string{"Home4Button", "Home4Button1Panel", "Home4Button1Panel1Tab"},
				HVScroll: false,
			},
			"Home4Button1Panel2Tab1Panel": project.SpawnFolders{
				Position: 4,
				Spawn:    false,
				Folders:  []string{"Home4Button", "Home4Button1Panel", "Home4Button1Panel2Tab"},
				HVScroll: false,
			},
		}
	)
	sl := slurp.NewSlurper()
	originalBuilder, err := sl.Gulp(buttonPanelsTabPanelsScrollPath)
	if err != nil {
		t.Fatal(err)
	}
	_ = originalBuilder.ToHTMLNode("", false)
	if err != nil {
		t.Fatal(err)
	}
	testbuildPanelNamePathMap("TestBuildPanelNamePathMap", originalBuilder, buildPanelNamePathMapWant, t)
}

func testbuildPanelNamePathMap(name string, b *project.Builder, want map[string]project.SpawnFolders, t *testing.T) {
	t.Run(name, func(t *testing.T) {
		got := buildPanelNamePathMap(b)
		if eq, fixedGot := spawnFoldersEqual(got, want); !eq {
			t.Errorf("buildPanelNamePathMap(%s) got = %#v\n\nwant %#v", name, fixedGot, want)
		}
	})
}

func spawnFoldersEqual(got map[string]*project.SpawnFolders, want map[string]project.SpawnFolders) (equal bool, fixedGot map[string]project.SpawnFolders) {
	fixedGot = make(map[string]project.SpawnFolders, len(got))
	for k, v := range got {
		fixedGot[k] = *v
	}
	if len(got) != len(want) {
		return
	}
	for pname, sfGot := range fixedGot {
		var sfWant project.SpawnFolders
		var found bool
		if sfWant, found = want[pname]; !found {
			return
		}
		if sfGot.Position != sfWant.Position {
			return
		}
		if sfGot.Spawn != sfWant.Spawn {
			return
		}
		if sfGot.HVScroll != sfWant.HVScroll {
			return
		}
		if len(sfGot.Folders) != len(sfWant.Folders) {
			return
		}
		for i, f := range sfGot.Folders {
			if sfWant.Folders[i] != f {
				return
			}
		}
	}
	equal = true
	return
}

func testDifPanelPaths(name string, changes, original *project.Builder, wantRemovals, wantAdditions map[string]SpawnPath, wantMoves map[string]MoveSpawnPath, wantPositionsChanged, wantscrollChanged bool, t *testing.T) {
	t.Run(name, func(t *testing.T) {
		gotRemovals, gotAdditions, gotMoves, gotPositionsChanged, gotscrollChanged := DifPanelPaths(changes, original)
		if !reflect.DeepEqual(gotRemovals, wantRemovals) {
			t.Errorf("DifPanelPaths(%s) gotRemovals = %#v\n\nwant %#v", name, gotRemovals, wantRemovals)
		}
		if !reflect.DeepEqual(gotAdditions, wantAdditions) {
			t.Errorf("DifPanelPaths(%s) gotAdditions = %#v\n\nwant %#v", name, gotAdditions, wantAdditions)
		}
		if !reflect.DeepEqual(gotMoves, wantMoves) {
			t.Errorf("DifPanelPaths(%s) gotMoves = %#v\n\nwant %#v", name, gotMoves, wantMoves)
		}
		if !reflect.DeepEqual(gotPositionsChanged, wantPositionsChanged) {
			t.Errorf("DifPanelPaths(%s) gotPositionsChanged = %#v\n\nwant %#v", name, gotPositionsChanged, wantPositionsChanged)
		}
		if gotscrollChanged != wantscrollChanged {
			t.Errorf("DifPanelPaths(%s) gotscrollChanged = %#v\n\nwant %#v", name, gotscrollChanged, wantscrollChanged)
		}
	})
}
