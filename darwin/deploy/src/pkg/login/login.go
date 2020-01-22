package login

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	authUrl  	= "https://darwin-keycloak.continuousplatform.com/auth/realms/darwin/protocol/openid-connect/token"
	clientId 	= "darwin-client"
	grantType	= "password"
)

type KeycloakResponse struct {
	AccessToken string `json:"access_token"`
}

func Login(user, pass string) string {
	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Set("client_id", clientId)
	data.Set("username", user)
	data.Set("password", pass)
	req, err := http.NewRequest("POST", authUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal("Failed to authentication!")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Failed to authentication request!")
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to auth darwin. %v\n", resp)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response KeycloakResponse
	json.Unmarshal(bodyBytes, &response)
	return response.AccessToken
}
