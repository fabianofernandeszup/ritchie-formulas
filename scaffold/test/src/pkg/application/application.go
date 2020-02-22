package application

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"test/pkg/fileutil"
)

const (
	starterURL  = "https://start.spring.io/starter.zip"
	jenkinsFile = "Jenkinsfile"
)

type Inputs struct {
	Type        string
	Language    string
	BootVersion string
	BaseDir     string
	GroupId     string
	ArtifactId  string
	Name        string
	Description string
	PackageName string
	Packaging   string
	JavaVersion string
}

func Run(inputs Inputs) {
	log.Println("Starting scaffold generation...")
	fmt.Printf("Name: %v\n", inputs.Name)
	fmt.Printf("Description: %v\n", inputs.Description)

	zipFile, err := downloadFile(inputs)

	if err != nil {
		log.Fatal("Failed to download starter project", err)
	}

	if err := unzipFile(zipFile); err != nil {
		log.Fatal("Failed to Unzip file", err)
	}

	/* if err := inputs.changePermissionJenkinsFile(); err != nil {
	 	log.Fatal("Failed to change permission to Jenkinsfile", err)
	} */

	log.Println("Finished scaffold generation")
}

func downloadFile(inputs Inputs) (string, error) {
	log.Println("Starting download project.")

	req, err := http.NewRequest("GET", starterURL, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("type", inputs.Type)
	q.Add("language", inputs.Language)
	q.Add("bootVersion", inputs.BootVersion)
	q.Add("baseDir", inputs.BaseDir)
	q.Add("artifactId", inputs.ArtifactId)
	q.Add("groupId", inputs.GroupId)
	q.Add("name", inputs.Name)
	q.Add("description", inputs.Description)
	q.Add("packageName", inputs.PackageName)
	q.Add("packaging", inputs.Packaging)
	q.Add("javaVersion", inputs.JavaVersion)
	req.URL.RawQuery = q.Encode()

	log.Println(req.URL)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	log.Println(resp.Status)

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

/* func (i Inputs) changePermissionJenkinsFile() error {
	file := fmt.Sprintf("%s/%s", i.Name, jenkinsFile)
	return fileutil.ChangePermission(file, 0755)
} */
