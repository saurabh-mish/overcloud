+ Testing

  ```
  go test ./auth
  go test ./model
  ```

+ Coverage

  + Enable code coverage for tests

    `go test -v ./auth -coverprofile profile.out`

  + View coverage data from raw file

    `go tool cover -func profile.out`

  + Translate raw data to HTML

    `go tool cover -html=profile.out -o coverage.html`

  + Open HTML file in browser for inspection
