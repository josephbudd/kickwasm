package templates

// BuildDotSH is the shell script that builds the renderer.
const BuildDotSH = `#!/bin/bash

GOARCH=wasm GOOS=js go build -o ../site/app.wasm main.go panels.go
`
