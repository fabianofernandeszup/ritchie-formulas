package handbook

import (
	"log"
)

const (
	zupGitUrl = "https://api.github.com/repos/ZupIT/{{REPOSITORY}}/git/refs"
)

type Inputs struct {
	RepositoryName string
	GitUser        string
	GitToken       string
}

func (in Inputs) Run() {
	log.Println("Handbook Starter ...")
}
