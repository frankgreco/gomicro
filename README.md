# gomicro

 [![Travis](https://img.shields.io/travis/petrasphere/gomicro.svg?style=flat-square)](https://travis-ci.org/petrasphere/gomicro) [![Coveralls](https://img.shields.io/coveralls/petrasphere/gomicro.svg?style=flat-square)]()  [![npm](https://img.shields.io/npm/dm/generator-gomicro.svg?style=flat-square)](https://www.npmjs.com/package/generator-gomicro)

> a yeoman generator for golang rest microservices

> project backlog [here](https://github.com/petrasphere/gomicro/projects/1)

### Quick Start

```sh
$ npm install -g generator-gomicro
$ yo gomicro
```

### Overview

**gomicro** aims to provide a production grade generator for applications that implement a restful api through create/retrieve/update/delete (CRUD) operations. Here are the aspects that gomicro implements to make your application production ready:

#### Command Line Interface
#### Logging
#### API Documentation
#### Unit Testing
#### Database Driver
#### HTTP Access Control (CORS)
#### Linting
#### Authorization
#### SSL
#### Error Handling
#### Health Check
#### Docker
#### Build Scripts

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
