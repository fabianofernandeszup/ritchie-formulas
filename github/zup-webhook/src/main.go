package main

import (
	"os"
	"webhook/pkg/webhook"
)

func main() {
	loadInputs().Run()
}

func loadInputs() webhook.Inputs {
	return webhook.Inputs{
		Repository:			os.Getenv("REPOSITORY"),
		GitUser:            os.Getenv("GIT_USER"),
		GitToken:           os.Getenv("GIT_TOKEN"),
	}
}
