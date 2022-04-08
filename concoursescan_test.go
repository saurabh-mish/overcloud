package main

import (
	"os"
	"testing"
)

func TestCredentialsSet(t *testing.T) {
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

func TestCredentialsUnset(t *testing.T) {
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
