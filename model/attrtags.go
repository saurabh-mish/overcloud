package model

import (
	//"os"
	//"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/saurabh-mish/overcloud/auth"
)

type ConcourseAuthData struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"token_type"`
	RefreshToken string     `json:"refresh_token"`
	ExpiresIn    int16      `json:"expires_in"`
	Scope        string     `json:"scope"`
	Jti          string     `json:"jti"`
}

func getAccessToken() string {
	var concourseAuth ConcourseAuthData

	username, password := auth.CheckCredentials()
	respData := auth.GetAuthData(username, password)

	jsonData := json.Unmarshal([]byte(respData), &concourseAuth)
	if jsonData != nil {
		log.Println("Error parsing JSON data %v", jsonData)
	}

	return concourseAuth.AccessToken
}


func GetAllAttributeTags() {
	const url      = "https://prod.concourselabs.io/api/model/v1"
	const resource = "/institutions/113/attribute-tags"
	endpoint := url + resource

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	// add bearer token to header and send request
	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	resp, _ := http.DefaultClient.Do(req)

    defer resp.Body.Close()

    // convert map object to byte array
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println(string(body))
}
