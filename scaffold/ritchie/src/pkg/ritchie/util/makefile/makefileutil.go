package makefile

import (
	"github.com/thoas/go-funk"
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/util/error"
	"ritchie/pkg/ritchie/util/input"
	"ritchie/pkg/ritchie/util/path"
	"strings"
)

func ChangeMakeFile(inputValue input.Input, mainPaths path.MainPaths) string {
	templateFile, err := fileutil.ReadFile(mainPaths.MakeFile)
	error.VerifyError(err)
	variableName := generateVariableName(inputValue)
	variable := variableName + "=" + strings.Join(inputValue.FullName, "/")
	templateFile = []byte(
		strings.ReplaceAll(
			string(templateFile),
			"\nFORMULAS=",
			"\n"+variable+"\nFORMULAS=",
		),
	)
	formulas := getFormulaValue(templateFile)

	templateFile = []byte(
		strings.ReplaceAll(
			string(templateFile),
			formulas,
			formulas+" $("+strings.ToUpper(inputValue.Name)+")",
		),
	)

	error.VerifyError(fileutil.WriteFile(mainPaths.MakeFile, templateFile))
	return variableName
}

func generateVariableName(inputValue input.Input) string {
	return strings.Join(toUpper(inputValue.FullName), "_")
}

func toUpper(fullName []string) []string {
	return funk.Map(fullName, func(name string) string {
		return strings.ToUpper(name)
	}).([]string)
}

func getFormulaValue(file []byte) string {
	fileString := string(file)
	return strings.Split(strings.Split(fileString, "FORMULAS=")[1], "\n")[0]
}
