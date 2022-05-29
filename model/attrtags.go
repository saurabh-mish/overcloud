package model

import (
	"strconv"
	"bytes"
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

type AttrTag struct {
    Name        string `json:"name"`
    Description string `json:"description"`
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


func CreateAttributeTag() {
	const url      = "https://prod.concourselabs.io/api/model/v1"
	const resource = "/institutions/113/attribute-tags"
	endpoint := url + resource

	jsonPayload := &AttrTag{
	    Name:    "saurabh test name",
	    Description: "Saurabh test description",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jsonPayload)
	req, err := http.NewRequest(http.MethodPost, endpoint, payloadBuf)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    log.Println(string(body))
}

/*
func ReadAttributeTag() {

}
*/

func DeleteAttributeTag(tagId int) {
	const url      = "https://prod.concourselabs.io/api/model/v1"
	const resource = "/institutions/113/attribute-tags"
	attrTag        := strconv.Itoa(tagId)
	endpoint       := url + resource + "/" + attrTag

	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

    // convert map object to byte array
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println(string(body))
}
