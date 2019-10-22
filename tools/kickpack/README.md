[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickpack)](https://goreportcard.com/report/github.com/josephbudd/kickpack)

kickpack is part of the kickwasm tool chain. It is only used by the framework's renderer build.sh script.

## Updated 2019-07-25

These changes were made for the build.sh script.

* Added the exit code.
* Added the -nu flag which prevents usage from being displayed following an error message.

To see how kickpack works just follow the steps under testing below.

If you just want to see what kind of package kickpack creates I posted tree/master/testdata/deep/gooutput/gooutput.go at the bottom of this page.

## If you want to test kickpack

First run create_output_test.go
Then run check_output_test.go

example:

``` shell

$ go test create_output_test.go
$ go test check_output_test.go

```

The output will be at ./testdata/deep/gooutput/gooutput.go
It just a tiny file for you to look at it in a text editor.

### Warning

When you use kickpack you will not want to look at the output file using a text editor.
That is because the file will probably be monsterously huge.

## A tiny example kickpack output file

This output file is safe to look using a text editor because it was made from tiny little text files.
This file is located at tree/master/testdata/deep/gooutput/gooutput.go.

``` go

package gooutput

Contents returns the contents of the file at path and if found.
func Contents(path string) (contents []byte, found bool) {
  contents, found = fileStore[path]
  return
}

// Paths returns a slice of the file paths.
func Paths() (paths []string) {
  l := len(fileStore)
  paths = make([]string, 0, l)
  for k := range fileStore {
    paths = append(paths, k)
  }
  return
}


// fileStore is a store of various files.
var fileStore =  map[string][]byte{
    "me.txt":[]byte{0x6d, 0x65, 0xa},
    "src1/css/images.css":[]byte{0x69, 0x6d, 0x67, 0x2e, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x7b, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x31, 0x70, 0x78, 0x20, 0x72, 0x65, 0x64, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x3b, 0x7d},
    "../src2/text/something.txt":[]byte{0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67},
    "../src2/text/hello_world.txt":[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64},
    "src1/css/forms.css":[]byte{0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x20, 0x7b, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x72, 0x65, 0x64, 0x3b, 0x7d},
}

```

## To install

``` shell

$ go get -u github.com/josephbudd/kickpack
$ cd ~/go/src/github.com/josephbudd/kickpack
$ go install

```
