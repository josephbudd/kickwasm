package about

import (
	"fmt"
)

const (
	versionBreaking = 4 // Each new version breaks backwards compatibility.
	versionFeature  = 0 // Each new version adds features. Retains backwards compatibility.
	versionPatch    = 0 // Each new version only fixes bugs. No added features. Retains backwards
)

// GetAuthorVersion returns the application's author and version.
func GetAuthorVersion() (author string, version []string) {
	author = "Joseph Budd"
	version = []string{
		`Contacts Example Application.`,
		`2019-04-06`,
		fmt.Sprintf("Version: %d.%d.%d", versionBreaking, versionFeature, versionPatch),
		"Updated to the experimental kickwasm version 4.02.",
	}
	return
}

// GetLicense returns the license
func GetLicense() (license []string) {
	license = []string{
		`The MIT License (MIT)`,

		`Copyright (c) 2018 Joseph Budd`,

		`Permission is hereby granted, free of charge, to any person obtaining a copy of
		this software and associated documentation files (the "Software"), to deal in
		the Software without restriction, including without limitation the rights to
		use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
		the Software, and to permit persons to whom the Software is furnished to do so,
		subject to the following conditions:`,

		`The above copyright notice and this permission notice shall be included in all
		copies or substantial portions of the Software.`,

		`THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
		FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
		COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
		IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
		CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`,
	}
	return
}
