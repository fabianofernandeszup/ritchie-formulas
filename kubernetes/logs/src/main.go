package main

import (
	"kubernetes/logs/pkg/logs"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() logs.Inputs {
	return logs.Inputs{
		Namespace:   os.Getenv("NAMESPACE"),
		PodPartName: os.Getenv("POD_PART_NAME"),
		Kubeconfig:  os.Getenv("KUBECONFIG"),
	}
}