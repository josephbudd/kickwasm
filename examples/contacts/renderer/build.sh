#!/bin/bash

GOARCH=wasm GOOS=js go build -o ../site/app.wasm main.go panels.go
