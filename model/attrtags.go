package model

import (
	"bytes"
	"encoding/json"
	"github.com/saurabh-mish/overcloud/auth"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	ID            int  `json:"id"`
	Version       int  `json:"version"`
	//Created       time `json:"created"`
	//Updated       time `json:"updated"`
	CreatedBy     int  `json:"created_by"`
	UpdatedBy     int  `json:"updated_by"`
	InstitutionId int  `json:"institutionId"`
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
}

/*
func ReadAttributeTag() {



}
*/

/*
func UpdateAttributeTag(tagId int) {
	// read file

	// unmarshall to struct

	// encode to JSON

	// perform request

	// unmarshall response body to AttrTagResp struct and compare 'created' and 'updated' times
}
*/

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

	// convert map object to byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
