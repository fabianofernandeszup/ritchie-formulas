package main

import (
	"jobs/pkg/jobs"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() jobs.Inputs {
	return jobs.Inputs{
		JobName:      os.Getenv("JOB_NAME"),
		JenkinsUser:  os.Getenv("JENKINS_USER"),
		JenkinsToken: os.Getenv("JENKINS_TOKEN"),
	}

}
