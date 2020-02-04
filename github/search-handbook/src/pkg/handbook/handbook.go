package handbook

import (
	"encoding/base64"
	"encoding/json"
	"handbook/pkg/prompt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	zupGitListUrl = "https://api.github.com/search/code?q={{WORD}}+in:file+repo:zupit/{{REPOSITORY}}"
	//Search code = https://api.github.com/search/code?q=main+in:file+repo:zupit/ritchie-cli
)

type Archive struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Content string `json:"content"`
}

type Inputs struct {
	RepositoryName string
	GitUser        string
	GitToken       string
}

func (in Inputs) Run() {
	log.Println("Handbook Search Code Starter ...")

	repository := readRepository()
	word := readWord()

	url := strings.ReplaceAll(zupGitListUrl, "{{REPOSITORY}}", repository)
	url = strings.ReplaceAll(url,"{{WORD}}",word)

	log.Println(url)

	archives := in.searchRepository(url)
	log.Println(archives)
}

func decodeContent(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("error:", err)
		return ""
	}
	return string(data)
}

func archivesToString(archives []Archive) []string {
	var str []string

	for _, a := range archives {
		str = append(str, a.Name)
	}
	return str
}

func readRepository() string {
	repository, err := prompt.String("Type name of application repository: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}

func readWord() string {
	repository, err := prompt.String("Type word of search: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}

func (in Inputs) searchRepository(url string) []Archive {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error to scan Repository Request: ", err)
	}
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process scan Repository: ", err)
	}
	defer resp.Body.Close()

	var archives []Archive
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &archives)
	if err != nil {
		log.Fatal("Error proccess convert json to struct:", err)
	}
	return archives
}

func verifyTypeFile(archives []Archive, str string) bool{

	for _,a := range archives  {
		if a.Name == str {
			switch a.Type {
			case "file":
				return true
			case "dir":
				return false
			default:
				log.Fatal("Type GitHub is not valid.")
			}
		}
	}
	return false
}