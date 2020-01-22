package deployment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	deployUrl = "https://darwin-api.continuousplatform.com/moove/deployments"
	cicleId = "83cfca63-25a5-4626-a92b-a2efc4b7346a"
)

type Deployment struct {
	AuthorId string `json:"authorId"`
	CircleId string	`json:"circleId"`
	BuildId string	`json:"buildId"`
}

type Response struct {
	Id string `json:"id"`
	Status string	`json:"status"`
}

func GetStatus(token, id string) Response {
	url := fmt.Sprint(deployUrl, "/", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Error to get status Deploy Request: ", err)
	}
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
	req.Header.Add("X-Circle-Id", cicleId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process to get status Deploy Request: ", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatal("Failed to call service status deploy.")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Response
	json.Unmarshal(bodyBytes, &response)
	return response
}

func CreateDeployment(d Deployment, token string) Response {
	jsonValue, _ := json.Marshal(d)
	req, err := http.NewRequest(http.MethodPost, deployUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal("Error to create Deploy Request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
	req.Header.Add("X-Circle-Id", cicleId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process Deploy Request: ", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to call service create deploy. %v\n", resp)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Response
	json.Unmarshal(bodyBytes, &response)
	return response
}
