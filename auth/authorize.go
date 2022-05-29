package auth

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func CheckCredentials() (*string, *string) {
	var concourseUser string
	var concoursePass string
	var present bool

	concourseUser, present = os.LookupEnv("CONCOURSE_USERNAME")
	if concourseUser == "" || !present {
		fmt.Fprint(os.Stdout, "Environment variable 'CONCOURSE_USERNAME' empty or not set ...")
	} else {
		flag.StringVar(&concourseUser, "username", concourseUser, "Username (Email) for Concourse Labs")
	}

	concoursePass, present = os.LookupEnv("CONCOURSE_PASSWORD")
	if concoursePass == "" || !present {
		fmt.Fprint(os.Stdout, "Environment variable 'CONCOURSE_PASSWORD' empty or not set ...")
	} else {
		flag.StringVar(&concoursePass, "password", concoursePass, "Password for Concourse Labs")
	}

	flag.Parse()
	return &concourseUser, &concoursePass
}

func GetAuthData(user *string, pass *string) string {
	var endpoint string = "https://auth.prod.concourselabs.io/api/v1/oauth/token"
	payload := url.Values{
		"username":   {*user},
		"password":   {*pass},
		"grant_type": {"password"},
		"scope":      {"INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA"},
	}

	// Efficient URL-encoded payload
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // required for form data encoded request
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//Common attributes: resp.Status, resp.Header
	body, err := ioutil.ReadAll(resp.Body) // resp.Body is a map object
	if err != nil {
		log.Fatal(err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Fatalf("JSON parse error: %v", error)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return prettyJSON.String() // string(prettyJSON.Bytes())    // string(body)
	} else {
		return ""
	}
}


/*func main() {
	username, password := checkCredentials()
	fmt.Printf("Concourse Email:    %s\n", *username)
	fmt.Printf("Concourse Password: %s\n", *password)

	respData := getAuthData(username, password)
	fmt.Println(respData)
}
*/
