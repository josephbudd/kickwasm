#!/bin/bash

# build wasm
echo "Building your wasm into ../site/app.wasm"
GOARCH=wasm GOOS=js go build -o ../site/app.wasm main.go panels.go
if [ $? -gt 0 ]
then
    echo "Oops! It's broken."
    exit
fi

echo ""
echo "Great! Your wasm has been compiled."

# collect paths
startwd=$PWD
appwd="${PWD%/*}"
prefix="${appwd##*/}"
authorwd="${appwd%/*}"
sitepackagename="${prefix}sitepack"
sitepackpath="${authorwd}/${sitepackagename}"

# pack ./site and .http.yaml into a new sitepack package
echo ""
echo "Now its time to write the source code for your new ${sitepackagename} package."
echo "The ${sitepackagename} package is your applications renderer process."
echo "( The stuff that gets loaded into the browser. )"
echo "This could take a while."
echo "cd ${appwd}"
cd ..
echo "kickpack -o ${sitepackpath} ./site ./http.yaml"
kickpack -o "${sitepackpath}" ./site ./http.yaml
if [ $? -gt 0 ]
then
    echo "Oops! Job Ended."
    exit
fi

# build the new package
echo ""
echo "Finally! Now its time to build your new ${sitepackagename} package."
echo "cd ${sitepackpath}"
cd "$sitepackpath"
echo "go build"
go build
if [ $? -gt 0 ]
then
    echo "Oops! Job Ended."
    exit
fi

echo ""
echo "You've done it!"
echo "The package at ${sitepackpath} contains the files from your renderer process."
