# kickbuild

Will only run in
    rendererprocess/
    mainprocess/
    application folder

Will need output to appname for OS.

``` go
var fname string
if len(GOOSFlag) == 0 {
    GOOSFlag = runtime.GOOS
}
if len(GOARCHFlag) == 0 {
    GOARCHFlag = runtime.GOARCH
}
switch GOOSFlag {
case "darwin":
  fname = appname
case "windows":
  fname = appname + ".exe"
default:
  fname = appname
}
```
