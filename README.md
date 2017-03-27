# gomicro

 [![Travis](https://img.shields.io/travis/petrasphere/gomicro.svg?style=flat-square)](https://travis-ci.org/petrasphere/gomicro) [![Coveralls](https://img.shields.io/coveralls/petrasphere/gomicro.svg?style=flat-square)]()  [![npm](https://img.shields.io/npm/dm/generator-gomicro.svg?style=flat-square)](https://www.npmjs.com/package/generator-gomicro)

> a yeoman generator for golang rest microservices

> project backlog [here](https://github.com/petrasphere/gomicro/projects/1)

## Quick Start

```sh
$ npm install -g generator-gomicro
$ yo gomicro
```

## Table of Contents
* [Overview](#overview)
* [Components](#components)
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

## Overview

gomicro aims to provide a production grade generator for applications that implement a restful api through create/retrieve/update/delete (CRUD) operations. Here are the aspects that gomicro implements to make your application production ready:

## Components

### HTTP Router

> [github.com/gorilla/mux](https://github.com/gorilla/mux)

### Command Line Interface

> [github.com/spf13/cobra](https://github.com/spf13/cobra)

### Database Driver

> [github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)

The generator provides support for multiple backend database drivers configurable by command line flags.

##### MySQL

```sh
--db-port=DATABASE-PORT or DB_PORT=DATABASE-PORT
--db-host=DATABASE-HOST or DB_HOST=DATABASE-HOST  
--db-user=DATABASE-USER or DB_USER=DATABASE-USER  
--db-pass=DATABASE-PASS or DB_PASS=DATABASE-PASS  
--db-name=DATABASE-NAME or DB_NAME=DATABASE-NAME  
```

##### Sqlite

```sh
--db-location=DATABASE-LOCATION or DB_PORT=DATABASE-LOCATION
--db-name=DATABASE-NAME or DB_NAME=DATABASE-NAME
```

##### PostgreSQL

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
### Unit Testing
### HTTP Access Control (CORS)
### Logging

> [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

### Authorization

##### Static Password File

Basic authentication is enabled by passing the `--basic-auth-file=SOMEFILE` option to the server. Currently, the basic auth credentials last indefinitely, and the password(s) cannot be changed without restarting the server.

The basic auth file is a csv file with exactly 2 columns: password, user name. When using basic authentication from an http client, the server expects an `Authorization` header with the value of `Basic BASE64ENCODED(USER:PASSWORD)`

##### Static Token File

The server reads bearer tokens from a file when given the `--token-auth-file=SOMEFILE` option on the command line. Currently, tokens last indefinitely, and the token list cannot be changed without restarting the server.

The token file is a csv file with 1 column: token. When using bearer token authentication from an http client, the server expects an `Authorization` header with a value of `Bearer THETOKEN`. The bearer token must be a character sequence that can be put in an HTTP header value using no more than the encoding and quoting facilities of HTTP.

### SSL

The server can be secured via SSL by setting the `--tls-cert-file=SOMEFILE` and `--tls-private-key-file=SOMEFILE` cli flags. The generator will automatically generate self-signed certificates with a common name of `localhost`. If these cli flags are not set, the server will be served insecurely over HTTP.

### Docker
### Build Scripts
### Local Development
### Deployment

Since it is recommended that your application be deployed inside of a Docker container, a container orchestration tool is needed in production. Robust configuration for both Kubernetes and Docker Swarm are provided.

### Error Handling
### Health Checks
