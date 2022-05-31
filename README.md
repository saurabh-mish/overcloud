[![Overcloud](https://github.com/saurabh-mish/overcloud/actions/workflows/ci.yaml/badge.svg)](https://github.com/saurabh-mish/overcloud/actions/workflows/ci.yaml)

# Overcloud

Consumes public APIs of Concourse Labs.

This CLI application consumes endpoints from *Auth* and *Model* services.

## Executing CLI Application

+ Run main

  `go run main.go`

  **NOTE**: Comment / Uncomment functions in this file.

## Testing APIs with HTTP

Authorization access is retrieved from the token endpoint `oauth/token` which is part of the auth service.
This token is then used to interact with the attribute tag endpoint `institutions/<id>/attribute-tags` - part of the model service.


### Getting Started

+ Install tools

  (PACKAGE_MANAGER: `apt`, `yum`, `brew`, etc.)

  ```
  PACKAGE_MANAGER install curl
  PACKAGE_MANAGER install jq
  ```

+ Set environment variables

  ```
  export CONCOURSE_USERNAME="user+113@concourselabs.com"
  export CONCOURSE_PASSWORD="decentPassword"
  ```


### [Authorization Service][1]

+ Auth Service `POST` request

  ```
  curl --request POST 'https://auth.prod.concourselabs.io/api/v1/oauth/token' \
    --header 'Accept: application/json' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'username='"$CONCOURSE_USERNAME"'' \
    --data-urlencode 'password='"$CONCOURSE_PASSWORD"'' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'scope=INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA' \
    | jq
  ```

+ Copy the value of `access_token` to environment variable `CONCOURSE_TOKEN`.

  ```
  export CONCOURSE_TOKEN=ey...
  ```


### [Model Service][2]

+ Read all attribute tags

  ```
  curl --request GET \
    --url https://prod.concourselabs.io/api/model/v1/institutions/113/attribute-tags \
    --header 'Authorization: Bearer '"$CONCOURSE_TOKEN"''
  ```

+ Create attribute tag

  ```
  curl --request POST \
    --url https://prod.concourselabs.io/api/model/v1/institutions/113/attribute-tags \
    --header 'Authorization: Bearer '"$CONCOURSE_TOKEN"'' \
    --header 'Content-Type: application/json' \
    --data '{
      "name": "some name",
      "description": "brief description"
    }'
  ```

+ Read attribute tag

  ```
  curl --request GET \
    --url https://prod.concourselabs.io/api/model/v1/institutions/113/attribute-tags/<id> \
    --header 'Authorization: Bearer '"$CONCOURSE_TOKEN"''
  ```

+ Update attribute tag

  ```
  curl --request PUT \
    --url https://prod.concourselabs.io/api/model/v1/institutions/113/attribute-tags/<id> \
    --header 'Authorization: Bearer '"$CONCOURSE_TOKEN"'' \
    --header 'Content-Type: application/json' \
    --data '{
      "name": "saurabh_updated_name",
      "description": "saurabh_updated_description"
    }'
  ```

+ Delete attribute tag

  ```
  curl --request DELETE \
    --url https://prod.concourselabs.io/api/model/v1/institutions/113/attribute-tags/<id> \
    --header 'Authorization: Bearer '"$CONCOURSE_TOKEN"''
  ```


### Data Structure

When create, read, and update opertions are performed on an attribute tag, response data with the below structure is returned:

```
{
  "id" : integer,
  "version" : integer,
  "created" : time (UTC),
  "updated" : time (UTC),
  "createdBy" : integer,
  "updatedBy" : integer,
  "institutionId" : integer,
  "name" : string,
  "description" : string
}
```

On deleting an attribute tag, a `204` response is returned without any data.

[1]: https://api-doc.prod.concourselabs.io/?urls.primaryName=Auth%20Service
[2]: https://api-doc.prod.concourselabs.io/?urls.primaryName=Model%20Service
