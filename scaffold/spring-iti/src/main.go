package main

import (
	"spring-iti/pkg/microservice"
	"os"
)

func main() {
	microservice.Run(loadInputs())
}

func loadInputs() microservice.Inputs {
	return microservice.Inputs{
		Packaging:   os.Getenv("PACKAGING"),
		JavaVersion: os.Getenv("JAVA_VERSION"),
		Language:    os.Getenv("LANGUAGE"),
		GroupId:     os.Getenv("GROUP_ID"),
		ArtifactId:  os.Getenv("ARTIFACT_ID"),
		Version:     os.Getenv("VERSION"),
		Name:        os.Getenv("NAME"),
		Description: os.Getenv("DESCRIPTION"),
		PackageName: os.Getenv("PACKAGE_NAME"),
	}
}
