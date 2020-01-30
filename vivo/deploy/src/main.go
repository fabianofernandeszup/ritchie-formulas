package main

import (
	"deploy/pkg/deploy"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() deploy.Inputs {
	return deploy.Inputs{
		Environment:		os.Getenv("ENVIRONMENT"),
		MicroserviceName:  	os.Getenv("MICROSERVICE_NAME"),
		Version:			os.Getenv("VERSION"),
		JenkinsUser:  		os.Getenv("JENKINS_USER"),
		JenkinsToken: 		os.Getenv("JENKINS_TOKEN"),
		EmailUser: 			os.Getenv("EMAIL_USER"),
		EmailToken: 		os.Getenv("EMAIL_TOKEN"),
		GitUser:            os.Getenv("GIT_USER"),
		GitToken:           os.Getenv("GIT_TOKEN"),
	}
}
