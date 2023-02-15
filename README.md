# Simple CRUD API

The objective of this project is to train a basic implementation of CRUD operations with golang.

This app allows you to create, read, update and delete shipments, with the help of PostgreSQL database

- Used [Gin framework](https://gin-gonic.com/) as web http framework to handle http requests.
- Used [Gorm](https://gorm.io/) as ORM to handle database connections.

## Requirements

You need to have golang installed on your machine - [Install golang](https://go.dev/doc/install)

You need to create a file in the root of the project as `.env` file, and change the values to your **`PostgresSQL`** db connection

EXAMPLE:

```bash
PORT="3000" # VARIABLE will be automatic identified by GIN framework and will set the HTTP port to the value of this variable

DB_HOST="${CHANGE_ME}"
DB_USER="${CHANGE_ME}"
DB_PASS="${CHANGE_ME}"
DB_NAME="${CHANGE_ME}"
DB_PORT="${CHANGE_ME}"
DB_SSL_MODE="${CHANGE_ME}"
```

## Start app locally

Run this command at the root of the project, to run the app from source code:

```bash
go run main.go
```

## Run app as Docker container

Run this command at the root of the project, to run the app as docker container:

Build docker image:

```bash
docker build -t docker.io/library/golang-crud-gin-gorm .
```

Run docker image:

```bash
docker run -p 8080:3000 docker.io/library/golang-crud-gin-gorm
```

Go to:

```bash
http://localhost:8080/api/...
```

**NOTE:**

To run the app as docker container you need to have docker installed on your machine - [Install docker](https://docs.docker.com/engine/install/ubuntu/)

## Endpoints available

- POST - Create shipment `/api/shipment`
- GET - Get all shipments `/api/shipment`
- GET - Get shipment by id `/api/shipment/:id`
- PUT - Update shipment by id `/api/shipment/:id`
- DELETE - Delete shipment by id `/api/shipment/:id`
