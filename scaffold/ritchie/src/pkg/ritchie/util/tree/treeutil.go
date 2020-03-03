package tree

import (
	"encoding/json"
	"github.com/thoas/go-funk"
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/util/error"
	"ritchie/pkg/ritchie/util/input"
	"ritchie/pkg/ritchie/util/path"
	"strings"
)

type Formula struct {
	Path    string
	Bin     string
	Config  *string
	RepoUrl string
}

type Command struct {
	Usage   string
	Help    string
	Formula *Formula
	Parent  string
}

type Tree struct {
	Commands []Command
	Version  string
}

func ChangeTreeFile(input input.Input, mainPaths path.MainPaths) {
	treeFile, err := fileutil.ReadFile(mainPaths.TreeFile)
	error.VerifyError(err)

	var jsonTree Tree
	error.VerifyError(json.Unmarshal(treeFile, &jsonTree))

	jsonTree = generateTreeFile(input, 0, jsonTree)
	jsonResult, _ := json.MarshalIndent(jsonTree, "", "  ")
	error.VerifyError(fileutil.WriteFile("tree/tree.json", jsonResult))
}

func generateTreeFile(input input.Input, i int, treeJson Tree) Tree {
	var parent = generateParent(input, i)
	command := funk.Filter(treeJson.Commands, func(command Command) bool {
		return command.Usage == input.FullName[i] && command.Parent == parent
	}).([]Command)

	if len(input.FullName)-1 == i {
		if len(command) == 0 {
			pathValue := strings.Join(input.FullName, "/")
			commands := append(treeJson.Commands, Command{
				Usage: input.Name,
				Help:  input.Description,
				Formula: &Formula{
					Path:    pathValue,
					Bin:     input.Name + "-${so}",
					Config:  nil,
					RepoUrl: "http://ritchie-cli-bucket152849730126474.s3-website-sa-east-1.amazonaws.com/formulas/" + pathValue,
				},
				Parent: parent,
			})
			treeJson.Commands = commands
			return treeJson
		} else {
			panic("Form already created.")
		}
	} else {
		if len(command) == 0 {
			commands := append(treeJson.Commands, Command{
				Usage:   input.FullName[i],
				Help:    "",
				Formula: nil,
				Parent:  parent,
			})
			treeJson.Commands = commands
		}
	}
	return generateTreeFile(input, i+1, treeJson)
}

func generateParent(input input.Input, index int) string {
	if index > 0 {
		return "root_" + strings.Join(input.FullName[0:index], "_")
	} else {
		return "root"
	}
}
