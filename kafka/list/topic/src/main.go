package main

import (
	"os"

	"topic/pkg/topic"
)

func main() {
	topic.List(loadInputs())
}

func loadInputs() topic.Inputs {
	u := os.Getenv("URL")
	return topic.Inputs{Url: u}
}
