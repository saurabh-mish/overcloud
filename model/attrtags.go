package model

import (
	"bytes"
	"encoding/json"
	"github.com/saurabh-mish/overcloud/auth"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ConcourseAuthData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int16  `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
	// ignoring JSON field "extra_info"
}

type AttrTagReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AttrTagResp struct {
	ID            int    `json:"id"`
	Version       int    `json:"version"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	CreatedBy     int    `json:"createdBy"`
	UpdatedBy     int    `json:"updatedBy"`
	InstitutionId int    `json:"institutionId"`
	// ignoring JSON fields "name" and "description"
}

const url      = "https://prod.concourselabs.io/api/model/v1"
const resource = "/institutions/113/attribute-tags"


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

/*
{
	"id": 212891,
	"version": 0,
	"created": "2022-05-29T20:18:50.190Z",
	"updated": "2022-05-29T20:18:50.190Z",
	"createdBy": 101685,
	"updatedBy": 101685,
	"institutionId": 113,
	"name": "saurabh_test_name",
	"description": "saurabh_test_description"
}
*/

func CreateAttributeTag() {
	endpoint := url + resource

	jsonPayload := &AttrTagReq{
		Name:        "saurabh test name",
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

	// unmarshall response body to AttrTagResp struct which will output the ID of attribute tag created

	var jsonData AttrTagResp
	json.Unmarshal(body, &jsonData)
	createdTime, err := time.Parse(time.RFC3339, jsonData.Created)
	updatedTime, err := time.Parse(time.RFC3339, jsonData.Updated)
	if err != nil {
	    log.Printf("Unable to parse given string to time: %v\n", err)
	}

	if createdTime == updatedTime {
		log.Printf("\nAttribute tag with ID %v created successfully!", jsonData.ID)
	} else {
		log.Println("\nError with time difference ...")
	}
}


func ReadAttributeTag(tagId int) {
	attrTag := strconv.Itoa(tagId)
	endpoint := url + resource + "/" + attrTag

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	// convert map object to byte array
	var jsonData AttrTagResp
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &jsonData)

	log.Println(jsonData)
}


func UpdateAttributeTag(tagId int) {
	attrTag := strconv.Itoa(tagId)
	endpoint := url + resource + "/" + attrTag
	log.Println(endpoint)

	jsonPayload := &AttrTagReq{
		Name:        "saurabh updated name",
		Description: "Saurabh updated description",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(jsonPayload)
	req, err := http.NewRequest(http.MethodPut, endpoint, payloadBuf)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData AttrTagResp
	json.Unmarshal(body, &jsonData)
	log.Println(jsonData)
	createdTime, cerr := time.Parse(time.RFC3339, jsonData.Created)
	if err != nil {
	    log.Printf("Unable to parse given string to time: %v\n", cerr)
	}
	updatedTime, uerr := time.Parse(time.RFC3339, jsonData.Updated)
	if err != nil {
	    log.Printf("Unable to parse given string to time: %v\n", uerr)
	}

	if updatedTime.Sub(createdTime) > 0 {
		log.Printf("\nAttribute tag with ID %v updated successfully!", jsonData.ID)
	} else {
		log.Printf("\nError updating attribute tag; time difference %v ...", updatedTime.Sub(createdTime))
	}

}


func DeleteAttributeTag(tagId int) {
	attrTag := strconv.Itoa(tagId)
	endpoint := url + resource + "/" + attrTag

	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		log.Println("Endpoint unavailable ...")
	}

	apiToken := "Bearer " + getAccessToken()
	req.Header.Add("Authorization", apiToken)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("Attribute tag deleted successfully!")
	} else {
		log.Printf("Deleting attribute tag failed with:\n%v", resp.StatusCode)
	}
}
