package templates

import (
	"strings"
)

// GetServicesAboutGo returns the mainprocess/services/about.go file.
func GetServicesAboutGo() string {
	return strings.Replace(aboutGo, "{{backTick}}", "`", -1)
}

// AboutGo is the mainprocess/about/about.go
const aboutGo = `package about

import "fmt"

const (
	major  uint = 0
	minor  uint = 1
	bugfix uint = 0
)

// String stringifys version.
func String() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, bugfix)
}

// GetReleases returns the history of changes to this application.
// Be careful editing this.
// Make sure you understand this data structure.
func GetReleases() map[string]map[string][]string {
	return map[string]map[string][]string{
		"0.1.0": map[string][]string{
			"Once Upon A Time...": []string{
				{{backTick}}I'm using a separate main process and renderer. The renderer is a obviously a browser. The separation means no data races between the two.{{backTick}},
				{{backTick}}The main process and renderer each have their own message dispatcher connected via web sockets.{{backTick}},
				{{backTick}}The user interface is started and working properly as are the services and models in the main process, however, my ALSA model needs improvement.{{backTick}},
				{{backTick}}I separated the reference tab panel information into individual tabs.{{backTick}},
				{{backTick}}Instead of a single wpm there are now two wpms. One wpm for copying and another wpm for keying.{{backTick}},
				{{backTick}}The main process now terminates if the web socket connection to the renderer terminates.{{backTick}},
				{{backTick}}The render now terminates if the web socket connection to the main process terminates.{{backTick}},
				{{backTick}}I pulled the message code out. Then i rewrote it again and again so its no longer messages.{{backTick}},
				{{backTick}}Now its local procedure calls. I named the package "lpc".{{backTick}},
			},
			"FYI...": []string{
				{{backTick}}You can edit all of the about information you see under the about tab.{{backTick}},
				{{backTick}}It is in /mainprocess/about/about.go{{backTick}},
				{{backTick}}But be careful not to break anything.{{backTick}},
			},
		},
		"0.1.1": map[string][]string{
			"More stuff.": []string{
				{{backTick}}This is just something to give you a better idea of how this data structure looks and works.{{backTick}},
				{{backTick}}You will need to understand this data structure before you edit it.{{backTick}},
				{{backTick}}Again, you can edit all of the about information you see under the about tab.{{backTick}},
				{{backTick}}It is in /mainprocess/about/about.go{{backTick}},
				{{backTick}}Happy coding!{{backTick}},
			},
		},
	}
}

// GetContributors returns those who contributed to this application in any way.
// This is where the names of each man or woman involve in this application should be listed.
// Each entry is a name and a saying or comment.
func GetContributors() map[string]string {
	return map[string]string{
		"Your Name":            "Something you want to say.",
		"Another Coder's Name": "Something he or she has to say.",
	}
}

// Credit is information about things and people you need to acknowledge.
type Credit struct {
	Name        string
	Description string
	URL         string
}

// GetCredits returns the credits.
// Be verbose! Show how appreciatative and good natured you are.
// Mention any software, speakers and writers that helped you build this application in any way.
// Don't forget people and things that helped you in more remote ways as well.
func GetCredits() []Credit {
	return []Credit{
		Credit{
			Name:        "kickwasm",
			Description: "A way to kick start your golang application.",
			URL:         "https://www.github.com/josephbudd/kickwasm",
		},
		Credit{
			Name:        "bolt",
			Description: "the bolt database.",
			URL:         "https://www.github.com/boltdb/bolt",
		},
	}
}

// License is software license information.
type License struct {
	Software string
	License  string
	Location string
}

// GetLicenses returns the licenses.
// Add your license info to this list.
// Add license info for other software that you use in this application.
func GetLicenses() []License {
	return []License{
		License{
			Software: "Some of the code in this application was liberally copied from go source files which are distributed under the following license.",
			License: {{backTick}}Copyright (c) 2009 The Go Authors. All rights reserved.
			
			Redistribution and use in source and binary forms, with or without
			modification, are permitted provided that the following conditions are
			met:
			
			   * Redistributions of source code must retain the above copyright
			notice, this list of conditions and the following disclaimer.
			   * Redistributions in binary form must reproduce the above
			copyright notice, this list of conditions and the following disclaimer
			in the documentation and/or other materials provided with the
			distribution.
			   * Neither the name of Google Inc. nor the names of its
			contributors may be used to endorse or promote products derived from
			this software without specific prior written permission.
			
			THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
			"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
			LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
			A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
			OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
			SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
			LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
			DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
			THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
			(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
			OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.{{backTick}},
			Location: "https://www.golang.org/",
		},
		License{
			Software: "bolt",
			License: {{backTick}}The MIT License (MIT)
			
			Copyright (c) 2013 Ben Johnson
			
			Permission is hereby granted, free of charge, to any person obtaining a copy of
			this software and associated documentation files (the "Software"), to deal in
			the Software without restriction, including without limitation the rights to
			use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
			the Software, and to permit persons to whom the Software is furnished to do so,
			subject to the following conditions:
			
			The above copyright notice and this permission notice shall be included in all
			copies or substantial portions of the Software.
			
			THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
			IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
			FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
			COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
			IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
			CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.{{backTick}},
			Location: "https://www.github.com/boltdb/bolt",
		},
		License{
			Software: "kickwasm",
			License: {{backTick}}The MIT License (MIT)
			
			Copyright (c) 2017 Joseph Budd
			
			Permission is hereby granted, free of charge, to any person obtaining a copy of
			this software and associated documentation files (the "Software"), to deal in
			the Software without restriction, including without limitation the rights to
			use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
			the Software, and to permit persons to whom the Software is furnished to do so,
			subject to the following conditions:
			
			The above copyright notice and this permission notice shall be included in all
			copies or substantial portions of the Software.
			
			THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
			IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
			FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
			COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
			IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
			CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
			{{backTick}},
			Location: "https://www.github.com/josephbudd/kickwasm",
		},
	}
}

`
