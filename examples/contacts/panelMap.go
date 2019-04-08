package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/filepaths"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

const (
	mainTemplate = "main.tmpl"
	headTemplate = "head.tmpl"
)

// serviceEmptyInsidePanelNamePathMap maps each markup panel template name to it's file path.
var serviceEmptyInsidePanelNamePathMap = map[string]map[string][]string{"About": map[string][]string{"AboutTabPanel": []string{"AboutButton", "AboutTabBarPanel", "AboutTab"}, "LiscenseTabPanel": []string{"AboutButton", "AboutTabBarPanel", "LiscenseTab"}, "RecordsTabPanel": []string{"AboutButton", "AboutTabBarPanel", "RecordsTab"}}, "AddAContact": map[string][]string{"AddContactPanel": []string{"AddButton"}}, "EditAContact": map[string][]string{"EditContactEditPanel": []string{"EditButton"}, "EditContactNotReadyPanel": []string{"EditButton"}, "EditContactSelectPanel": []string{"EditButton"}}, "RemoveAContact": map[string][]string{"RemoveContactConfirmPanel": []string{"RemoveButton"}, "RemoveContactNotReadyPanel": []string{"RemoveButton"}, "RemoveContactSelectPanel": []string{"RemoveButton"}}}

// serveMainHTML only serves up main.tmpl with all of the templates for your markup panels.
func serveMainHTML(w http.ResponseWriter) {
	var err error
	templateFolderPath := filepaths.GetTemplatePath()
	var t *template.Template
	t = template.New(mainTemplate)
	if t, err = t.ParseFiles(filepath.Join(templateFolderPath, mainTemplate)); err != nil {
		http.Error(w, err.Error(), 300)
		return
	}
	for _, namePathMap := range serviceEmptyInsidePanelNamePathMap {
		for name, folders := range namePathMap {
			folderPath := strings.Join(folders, string(os.PathSeparator))
			tpath := filepath.Join(templateFolderPath, folderPath, name+".tmpl")
			if t, err = t.ParseFiles(tpath); err != nil {
				http.Error(w, err.Error(), 300)
				return
			}
		}
	}
	// the head template which contains
	//  * any css imports
	//  * any javascript imports
	// needed for this applicaion
	tpath := filepath.Join(templateFolderPath, headTemplate)
	// it's ok if the template is not there
	// but if it's there use it.
	if _, err = os.Stat(tpath); os.IsNotExist(err) {
		// the template file does not exist so inform the developer.
		temp := fmt.Sprintf("%[1]s%[1]s define %[3]q %[2]s%[2]s<!-- You do not have a %[3]s file to import your css files. Feel free to add one in the render/template folder. -->%[1]s%[1]s end %[2]s%[2]s", "{", "}", headTemplate)
		if t, err = t.Parse(temp); err != nil {
			http.Error(w, err.Error(), 300)
			return
		}
	} else {
		// the file exists so parse it
		if t, err = t.ParseFiles(tpath); err != nil {
			http.Error(w, err.Error(), 300)
			return
		}
	}
	// do the template
	if err = t.ExecuteTemplate(w, mainTemplate, nil); err != nil {
		if !strings.Contains(err.Error(), "reset") {
			http.Error(w, err.Error(), 300)
		}
	}
}

