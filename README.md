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

#### Command Line Interface
#### Database Driver
#### API Documentation
#### Unit Testing
#### HTTP Access Control (CORS)
#### Logging
#### Authorization
#### SSL
#### Docker
#### Build Scripts
#### Local Development
#### Deployment
#### Error Handling
#### Health Checks

## database
A user will be prompted to choose between the following database options for their datastore:
* MySQL
* Sqlite
* PostgreSQL

## ssl
If `https` is selected by the user when prompted, self-signed certificates will automatically be generated and the application will be served over ssl.

## go libraries
This project aims to use the best and most widely used libraries. Here are the main libraries that gomicro uses:

**command line interface**: [spf13/cobra](https://github.com/spf13/cobra)  
**http router**: [gorilla/mux](https://github.com/gorilla/mux)  
**logging**: [sirupsen/logrus](https://github.com/sirupsen/logrus)
**ORM library**: [jinzhu/gorm](https://github.com/jinzhu/gorm)  
