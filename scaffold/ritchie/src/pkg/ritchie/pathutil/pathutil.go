package pathutil

import (
	"github.com/fatih/color"
	"ritchie/pkg/file/fileutil"
)

type MainPaths struct {
	RitchieScaffold         string
	RitchieScaffoldTemplate string
	MakeFile                string
	TreeFile                string
}

func RightDir(mainPaths MainPaths) bool {
	if fileutil.Exists(mainPaths.TreeFile) ||
		fileutil.Exists(mainPaths.MakeFile) ||
		fileutil.Exists(mainPaths.RitchieScaffoldTemplate) ||
		fileutil.Exists(mainPaths.RitchieScaffold) {
		return true
	}
	color.Red("Please go to the ritchie-formulas path.")
	return false
}

func BuildMainPaths() MainPaths {
	return MainPaths{
		RitchieScaffold:         "scaffold/ritchie",
		RitchieScaffoldTemplate: "scaffold/ritchie/template",
		MakeFile:                "Makefile",
		TreeFile:                "tree/tree.json",
	}
}
