package makefile

import (
	"ritchie/pkg/file/fileutil"
	"ritchie/pkg/ritchie/util/error"
	"ritchie/pkg/ritchie/util/input"
	"ritchie/pkg/ritchie/util/path"
	"strings"
)

func ChangeMakeFile(inputValue input.Input, mainPaths path.MainPaths) {
	templateFile, err := fileutil.ReadFile(mainPaths.MakeFile)
	error.VerifyError(err)
	variable := strings.ToUpper(inputValue.Name) + "=" + strings.Join(inputValue.FullName, "/")
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
}

func getFormulaValue(file []byte) string {
	fileString := string(file)
	return strings.Split(strings.Split(fileString, "FORMULAS=")[1], "\n")[0]
}
