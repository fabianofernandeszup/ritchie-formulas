package microservice

import (
	"fmt"
	"io"
	"log"
	"spring-iti/pkg/file/fileutil"
	"net/http"
	"os"
	"strings"
)

const genURL = "https://iti-initializr.itiaws.dev/starter.zip"

type Inputs struct {
	Packaging string
	JavaVersion string
	Language string
	GroupId string
	ArtifactId string
	Version string
	Name string
	Description string
	PackageName string
}

func Run(inputs Inputs) {
	log.Println("Starting scaffold generation...")
	zipFile, err := downloadZipProject(inputs)
	if err != nil {
		log.Fatal("Failed to download starter project", err)
	}
	err = unzipFile(zipFile)
	if err != nil {
		log.Fatal("Failed to Unzip file", err)
	}
	log.Println("Finished scaffold generation")
}

func unzipFile(filename string) error {
	log.Println("Unzip files...")
	destFolder := strings.Replace(filename, ".zip", "",1)
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
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
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
