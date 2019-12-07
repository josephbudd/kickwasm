package renderer

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createProof(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	data := struct {
		HomeButtonNames      string
		ButtonNamePanelNames string
		TabNamePanelNames    string

		PanelNameButtonNames string
		PanelNameTabNames    string
	}{
		HomeButtonNames:      fmt.Sprintf("%#v", builder.ProofsHomeButtonNames()),
		ButtonNamePanelNames: fmt.Sprintf("%#v", builder.ProofsButtonNamePanelNames()),
		TabNamePanelNames:    fmt.Sprintf("%#v", builder.ProofsTabNamePanelNames()),

		PanelNameButtonNames: fmt.Sprintf("%#v", builder.ProofsPanelNameButtonNames()),
		PanelNameTabNames:    fmt.Sprintf("%#v", builder.ProofsPanelNameTabNames()),
	}
	folderpaths := appPaths.GetPaths()
	filenames := appPaths.GetFileNames()
	fname := filenames.ProofsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererProofs, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ProofsGo, data, appPaths)
	return

}
