package main

import (
	"os"

	"produce/pkg/produce"
)

func main() {
	produce.Run(loadInputs())
}

func loadInputs() produce.Inputs {
	return produce.Inputs{
		Urls:  os.Getenv("URLS"),
		Topic: os.Getenv("TOPIC"),
	}
}
