package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	zupGitUrl  = "https://api.github.com/repos/ZupIT/"
	bodyCi = `{
  "name": "web",
  "active": true,
  "events": [
    "create",
    "pull_request",
    "push"
  ],
  "config": {
    "url": "https://ci.zup.com.br/github-webhook/",
    "content_type": "form",
    "insecure_ssl": "0"
  }
}`
	bodyBot = `{
  "name": "web",
  "active": true,
  "events": [
    "issue_comment"
  ],
  "config": {
    "url": "http://zupbot.herokuapp.com/payload",
    "content_type": "json",
    "insecure_ssl": "0"
  }
}`
)


type Inputs struct {
	Repository string
	GitUser string
	GitToken string
}

type Response struct {
	Status	int
	Body	string
}

func (in Inputs) Run()  {
	log.Println("Zup Web Hook Formula Starter!")
	log.Println("Creating Web Hook CI")
	resp := in.createWebHook(bodyCi)
	log.Printf("Web Hook CI Response code: %d, \nResponse body: \n%v\n", resp.Status, resp.Body)
	log.Println("Creating Web Hook ZupBot")
	resp = in.createWebHook(bodyBot)
	log.Printf("Web Hook ZupBot Response code: %d, \nResponse body: \n%v\n", resp.Status, resp.Body)
	log.Println("Zup Web Hook Formula Finished!!!")
}

func (in Inputs) createWebHook(body string) Response {
	url := fmt.Sprint(zupGitUrl, in.Repository, "/hooks")
	jsonCi := []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonCi))
	if err != nil {
		log.Fatal("Error to create WebHook Request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process WebHook Request: ", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, bodyBytes, "", "\t")
	if error != nil {
		log.Fatal(error)
	}
	return Response{
		Status: resp.StatusCode,
		Body:   string(prettyJSON.Bytes()),
	}
}

