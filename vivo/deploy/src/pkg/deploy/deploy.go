package deploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

const (
	zupGitUrl  = "https://api.github.com/repos/ZupIT/{{MICROSERVICE}}/git/refs"
	zupGitTagsUrl = "https://api.github.com/repos/ZupIT/{{MICROSERVICE}}/tags"
	zupJenkinsURL = "https://ci.zup.com.br/job/deploy-vivo-easy-dev-k8s/buildWithParameters?MICROSERVICE={{MICROSERVICE}}&VERSION={{VERSION}}" //TODO:: SWITCH ENVIROMENT TO PRD
	bodyBrancheGit = `{
    "ref": "refs/heads/{{NEWBRANCH}}",
    "sha": "{{SHA}}"
	}`
	addressSMTP = "smtp.gmail.com:587"
	hostSMTP = "smtp.gmail.com"
)

type Inputs struct {
	MicroserviceName	string
	Version 			string
	JenkinsUser			string
	JenkinsToken 		string
	EmailUser 			string
	EmailToken 			string
	GitUser				string
	GitToken			string
}

type RespGet struct {
	Ref	string `json:"ref"`
	Object Object `json:"object"`
}
type Object struct {
	Sha string `json:"sha"`
}
type Response struct {
	Status int
	Body   string
}

func (in Inputs) Run() {
	log.Println("Vivo Easy Deploy Starter...")

	urlGit := strings.ReplaceAll(zupGitUrl, "{{MICROSERVICE}}", in.MicroserviceName)
	in.createBranch(urlGit)
	urlJenkins := strings.ReplaceAll(zupJenkinsURL,"{{MICROSERVICE}}",in.MicroserviceName)
	urlJenkins = strings.ReplaceAll(urlJenkins,"{{VERSION}}",in.Version)
	resp := in.executeJobJenkins(urlJenkins)

	if resp.Status == 201 {
		log.Println("Pipeline running ...")
		urlApprove := "https://ci.zup.com.br/job/deploy-vivo-easy-dev-k8s"
		log.Println("Please approve changes to: ")
		log.Println(urlApprove)
		in.sendMail("Change executada com sucesso...\n Microserviço: "+in.MicroserviceName+" Versão: "+in.Version)
	}else {
		log.Println("Error to run pipeline in Jenkins...")
	}

}
func (in Inputs) executeJobJenkins(urlJenkins string) Response {

	urlGet := strings.ReplaceAll(zupGitTagsUrl,"{{MICROSERVICE}}",in.MicroserviceName)

	b := true
	for b {
		i := 1
		b = in.verifyVersionCreated(urlGet)
		log.Println("Waiting for generate release.")
		time.Sleep(time.Minute)
		i++
		if i > 10{
			b = false
			log.Fatal("Timeout with Jenkins, please verify...")
		}
	}

	urlZupJenkins := strings.ReplaceAll(zupJenkinsURL,"{{MICROSERVICE}}",in.MicroserviceName)
	urlZupJenkins = strings.ReplaceAll(urlZupJenkins,"{{VERSION}}",in.Version)

	req, err := http.NewRequest("POST", urlZupJenkins, nil)
	if err != nil {
		log.Fatal("Error to execute Jenkins Pipeline Request: ", err)
	}
	req.SetBasicAuth(in.JenkinsUser, in.JenkinsToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error request Jenkins execute pipeline: ", err)
	}
	defer resp.Body.Close()

	return Response{
		Status: resp.StatusCode,
		Body:   "",
	}
}
func (in Inputs) verifyVersionCreated(url string) bool{

	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		log.Fatal("Error to GET verify Tag in Git Request: ", err)
	}
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process verify release Request: ", err)
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
    return !strings.Contains(string(prettyJSON.Bytes()), in.Version)
}
func (in Inputs) createBranch(url string) Response{
	//URL+/heads/master
	urlGet := fmt.Sprint(url,"/heads/master")
	req, err := http.NewRequest("GET", urlGet,nil)
	if err != nil {
		log.Fatal("Error to GET create branch release Request: ", err)
	}
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process create branch Request: ", err)
	}
	defer resp.Body.Close()
	respGet := RespGet{}

	err = json.NewDecoder(resp.Body).Decode(&respGet)
	if err != nil {
		log.Fatal(err)
	}
	if in.Version == ""{
		log.Fatal("Error Version is empty...")
	}
	body := strings.ReplaceAll(bodyBrancheGit, "{{NEWBRANCH}}", "release-"+in.Version)
	body = strings.ReplaceAll(body,"{{SHA}}",respGet.Object.Sha)

	jsonNewBranch := []byte(body)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonNewBranch))
	if err != nil {
		log.Fatal("Error to POST create branch release Request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process POST create branch release Request: ", err)
	}
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

func (in Inputs) sendMail(body string) {
	to := []string{"thiago.oliveira@zup.com.br","gabriel.pinheiro@zup.com.br", "barbara.rocha@zup.com.br",
		           "rodrigo.pereira@zup.com.br","petterson.santos@zup.com.br","juliano.borges@zup.com.br",
				   "nicolas.peixoto@zup.com.br"}

	msg := "From: " + in.EmailUser + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: [Ritchie]Change realizada com sucesso!\n\n" +
		body

	err := smtp.SendMail(addressSMTP,
		smtp.PlainAuth("", in.EmailUser, in.EmailToken, hostSMTP),
		in.EmailUser, to, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Email enviado para: ")
	for _,v := range to{
		log.Println(v)
	}
}