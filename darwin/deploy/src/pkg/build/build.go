package build

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	buildUrl = "https://darwin-api.continuousplatform.com/moove/builds/compose"
	cicleId = "83cfca63-25a5-4626-a92b-a2efc4b7346a"
)


type Build struct {
	ReleaseName string 	`json:"releaseName"`
	AuthorId string 	`json:"authorId"`
	Modules []Module 	`json:"modules"`
}

type Module struct {
	Id string 				`json:"id"`
	Components []Component 	`json:"components"`
}

type Component struct {
	Id string		`json:"id"`
	Version string	`json:"version"`
	Artifact string	`json:"artifact"`
}

type Response struct {
	Id string		`json:"id"`
}


func CreateBuild(b Build, token string) Response {
	jsonValue, _ := json.Marshal(b)
	req, err := http.NewRequest("POST", buildUrl, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal("Error to create Build Request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
	req.Header.Add("X-Circle-Id", cicleId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process Build Request: ", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to call service to create build. %v\n", resp)
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
