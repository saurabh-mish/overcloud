# overcloud

Concourse Labs command-line application to evaluate cloud infrastructure.

## cURL Request

+ Postman Export

  ```
  curl --location --request POST 'https://auth.prod.concourselabs.io/api/v1/oauth/token' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'username=user+113@concourselabs.com' \
  --data-urlencode 'password=decentPassword' \
  --data-urlencode 'grant_type=password' \
  --data-urlencode 'scope=INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA'
  ```

+ Modified `POST` request

  ```
  curl --request POST 'https://auth.prod.concourselabs.io/api/v1/oauth/token' \
    --header 'Accept: application/json' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'username=user+113@concourselabs.com' \
    --data-urlencode 'password=decentPassword' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'scope=INSTITUTION POLICY MODEL IDENTITY RUNTIME_DATA'
  ```
