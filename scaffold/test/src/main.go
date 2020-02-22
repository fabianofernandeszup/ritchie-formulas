package main

import (
	"os"
	"test/pkg/application"
)

func main() {
	application.Run(loadInputs())
}

func loadInputs() application.Inputs {
	return application.Inputs{
		Type:        os.Getenv("TYPE"),
		Language:    os.Getenv("LANGUAGE"),
		BootVersion: os.Getenv("BOOT_VERSION"),
		BaseDir:     os.Getenv("BASE_DIR"),
		GroupId:     os.Getenv("GROUP_ID"),
		ArtifactId:  os.Getenv("ARTIFACT_ID"),
		Name:        os.Getenv("NAME"),
		Description: os.Getenv("DESCRIPTION"),
		PackageName: os.Getenv("PACKAGE_NAME"),
		Packaging:   os.Getenv("PACKAGING"),
		JavaVersion: os.Getenv("JAVA_VERSION"),
	}

}
