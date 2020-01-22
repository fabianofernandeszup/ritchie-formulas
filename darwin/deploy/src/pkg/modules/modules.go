package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	moduleUrl  = "https://darwin-api.continuousplatform.com/moove/modules?name={name}&page={page}&size={size}"
	cicleId = "83cfca63-25a5-4626-a92b-a2efc4b7346a"
)

type PageModule struct {
	Content []Module 	`json:"content"`
	Page int 			`json:"page"`
	Size int 			`json:"size"`
	TotalPages int 		`json:"totalPages"`
	Last bool 			`json:"last"`
}

type Module struct {
	Id string              `json:"id"`
	Name string            `json:"name"`
	Components []Component `json:"components"`
}

type Component struct {
	Id string 	`json:"id"`
	Name string `json:"name"`
}

func SearchModules(token, partName string) []Module {
	initialPage := 0
	size := 100
	page := Search(token, partName, initialPage, size)
	modules := page.Content
	for page.Last != true {
		initialPage++
		page = Search(token, partName, initialPage, size)
		modules = append(modules, page.Content...)
	}
	return modules
}

func Search(token, partName string, page, size int) PageModule {
	url := strings.ReplaceAll(moduleUrl, "{page}", strconv.Itoa(page))
	url = strings.ReplaceAll(url, "{size}", strconv.Itoa(size))
	url = strings.ReplaceAll(url, "{name}", partName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
	req.Header.Add("X-Circle-Id", cicleId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to call service search modules. %v\n", resp)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response PageModule
	json.Unmarshal(bodyBytes, &response)
	return response
}
