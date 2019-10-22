package main

import (
	"testing"

	"github.com/josephbudd/kickwasm/tools/kickpack/testdata/deep/gooutput"
)

func Test2(t *testing.T) {
	want := map[string]string{
		"../src2/text/hello_world.txt": "hello world",
		"me.txt":                       "me\n",
		"src1/css/images.css":          "img.fancy{border: 1px red solid;}",
		"../src2/text/something.txt":   "something",
		"src1/css/forms.css":           "form.fancy {color:red;}",
	}
	for path, wantContents := range want {
		if gotContents, found := gooutput.Contents(path); !found {
			t.Errorf("cant find %q in output", path)
		} else {
			if wantContents != string(gotContents) {
				t.Errorf("want %q got %q", wantContents, string(gotContents))
			}
		}
	}
}
