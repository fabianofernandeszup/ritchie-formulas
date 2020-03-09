package input

import (
	"github.com/thoas/go-funk"
	"os"
	"strings"
)

type Input struct {
	Name        string
	FullName    []string
	Description string
}

func BuildInput() Input {
	fullName := buildFullName(os.Getenv("FORMULA_NAME"))
	return Input{
		Name:        fullName[len(fullName)-1],
		FullName:    fullName,
		Description: os.Getenv("FORMULA_DESCRIPTION"),
	}
}

func buildFullName(name string) []string {
	return funk.Filter(strings.Split(name, " "), func(input string) bool {
		return input != "" && input != "rit"
	}).([]string)
}
