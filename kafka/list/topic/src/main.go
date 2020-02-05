package main

import (
	"os"

	"topic/pkg/topic"
)

func main() {
	topic.List(loadInputs())
}

func loadInputs() topic.Inputs {
	u := os.Getenv("URLS")
	return topic.Inputs{Urls: u}
}
