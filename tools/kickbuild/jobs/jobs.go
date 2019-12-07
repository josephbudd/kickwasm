package jobs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/tools/common"
	"github.com/josephbudd/kickwasm/tools/script"
)

// RemoveOldSpawnPackPackage removes a spawn package.
func RemoveOldSpawnPackPackage(packagePath string) (err error) {
	if !common.PathFound(packagePath) {
		return
	}
	printNextStep("REMOVE THE OLD " + filepath.Base(packagePath) + ".")
	if err = os.RemoveAll(packagePath); err != nil {
		printError(err)
		return
	}
	printSuccess()
	return
}

// WriteNewSpawnPackage builds a new spawn package.
func WriteNewSpawnPackage(packagePath, templatesFolder, sitePath string) (err error) {
	packageName := filepath.Base(packagePath)
	printNextStep("WRITE THE NEW " + packageName + " SOURCE CODE.")
	PrintLine(" * The " + packageName + " package contains your site folder's spawn html templates.")
	args := []string{"-nu", "-o=" + packagePath, templatesFolder}
	var dump []byte
	PrintLine("packagePath is " + packagePath)
	if dump, err = script.RunDump(nil, sitePath, "kickpack", args...); err != nil {
		printDump(dump, err)
		return
	}
	printSuccess()
	return
}

// BuildWASM builds the wasm code.
func BuildWASM(rootFolderPath, sitePath, frameworkPath, rendererPath string) (err error) {
	printNextStep("BUILD THE RENDERER GO CODE INTO WASM AT /site/app.wasm")
	//GOARCH=wasm GOOS=js go build -o ../site/app.wasm Main.go panels.go
	env := os.Environ()
	env = append(env, "GOARCH=wasm")
	env = append(env, "GOOS=js")
	// args := []string{"build", "-o", filepath.Join("site", "app.wasm"), filepath.Join("rendererprocess", "Main.go"), filepath.Join("rendererprocess", "panels.go")}
	args := []string{"build", "-o", filepath.Join(sitePath, "app.wasm"), filepath.Join(rendererPath, "Main.go")}
	var dump []byte
	if dump, err = script.RunDump(env, "", "go", args...); err != nil {
		fixPrintDump(rootFolderPath, dump, err)
		return
	}
	printSuccess()
	return
}

// RemoveOldSitePackPackage removes a spawn package.
func RemoveOldSitePackPackage(packagePath string) (err error) {
	packageName := filepath.Base(packagePath)
	if !common.PathFound(packagePath) {
		return
	}
	printNextStep("REMOVE THE OLD " + packageName + " PACKAGE.")
	if err = os.RemoveAll(packagePath); err != nil {
		printError(err)
		return
	}
	printSuccess()
	return
}

// WriteSitePackPackageDontPack writes the site package.
func WriteSitePackPackageDontPack(appWdPath, packagePath string) (err error) {
	packageName := filepath.Base(packagePath)
	printNextStep("WRITE THE NEW " + packagePath + " PACKAGE SOURCE CODE.")
	PrintLine(" * The NEW " + packageName + " package will pretend to be a file store")
	PrintLine("     but it will actually just read the files")
	PrintLine("     in the site folder and return their contents.")
	PrintLine(" * See func serveFileStore in Serve.go.")

	args := []string{"-nu", "-nopack", "-o=" + packagePath, filepath.Join(".", "site"), filepath.Join(".", "Http.yaml")}
	var dump []byte
	if dump, err = script.RunDump(nil, appWdPath, "kickpack", args...); err != nil {
		printDump(dump, err)
		return
	}
	printSuccess()
	return
}

// WriteSitePackPackagePack writes the site package.
func WriteSitePackPackagePack(appWdPath, packagePath string) (err error) {
	packageName := filepath.Base(packagePath)
	printNextStep("WRITE THE NEW " + packageName + " PACKAGE SOURCE CODE.")
	PrintLine(" * The NEW " + packageName + " package will be a file store")
	PrintLine("     of the site/ folder.")
	PrintLine(" * See func serveFileStore in Serve.go.")

	args := []string{"-nu", "-strict", "-o=" + packagePath, filepath.Join(".", "site"), filepath.Join(".", "Http.yaml")}
	var dump []byte
	if dump, err = script.RunDump(nil, appWdPath, "kickpack", args...); err != nil {
		printDump(dump, err)
		return
	}
	printSuccess()
	return
}

// BuildSitePackPackage builds the site pack package.
func BuildSitePackPackage(sitepackPackagePath string) (err error) {
	sitepackPackageName := filepath.Base(sitepackPackagePath)
	printNextStep("BUILD THE NEW " + sitepackPackageName + " PACKAGE.")

	var dump []byte
	if dump, err = script.RunDump(nil, sitepackPackagePath, "go", "build"); err != nil {
		printDump(dump, err)
		return
	}
	printSuccess()
	return
}

// BuildMain builds the application.
// var GOOSFlag string
// var GOARCHFlag string
func BuildMain(rootFolderPath, appName, goos, goarch string) (err error) {
	printNextStep("BUILD THE NEW " + appName + " APPLICATION.")
	env := os.Environ()
	env = append(env, fmt.Sprintf("GOOS=%s", goos))
	env = append(env, fmt.Sprintf("GOARCH=%s", goarch))
	executableName := executabaleName(appName, goos)
	// args := []string{"build", "-o", executableName, filepath.Join(".", "mainprocess/")}
	args := []string{"build", "-o", executableName, filepath.Join(rootFolderPath, "mainprocess")}
	var dump []byte
	if dump, err = script.RunDump(env, rootFolderPath, "go", args...); err != nil {
		fixPrintDump(rootFolderPath, dump, err)
		return
	}
	printSuccess()
	return
}

// RunExecutable runs the framework's executable.
func RunExecutable(rootFolderPath, appName, goos string) (err error) {
	if err = os.Chdir(rootFolderPath); err != nil {
		printError(err)
		return
	}
	ex := executabaleName(appName, goos)
	var dump []byte
	env := os.Environ()
	log.Println("running in: " + rootFolderPath)
	if dump, err = script.RunDump(env, rootFolderPath, filepath.Join(rootFolderPath, ex)); err != nil {
		printDump(dump, err)
		return
	}
	return
}

func executabaleName(appName, goos string) (name string) {
	switch goos {
	case "darwin":
		name = appName
	case "windows":
		name = appName + ".exe"
	default:
		name = appName
	}
	return
}
