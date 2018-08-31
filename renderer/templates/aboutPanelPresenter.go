package templates

// AboutPanelPresenter is the main process about panel presenter.
const AboutPanelPresenter = `package AboutPanel

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/josephbudd/kicknotjs"

	"{{.ApplicationGitPath}}{{.ImportMainProcessServicesAbout}}"
)

// Presenter writes to the panel
type Presenter struct {
	notjs     *kicknotjs.NotJS

	// Declare your presenter members here.

	output js.Value
}

// defineMembers is where you define your members by their html elements.
func (presenter *Presenter) defineMembers() {

	// Define the presenter members.

	// text output
	presenter.output = presenter.notjs.GetElementByID("output")
}

// Define your presenter functions.

// displayReleases displays the releases
func (presenter *Presenter) displayReleases(thisVersion string, releases map[string]map[string][]string) {
	notjs := presenter.notjs
	div := notjs.GetElementByID("{{.ReleasesTabPanelInnerID}}")
	h2 := notjs.CreateElement("h2")
	notjs.SetInnerText(h2, "Current Release.")
	notjs.AppendChild(div, h2)
	h3 := notjs.CreateElement("h3")
	notjs.SetInnerText(h3, fmt.Sprintf("Version %s", thisVersion))
	notjs.AppendChild(div, h3)
	versionRelease := releases[thisVersion]
	// versionRelease = map[string][]string
	for release, lines := range versionRelease {
		h4 := notjs.CreateElement("h4")
		notjs.SetInnerText(h4, release)
		notjs.AppendChild(div, h4)
		for _, line := range lines {
			p := notjs.CreateElement("p")
			notjs.SetInnerText(p, line)
			notjs.AppendChild(div, p)
		}
	}
	h2 = notjs.CreateElement("h2")
	notjs.SetInnerText(h2, "Earlier Releases.")
	notjs.AppendChild(div, h2)
	for version, versionRelease := range releases {
		// versionRelease is map[string][]string
		if version != thisVersion {
			h3 := notjs.CreateElement("h3")
			notjs.SetInnerText(h3, version)
			notjs.AppendChild(div, h3)
			for release, lines := range versionRelease {
				h4 := notjs.CreateElement("h4")
				notjs.SetInnerText(h4, release)
				notjs.AppendChild(div, h4)
				for _, line := range lines {
					p := notjs.CreateElement("p")
					notjs.SetInnerText(p, line)
					notjs.AppendChild(div, p)
				}
			}
		}
	}
}

func (presenter *Presenter) displayContributors(contributors map[string]string) {
	notjs := presenter.notjs
	div := notjs.GetElementByID("{{.ContributorsTabPanelInnerID}}")
	table := notjs.CreateElement("table")
	for name, comment := range contributors {
		tr := notjs.CreateElement("tr")
		td := notjs.CreateElement("td")
		b := notjs.CreateElement("b")
		notjs.SetInnerText(b, name)
		notjs.AppendChild(td, b)
		notjs.AppendChild(tr, td)
		td = notjs.CreateElement("td")
		notjs.SetInnerText(td, comment)
		notjs.AppendChild(tr, td)
		notjs.AppendChild(table, tr)
	}
	notjs.AppendChild(div, table)
}

func (presenter *Presenter) displayCredits(credits []about.Credit) {
	/*
		type Credit struct {
			Name        string
			Description string
			URL         string
		}
	*/
	notjs := presenter.notjs
	div := notjs.GetElementByID("{{.CreditsTabPanelInnerID}}")
	for _, credit := range credits {
		p := notjs.CreateElement("p")
		if len(credit.URL) > 0 {
			a := notjs.CreateElement("a")
			notjs.SetAttribute(a, "href", credit.URL)
			notjs.SetAttribute(a, "target", "_blank")
			b := notjs.CreateElement("b")
			notjs.SetInnerText(b, credit.Name)
			notjs.AppendChild(a, b)
			notjs.AppendChild(p, a)
			t := notjs.CreateTextNode(" " + credit.Description)
			notjs.AppendChild(p, t)
		} else {
			b := notjs.CreateElement("b")
			notjs.SetInnerText(b, credit.Name)
			notjs.AppendChild(p, b)
			t := notjs.CreateTextNode(" " + credit.Description)
			notjs.AppendChild(p, t)
		}
		notjs.AppendChild(div, p)
	}
}

func (presenter *Presenter) displayLicenses(licenses []about.License) {
	/*
		type License struct {
			Software string
			License  string
			Location string
		}
	*/
	notjs := presenter.notjs
	inner := notjs.GetElementByID("{{.LicensesTabPanelInnerID}}")
	for _, license := range licenses {
		div := notjs.CreateElement("div")
		notjs.ClassListAddClass(div, "about-license-wrapper")
		h2 := notjs.CreateElement("h2")
		notjs.SetInnerText(h2, license.Software)
		notjs.AppendChild(div, h2)
		p := notjs.CreateElement("p")
		a := notjs.CreateElement("a")
		notjs.SetAttributeHref(a, license.Location)
		notjs.SetInnerText(a, license.Location)
		notjs.AppendChild(p, a)
		notjs.AppendChild(div, p)
		h3 := notjs.CreateElement("h3")
		notjs.SetInnerText(h3, "License")
		notjs.AppendChild(div, h3)
		lines := presenter.fixLicense(license.License)
		for _, line := range lines {
			p := notjs.CreateElement("p")
			notjs.SetInnerText(p, line)
			notjs.AppendChild(div, p)
		}
		notjs.AppendChild(inner, div)
	}
}

func (presenter *Presenter) fixLicense(license string) []string {
	s := strings.Replace(license, "\r\n", "\n", -1)
	s = strings.Replace(license, "\r", "\n", -1)
	lines := strings.Split(s, "\n")
	newlines := make([]string, 0, len(lines))
	adding := true
	i := -1
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) > 0 {
			if adding {
				newlines = append(newlines, line)
				i++
				adding = false
			} else {
				newlines[i] = newlines[i] + " " + line
			}
		} else {
			adding = true
		}
	}
	return newlines
}
`
