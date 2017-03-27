# gomicro

 [![Travis](https://img.shields.io/travis/petrasphere/gomicro.svg?style=flat-square)](https://travis-ci.org/petrasphere/gomicro) [![Coveralls](https://img.shields.io/coveralls/petrasphere/gomicro.svg?style=flat-square)]()  [![npm](https://img.shields.io/npm/dm/generator-gomicro.svg?style=flat-square)](https://www.npmjs.com/package/generator-gomicro)

> a yeoman generator for golang rest microservices

> project backlog [here](https://github.com/petrasphere/gomicro/projects/1)

## Quick Start

```sh
$ npm install -g generator-gomicro
$ yo gomicro
```

## Overview

gomicro aims to provide a production grade generator for applications that implement a restful api through create/retrieve/update/delete (CRUD) operations. Below are the aspects that gomicro implements to make your application production ready:

## Table of Contents
* [Quick Start](#quick-start)
* [Overview](#overview)
* [Components](#components)
  * [HTTP Router](#http-router)
  * [Command Line Interface](#command-line-interface)
  * [Database Driver](#database-driver)
  * [API Documentation](#api-documentation)
  * [Unit Testing](#unit-testing)
  * [HTTP Access Control (CORS)](#http-access-control-cors)
  * [Logging](#logging)
  * [Authorization](#authorization)
  * [SSL](#ssl)
  * [Docker](#docker)
  * [Build Scripts](#build-scripts)
  * [Local Development](#local-development)
  * [Deployment](#deployment)
  * [Error Handling](#error-handling)
  * [Health Checks](#health-checks)

## Components

### HTTP Router

> [github.com/gorilla/mux](https://github.com/gorilla/mux)

Start the server with `./gomicro start` (replace "gomicro" with the name of your binary). The server implements CRUD operations on your resource. Reference the api documentation for full api usage. Drag and Drop the generated `swagger.json` onto the [swagger editor](http://editor.swagger.io/#!/)

### Command Line Interface

> [github.com/spf13/cobra](https://github.com/spf13/cobra)

A robust command line tool is included in your project. Available cli flags can are documented throughout this README. However, you can always find out more via the command line tool itself (replace "gomicro" with the name of your binary).

```sh
$ ./gomicro --help
gomicro is a RESTful microservice that performs CRUD operations on the Phone resources

Usage:
  gomicro [command]

Available Commands:
  help        Help about any command
  start       start a new server
  test        run all unit tests
  version     print the version information

Use "gomicro [command] --help" for more information about a command.
```

### Database Driver

> [github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)

The generator provides support for multiple backend database drivers configurable by command line flags.

#### MySQL

```sh
--db-port=DATABASE-PORT or DB_PORT=DATABASE-PORT
--db-host=DATABASE-HOST or DB_HOST=DATABASE-HOST  
--db-user=DATABASE-USER or DB_USER=DATABASE-USER  
--db-pass=DATABASE-PASS or DB_PASS=DATABASE-PASS  
--db-name=DATABASE-NAME or DB_NAME=DATABASE-NAME  
```

#### Sqlite

```sh
--db-location=DATABASE-LOCATION or DB_PORT=DATABASE-LOCATION
--db-name=DATABASE-NAME or DB_NAME=DATABASE-NAME
```

#### PostgreSQL

```sh
--db-port=DATABASE-PORT or DB_PORT=DATABASE-PORT  
--db-host=DATABASE-HOST or DB_HOST=DATABASE-HOST  
--db-user=DATABASE-USER or DB_USER=DATABASE-USER  
--db-pass=DATABASE-PASS or DB_PASS=DATABASE-PASS  
--db-name=DATABASE-NAME or DB_NAME=DATABASE-NAME  
```

##### MongoDB

*coming soon*

### API Documentation

[Swagger](http://swagger.io/specification/) is the industry standard for documenting APIs. Because of this, you will find complete API documentation for your application in your project. Note that you will only have to modify the data model to match yours.

### Unit Testing

*tack this backlog item [here](https://github.com/petrasphere/gomicro/projects/1#card-2228159)*

No project is complete with comprehensive unit testing. The goal is to have over 90% code coverage. You can run these unit tests by executing the following command (replace "gomicro" with the name of your binary):

```sh
./gomicro test
```

### HTTP Access Control (CORS)

[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS) give you the ability to secure cross-origin HTTP requests. For example, the following code allows the `Content-Type` header, allows requests from all origin domains, and only allows four HTTP methods:

```go
headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
originsOk := handlers.AllowedOrigins([]string{"*"})
methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
```

### Logging

> [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

Robust logging is streamed to stdout with the following json format:

```js
{"duration":0,"level":"info","method":"GET","msg":"http server started","name":"RetrievePhone","time":"2017-03-26T22:55:25-05:00","uri":"/phone/6"}
```

### Authorization

#### Static Password File

Basic authentication is enabled by passing the `--basic-auth-file=SOMEFILE` option to the server. Currently, the basic auth credentials last indefinitely, and the password(s) cannot be changed without restarting the server.

The basic auth file is a csv file with exactly 2 columns: password, user name. When using basic authentication from an http client, the server expects an `Authorization` header with the value of `Basic BASE64ENCODED(USER:PASSWORD)`

#### Static Token File

The server reads bearer tokens from a file when given the `--token-auth-file=SOMEFILE` option on the command line. Currently, tokens last indefinitely, and the token list cannot be changed without restarting the server.

The token file is a csv file with 1 column: token. When using bearer token authentication from an http client, the server expects an `Authorization` header with a value of `Bearer THETOKEN`. The bearer token must be a character sequence that can be put in an HTTP header value using no more than the encoding and quoting facilities of HTTP.

### SSL

The server can be secured via SSL by setting the `--tls-cert-file=SOMEFILE` and `--tls-private-key-file=SOMEFILE` cli flags. The generator will automatically generate self-signed certificates with a common name of `localhost`. If these cli flags are not set, the server will be served insecurely over HTTP.

### Docker

The generator create a `Dockerfile` that you can use to build a lightweight docker image for your application. Image size `~10MB`

```sh
make docker
docker build -t petrasphere/gomicro:latest .
```

### Build Scripts

Various aspects of your application are build via the `make` utility. Here are the provided scripts:

#### `make`

Install dependencies and build binary

#### `make install`

Install dependencies

#### `make local-dev`

Install dependencies, build binary, and start backend database

#### `make clean`

Remove binary

### Local Development

Unless you are using sqlite as a database driver, you will need too start up a database to test your application. So assist with this, a `docker-compose.yaml` file is provided and will start a database matching that of your driver. Here are the necessary configuration items that might be relevant as a side effect of using this feature:

#### MySQL

**Host**: `127.0.0.1`  
**Port**: `3306`  
**User**: admin  
**Password**: password  
**Database Name**: plural of resource noun  

#### PostgreSQL

**Host**: `127.0.0.1`  
**Port**: `5432`  
**User**: root  
**Password**: password  
**Database Name**: plural of resource noun  

### Deployment

Since it is recommended that your application be deployed inside of a Docker container, a container orchestration tool is needed in production. Robust configuration for both Kubernetes and Docker Swarm are provided.

### Error Handling

The server properly handlers all errors so that it does not have any unexpected results. If errors are encountered, they are written to the HTTP response with this body (as per the API documentation):

```json
{
  "code": 401,
  "msg": "unauthorized"
}
```

### Health Checks

It is often useful to have an HTTP route to diagnose application health. For example, this is used in the Kubernetes deployments. The `/health` route will attempt to connect to the database and return a `200` if successful or a `500` if it is not.
