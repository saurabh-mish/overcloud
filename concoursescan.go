package main

import (
	"flag"
	"fmt"
	"os"
)

func checkCredentials() (*string, *string) {
	var concourseUser string
	var concoursePass string
	var present bool

	concourseUser, present = os.LookupEnv("CONCOURSE_USERNAME")
	if concourseUser == "" || !present  {
		fmt.Println("Environment variable 'CONCOURSE_USERNAME' empty or not set ...")
	} else {
		flag.StringVar(&concourseUser, "username", concourseUser, "Username (Email) for Concourse Labs")	
	}
	
	concoursePass, present = os.LookupEnv("CONCOURSE_PASSWORD")
	if concoursePass == "" || !present {
		fmt.Println("Environment variable 'CONCOURSE_PASSWORD' empty or not set ...")
	} else {
		flag.StringVar(&concoursePass, "password", concoursePass, "Password for Concourse Labs")	
	}
	
	flag.Parse()
	return &concourseUser, &concoursePass
}

func main() {

	username, password := checkCredentials()

	fmt.Printf("Concourse Email:    %s\n", *username)
	fmt.Printf("Concourse Password: %s\n", *password)
}
