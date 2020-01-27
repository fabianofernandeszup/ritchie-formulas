package tags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	tagUrl  = "https://darwin-api.continuousplatform.com/moove/modules/components/{componentId}/tags"
	cicleId = "83cfca63-25a5-4626-a92b-a2efc4b7346a"
)

type TagsResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Version  string `json:"name"`
	Artifact string `json:"artifact"`
}

func GetTagsByComponentId(token, componentId, partName string) TagsResponse {
	url := strings.ReplaceAll(tagUrl, "{componentId}", componentId)
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
		log.Fatalf("Failed to call service get tags. %v\n", resp)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response TagsResponse
	json.Unmarshal(bodyBytes, &response)

	var tags []Tag

	for _, t := range response.Tags {
		if strings.Contains(strings.ToLower(t.Version), strings.ToLower(partName)) {
			tags = append(tags, t)
		}
	}
	return TagsResponse{Tags: tags}
}
