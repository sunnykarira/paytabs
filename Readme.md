# Paytabs
## Transfer money from one account to another account

A small application to simulate money transfer one account to another.

## Features

- Create Account
- Fetch Account Details
- Send Money from one account to another account
- Add money to account


## Run the application
> Application built on version go version go1.14.4 darwin/amd64
```sh
go mod vendor
go build
./paytabs
```

or

Just run the package exe cross compiled for mac and linux
env GOOS=linux GOARHC=386
```sh
./paytabs_linux
```

env GOOS=darwin GOARHC=386
```sh
./paytabs_mac
```

# API SAMPLE REQUEST AND RESPONSES
The sample script for internal curl is embedded on the server itself. So, it will automatically run to simulate sample scenarios. You can still use the api endpoints exposed to manipulate the data.

> Create Account Request
```curl
    curl --location --request POST 'localhost:8080/api/v1/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": 11,
    "balance": 10,
    "location": "test",
    "account_status": 1

}'
```

> Create Account Response
```json
{
    "code": "",
    "message": "account created successfully",
    "http_code": 200
}
```


-----

> Fetch Account Details
```curl
curl --location --request POST 'localhost:8080/api/v1/fetch' \
--header 'Content-Type: application/json' \
--data-raw '{
    
    "account_id": 10

}'
```


> Fetch Account Details Response
```json
{
    "id": 10,
    "balance": 8,
    "location": "test",
    "account_status": 1
}

OR

{"code":"","message":"account does not exist account id 66"}
```

-----

>Send Money from one account to another account Request

```curl
curl --location --request POST 'localhost:8080/api/v1/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source_account_id": 10,
    "destination_account_id": 11,
    "amount": 1

}'
```

> Send Money from one account to another account Response

```json
{
    "code": "",
    "message": "sent money successfully",
    "http_code": 200
}
```

----

> Add money to account Request

```curl
curl --location --request POST 'localhost:8080/api/v1/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": 10,
    "amount": 1

}'

```

> Add money to account Response

```json
{
    "code": "",
    "message": "added money successfully",
    "http_code": 200
}
```

-----
## All invalid cases respond with appropriate errors
