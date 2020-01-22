package user

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Id string 	`json:"id"`
	Name string `json:"name"`
}

const userUrl = "https://darwin-api.continuousplatform.com/moove/users/{emailBase64}"

func GetUserId(token, username string) User {
	user := base64.StdEncoding.EncodeToString([]byte(username))
	url := strings.ReplaceAll(userUrl, "{emailBase64}", user)
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
		log.Fatalf("Failed to call service get user info. %v\n", resp)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response User
	json.Unmarshal(bodyBytes, &response)
	return response
}
