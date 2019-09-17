#!/bin/bash

# Paths.
startwd=$PWD
appwd="${PWD%/*}"
prefix="${appwd##*/}"
authorwd="${appwd%/*}"
step=0

sitepack_package_name="${prefix}sitepack"
sitepack_output_package_path="${authorwd}/${sitepack_package_name}"

spawnpack_package_name="spawnpack"
spawnpack_package_path="${appwd}/renderer/${spawnpack_package_name}"
spawn_templates_folder="spawnTemplates"
spawnpack_templates_path="${appwd}/site/${spawn_templates_folder}"

# Remove the old spawn package if it's there.
if [ -d "${spawnpack_package_path}" ]
then
    echo ""
    (( step += 1 ))
    echo "STEP ${step}:"
    echo "REMOVE YOUR PREVIOUS BUILD OF ${spawnpack_package_path}"
    echo "rm -r ${spawnpack_package_path}"
    rm -r "${spawnpack_package_path}"
fi

# Write the new spawn package source code.
# This is for the renderer so it must be a file store not a reader.
echo ""
(( step += 1 ))
echo "STEP ${step}:"
echo "WRITE THE SOURCE CODE FOR YOUR NEW ${spawnpack_package_name} PACKAGE."
echo " * The ${spawnpack_package_name} package is your renderer's spawn html templates."
echo "cd ${appwd}/site"
cd "${appwd}/site"
echo "kickpack -nu -o=${spawnpack_package_path} ./${spawn_templates_folder}"
kickpack -nu -o="${spawnpack_package_path}" "./${spawn_templates_folder}"
if [ $? -gt 0 ]
then
    echo "Oops! Job.Ended."
    exit
fi
echo "cd ${startwd}"
cd "${startwd}"

echo " * Success. The source code for your new ${spawnpack_package_name} package is written."

# Build the wasm.
echo ""
(( step += 1 ))
echo "STEP ${step}:"
echo "BUILD THE RENDERER GO CODE INTO WEB ASSEMBLY CODE AT ../site/app.wasm"
echo "GOARCH=wasm GOOS=js go build -o ../site/app.wasm Main.go panels.go"
GOARCH=wasm GOOS=js go build -o ../site/app.wasm Main.go panels.go
if [ $? -gt 0 ]
then
    echo "Oops!"
    exit
fi

echo " * Success. The renderer go code is compiled into web assembly code at ../site/app.wasm"
    
# Remove the old sitepack package if it's there.
if [ -d "${sitepack_output_package_path}" ]
then
    echo ""
    (( step += 1 ))
    echo "STEP ${step}:"
    echo "REMOVE YOUR PREVIOUS BUILD OF ${sitepack_package_name}"
    echo "rm -r ${sitepack_output_package_path}"
    rm -r "${sitepack_output_package_path}"
fi

# Pack ./site and .Http.yaml into a new sitepack package.
# This will be a file reader not a file store.
echo ""
(( step += 1 ))
echo "STEP ${step}:"
echo "WRITE THE ${sitepack_package_name} PACKAGE SOURCE CODE."
echo " * The ${sitepack_package_name} package will pretend to be a file store"
echo "     but it will actually just read the files"
echo "     in the site folder and return their contents."
echo " * See func serveFileStore in Serve.go."
echo "cd ${appwd}"
cd ..
echo "kickpack -nu -nopack -o=${sitepack_output_package_path} ./site ./Http.yaml"
kickpack -nu -nopack -o="${sitepack_output_package_path}" ./site ./Http.yaml
if [ $? -gt 0 ]
then
    echo "Oops! That shouldn't happen."
    exit
fi

# Build the new site pack package.
echo " * Success. The ${sitepack_package_name} package source code has been written."

echo ""
(( step += 1 ))
echo "STEP ${step}:"
echo "BUILD THE ${sitepack_package_name} PACKAGE."
echo "cd ${sitepack_output_package_path}"
cd "$sitepack_output_package_path"
echo "go build"
go build
if [ $? -gt 0 ]
then
    echo "Oops! Your code is broken."
    exit
fi

echo " * Success."
echo " * You have successfully compiled the ${sitepack_package_name} package object code."

echo ""
(( step += 1 ))
echo "STEP ${step}:"
echo "BUILD THE MAIN PROCESS GO CODE INTO THE EXECUTABLE ${prefix}."
echo "   You will do so with the following 2 commands..."
echo "   cd .."
echo "   go build"
