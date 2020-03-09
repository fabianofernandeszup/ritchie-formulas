package path

import (
	"github.com/fatih/color"
	"ritchie/pkg/file/fileutil"
)

type MainPaths struct {
	RitchieScaffoldTemplate string
	MakeFile                string
	TreeFile                string
}

func IsOnRightDir(mainPaths MainPaths) bool {
	if fileutil.Exists(mainPaths.TreeFile) &&
		fileutil.Exists(mainPaths.MakeFile) &&
		fileutil.Exists(mainPaths.RitchieScaffoldTemplate) {
		return true
	}
	color.Red("Please go to the ritchie-formulas path.")
	return false
}

func BuildMainPaths() MainPaths {
	return MainPaths{
		RitchieScaffoldTemplate: "scaffold/ritchie/src/template",
		MakeFile:                "Makefile",
		TreeFile:                "tree/tree.json",
	}
}
