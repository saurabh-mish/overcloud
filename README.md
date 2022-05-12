# overcloud

Concourse Labs endpoint consumption

## [Authorization Service][1]

### Endpoint Test with cURL

+ Install tools

  (PACKAGE_MANAGER: `apt`, `yum`, `brew`, etc.)

  ```
  PACKAGE_MANAGER install curl
  PACKAGE_MANAGER install jq
  ```

+ Auth Service `POST` request

  ```
  curl --request POST 'https://auth.prod.concourselabs.io/api/v1/oauth/token' \
    --header 'Accept: application/json' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'username=user+113@concourselabs.com' \
    --data-urlencode 'password=decentPassword' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'scope=INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA' \
    | jq
  ```

### CLI Application

+ Set environment variables

  ```
  export CONCOURSE_USERNAME="user+113@concourselabs.com"
  export CONCOURSE_PASSWORD="decentPassword"
  export CONCOURSE_ATTRIBUTE_TAGS="192077,154840"
  ```

+ Testing

  ```
  go test ./authorization
  ```

+ Coverage

  + Enable code coverage for tests

    `go test -v ./authorization -coverprofile profile.out`

  + View coverage data from raw file

    `go tool cover -func profile.out`

  + Translate raw data to HTML

    `go tool cover -html=profile.out -o coverage.html`

  + Open HTML file in browser for inspection


[1]: https://api-doc.prod.concourselabs.io/?urls.primaryName=Auth%20Service
[2]: https://api-doc.prod.concourselabs.io/?urls.primaryName=Model%20Service
