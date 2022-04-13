package main

import (
	"os"
	"testing"
	"encoding/json"
)


type ConcourseAuth struct {
	AccessToken  string     `json:"access_token"`
	TokenType    string     `json:"token_type"`
	RefreshToken string     `json:"refresh_token"`
	ExpiresIn    int16      `json:"expires_in"`
	Scope        string     `json:"scope"`
	Extra        Extra_info `json:"extra"`
	Jti          string     `json:"jti"`
}

type Extra_info struct {
	InstitutionID int
	UserID int
	UserEmail string
	GroupIDs []int
	SurfaceIDs []int
}


func TestCredentialsPresent(t *testing.T) {
	var testuser string = "user+113@concourselabs.com"
	var testpass string = "somepassword"
	os.Setenv("CONCOURSE_USERNAME", testuser)
	os.Setenv("CONCOURSE_PASSWORD", testpass)
	user, pass := checkCredentials()

	t.Run("checking set username", func(t *testing.T) {
		if *user != testuser {
			t.Errorf("Unable to retrieve 'CONCOURSE_USERNAME' environment variable %s", testuser)
		}
	})

	t.Run("checking set password", func(t *testing.T) {
		if *pass != testpass {
			t.Errorf("Unable to retrieve 'CONCOURSE_PASSWORD' environment variable %s", testpass)
		}
	})

	os.Clearenv()
}


func TestCredentialsAbsent(t *testing.T) {
	os.Unsetenv("CONCOURSE_USERNAME")
	os.Unsetenv("CONCOURSE_PASSWORD")
	user, pass := checkCredentials()

	t.Run("checking unset username", func(t *testing.T) {
		if *user != "" {
			t.Errorf("Incorrect value for 'CONCOURSE_USERNAME' being used: %v", user)
		}
	})

	t.Run("checking unset password", func(t *testing.T) {
		if *pass != "" {
			t.Errorf("Incorrect value for 'CONCOURSE_PASSWORD' being used: %v", pass)
		}
	})
}


func TestValidResponseData(t *testing.T) {
	var testuser string = "saurabh+113@concourselabs.com"  
	var testpass string = "S@ura8hM2906"
	responseData := getAuthData(&testuser, &testpass)
	
	t.Run("checking valid response data", func(t *testing.T) {
		if responseData == "" {
			t.Errorf("Error with response data: %v", responseData)
		}	
	})

	//unmarshall JSON response data to struct
	var concourseRespData ConcourseAuth
    jsonData := json.Unmarshal([]byte(responseData), &concourseRespData)

    t.Run("checking JSON data structure", func(t *testing.T) {
    	if jsonData != nil {
    	    t.Errorf("Error parsing JSON data %v", jsonData)
    	}
    })

    t.Run("checking access token exists", func(t *testing.T) {
    	if concourseRespData.AccessToken == "" {
    		t.Errorf("Error retrieving access token %v", concourseRespData.AccessToken)
    	}
    })
	
	// use json.Unmarshall instead of json.Decode - https://stackoverflow.com/a/31129967/13055097
}
