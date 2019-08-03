package cases

import (
	"bytes"
	"strings"
)

const (
	underscoreRune = '_'
	spaceString    = " "
	emptyString    = ""
)

// CamelCase consverts a string to camel case.
// ex: "abc def ghi" becomes "AbcDefGhi"
func CamelCase(s string) string {
	ss := strings.Split(s, spaceString)
	for i, part := range ss {
		ss[i] = strings.ToUpper(part[0:1]) + part[1:]
	}
	return toJSVarName(strings.Join(ss, emptyString))
}

// LowerCamelCase consverts a string to lower camel case.
// ex: "abc def ghi" becomes "abcDefGhi"
func LowerCamelCase(s string) string {
	ss := strings.Split(s, spaceString)
	ss[0] = strings.ToLower(ss[0][0:1]) + ss[0][1:]
	l := len(ss)
	for i := 1; i < l; i++ {
		ss[i] = strings.ToUpper(ss[i][0:1]) + ss[i][1:]
	}
	return toJSVarName(strings.Join(ss, emptyString))
}

// ToGoPackageName makes a word a valid go package name.
func ToGoPackageName(s string) (newName string) {
	return strings.ToLower(s)
}

// toJSVarName removes chars unless a-z, A-Z, 0-9, underscore.
// If s begins with a digit (0-9)
//  then the new name is prefixed with an underscore.
func toJSVarName(s string) string {
	w := new(bytes.Buffer)
	ch := s[:1]
	if ch >= "0" && ch <= "9" {
		// Source name begins with a digit so prefix the new name with an underscore
		w.WriteRune(underscoreRune)
	}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			w.WriteRune(r)
			continue
		}
		if r >= 'A' && r <= 'Z' {
			w.WriteRune(r)
			continue
		}
		if r >= '0' && r <= '9' {
			w.WriteRune(r)
			continue
		}
		if r == underscoreRune {
			w.WriteRune(r)
		}
		// ignore other chars
	}
	return w.String()
}
