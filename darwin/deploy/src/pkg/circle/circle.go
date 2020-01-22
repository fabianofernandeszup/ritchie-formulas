package circle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const circleUrl  = "https://darwin-api.continuousplatform.com/moove/circles?page={page}&size={size}"

type PageModule struct {
	Content []Circle `json:"content"`
	Page int `json:"page"`
	Size int `json:"size"`
	TotalPages int `json:"totalPages"`
	Last bool `json:"last"`
}

type Circle struct {
	Id string `json:"id"`
	Name string `json:"name"`
}


func SearchCircles(token, partName string) []Circle {
	initialPage := 0
	size := 100
	page := Search(token, initialPage, size)
	circles := page.Content
	for page.Last != true {
		initialPage++
		page = Search(token, initialPage, size)
		circles = append(circles, page.Content...)
	}
	var circlesName []Circle
	for _, circle := range circles {
		partName = strings.ToLower(partName)
		circleName := strings.ToLower(circle.Name)
		if strings.Contains(circleName, partName) {
			circlesName = append(circlesName, circle)
		}
	}
	return circlesName
}

func Search(token string, page, size int) PageModule {
	url := strings.ReplaceAll(circleUrl, "{page}", strconv.Itoa(page))
	url = strings.ReplaceAll(url, "{size}", strconv.Itoa(size))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to call service find circle. %v\n", resp)
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
