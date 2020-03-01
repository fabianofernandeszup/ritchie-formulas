package main

import (
	"github.com/fatih/color"
	"os"
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/pathutil"
)

func main() {

	name := os.Getenv("FORMULA_NAME")
	description := os.Getenv("FORMULA_DESCRIPTION")
	mainPaths := pathutil.BuildMainPaths()
	if !pathutil.RightDir(mainPaths) {
		return
	}

	fileutil.CreateIfNotExists(name)
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-config.json")
	if err != nil {
		panic(err)
	}
	fileutil.WriteFile(name+"/config.json", templateFile)
	color.Green("Generate formula:" + name + " with description:" + description + " .")

}
