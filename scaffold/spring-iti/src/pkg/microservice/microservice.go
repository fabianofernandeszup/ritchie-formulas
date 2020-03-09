package microservice

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"spring-iti/pkg/file/fileutil"
	"strings"
)

const (
	genURL      = "https://iti-initializr.itiaws.dev/starter.zip"
	jenkinsFile = "Jenkinsfile"
)

type Inputs struct {
	Packaging    string
	JavaVersion  string
	Language     string
	GroupId      string
	ArtifactId   string
	Version      string
	Name         string
	Description  string
	PackageName  string
	Dependencies string
}

func Run(inputs Inputs) {
	log.Println("Starting scaffold generation...")
	zipFile, err := downloadZipProject(inputs)
	if err != nil {
		log.Fatal("Failed to download starter project\n", err)
	}
	err = unzipFile(zipFile)
	if err != nil {
		log.Fatal("Failed to Unzip file", err)
	}
	err = inputs.changePermissionJenkinsFile()
	if err != nil {
		log.Fatal("Failed to change permission to Jenkinsfile", err)
	}
	log.Println("Finished scaffold generation")
}

func (i Inputs) changePermissionJenkinsFile() error {
	file := fmt.Sprintf("%s/%s", i.Name, jenkinsFile)
	return fileutil.ChangePermission(file, 0755)
}

func unzipFile(filename string) error {
	log.Println("Unzip files...")
	destFolder := strings.Replace(filename, ".zip", "", 1)
	fileutil.CreateIfNotExists(destFolder, 0755)
	err := fileutil.Unzip(filename, destFolder)
	if err != nil {
		return err
	}
	err = fileutil.RemoveFile(filename)
	if err != nil {
		return err
	}
	log.Println("Unzip done.")
	return nil
}

func downloadZipProject(inputs Inputs) (string, error) {
	log.Println("Starting download project.")
	req, err := http.NewRequest("GET", genURL, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("packaging", inputs.Packaging)
	q.Add("javaVersion", inputs.JavaVersion)
	q.Add("language", inputs.Language)
	q.Add("groupId", inputs.GroupId)
	q.Add("artifactId", inputs.ArtifactId)
	q.Add("version", inputs.Version)
	q.Add("name", inputs.Name)
	q.Add("description", inputs.Description)
	q.Add("packageName", inputs.PackageName)
	q.Add("dependencies", inputs.Dependencies)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println(req.URL)
		err := fmt.Errorf("Invalid parameters ou dependencies! Response Status Code: %s", resp.Status)
		return "", err
	}
	prjfile := fmt.Sprintf("%s.zip", inputs.Name)
	out, err := os.Create(prjfile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("Download done.")
	return prjfile, nil
}
