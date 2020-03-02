package main

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/thoas/go-funk"
	"os"
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/pathutil"
	"ritchie/pkg/ritchie/tree"
	"strings"
)

func main() {

	name := os.Getenv("FORMULA_NAME")
	description := os.Getenv("FORMULA_DESCRIPTION")
	mainPaths := pathutil.BuildMainPaths()
	if !pathutil.RightDir(mainPaths) {
		return
	}

	nameList := splitName(name)
	generateFiles(nameList, mainPaths, 0)

	treeFile, err := fileutil.ReadFile(mainPaths.TreeFile)
	verifyError(err)
	var jsonTree tree.Tree
	verifyError(json.Unmarshal(treeFile, &jsonTree))

	commands := funk.Filter(jsonTree.Commands, func(command tree.Command) bool {
		return command.Parent == "root_github"
	})

	jsonTree.Commands = commands.([]tree.Command)

	jsonResult, _ := json.MarshalIndent(jsonTree, "", "  ")
	verifyError(fileutil.WriteFile("tree/tree2.json", jsonResult))

	//nameList := strings.Split(name, " ")

	color.Green("Generate formula:" + name + " with description:" + description + " .")

}

func splitName(name string) []string {
	return funk.Filter(strings.Split(name, " "), func(input string) {
		input = ""
	}).([]string)
}

func generateFiles(nameList []string, mainPaths pathutil.MainPaths, i int) {
	dir := strings.Join(nameList[0:i], "/")
	verifyError(fileutil.CreateIfNotExists(dir))
	if len(nameList)-1 == i {
		createConfigFile(dir, mainPaths)
	} else {
		generateFiles(nameList, mainPaths, i+1)
	}
}

func createConfigFile(dir string, mainPaths pathutil.MainPaths) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-config.json")
	templateFile = []byte(strings.ReplaceAll(string(templateFile), "{{description}}", ""))
	verifyError(err)
	err = fileutil.WriteFile(dir+"/config.json", templateFile)
	verifyError(err)
}

func verifyError(err error) {
	if err != nil {
		panic(err)
	}
}
