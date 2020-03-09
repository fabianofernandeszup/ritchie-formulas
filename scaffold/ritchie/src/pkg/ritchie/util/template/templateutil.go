package template

import (
	"github.com/fatih/color"
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/util/error"
	"ritchie/pkg/ritchie/util/input"
	"ritchie/pkg/ritchie/util/path"
	"strings"
)

func GenerateFiles(input input.Input, paths path.MainPaths, index int) {
	dir := strings.Join(input.FullName[0:index+1], "/")
	color.Green("create dir:" + dir)
	error.VerifyError(fileutil.CreateIfNotExists(dir))
	if len(input.FullName)-1 == index {
		createConfigFile(dir, paths, input)
		createSrcDir(dir, paths, input)
	} else {
		GenerateFiles(input, paths, index+1)
	}
}

func createSrcDir(dir string, mainPaths path.MainPaths, inputValue input.Input) {
	srdDir := dir + "/src"
	error.VerifyError(fileutil.CreateIfNotExists(srdDir))
	createMainFile(srdDir, mainPaths)
	createGoModFile(srdDir, mainPaths)
	createMakeFile(srdDir, mainPaths, inputValue.Name)
	error.VerifyError(fileutil.CreateIfNotExists(srdDir + "/pkg/hello"))
	createHelloFile(srdDir, mainPaths)
}

func createGoModFile(dir string, mainPaths path.MainPaths) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-go.mod")
	error.VerifyError(err)
	error.VerifyError(fileutil.WriteFile(dir+"/go.mod", templateFile))
}

func createHelloFile(dir string, mainPaths path.MainPaths) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-hello.txt")
	error.VerifyError(err)
	error.VerifyError(fileutil.WriteFile(dir+"/pkg/hello/hello.go", templateFile))
}

func createMainFile(dir string, mainPaths path.MainPaths) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-main.txt")
	error.VerifyError(err)
	error.VerifyError(fileutil.WriteFile(dir+"/main.go", templateFile))
}

func createConfigFile(dir string, mainPaths path.MainPaths, inputValue input.Input) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-config.json")
	error.VerifyError(err)
	templateFile = []byte(
		strings.ReplaceAll(
			string(templateFile),
			"{{description}}",
			inputValue.Name,
		),
	)
	error.VerifyError(fileutil.WriteFile(dir+"/config.json", templateFile))
}

func createMakeFile(dir string, mainPaths path.MainPaths, name string) {
	templateFile, err := fileutil.ReadFile(mainPaths.RitchieScaffoldTemplate + "/template-Makefile")
	error.VerifyError(err)
	templateFile = []byte(
		strings.ReplaceAll(
			string(templateFile),
			"{{name}}",
			name,
		),
	)
	error.VerifyError(fileutil.WriteFile(dir+"/Makefile", templateFile))
}
