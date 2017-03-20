> a yeoman generator for golang rest microservices

__*NOTE*__: this project is under active development. Here is the current backlog:
- [x] CRUD http router
- [x] mySQL db driver
- [x] robust logging
- [x] Docker configuration
- [ ] linting
- [x] build scripts
- [ ] mongodb db driver
- [x] cli tools
- [ ] unit testing
- [ ] coverage metrics with Coveralls
- [ ] build metrics with TravisCI
- [ ] support for multiple go package managers

## why gomicro
Whether you're a beginner with go or not, valuable time is spent setting up a new workspace. Here are some of the reasons to use gomicro to setup yours:

* **more than boilerplate**: most of the time, the majority of code when writing multiple RESTful microservices is the same. Things such as CORS, auth, logging, cli, testing, database drivers, error handling, etc. remain the same while the only thing changing is the data model. gomicro aims to take advantage of this by providing a standardized project utilizing the most widely used libraries so that all you have to worry about is the data model.  
* **documentation**: use the dynamically generated Swagger API documentation to professionally document your RESTful API   
* **unit testing**: gomicro provides robust unit testing for all your code   
* **containerized**: use the provided Dockerfile to containerize your application. Also, use the provided configuration to quickly spin up a lightweight database of your choice using Docker to enhance local development.   
* **command line interface**: use a cli to dynamically configure your application at runtime  
* **multiple backends**: choose from a variety of backends to be automatically generated for you.

## dependencies
* **go**: since this is a template for a project written in go we should install it. Find full installation instructions [here](https://golang.org/doc/install).
* **glide**: [glide](https://glide.sh) is a popular dependency management tool for go. This project uses it to manage dependencies in a fine grained way. Find full installation instructions [here](https://glide.sh).
* **gopath**: in order to properly setup a new go project, it is important that we have a proper gopath configured. Find out how to do so [here](https://golang.org/doc/code.html#GOPATH).
* **docker(optional)**: this dependency is used to create a lightweight backend to enhance local development. Full installation instructions for Mac can be found [here](https://docs.docker.com/docker-for-mac/install/) and for Windows [here](https://docs.docker.com/docker-for-windows/install/).
* **npm**: [npm](https://www.npmjs.com) is a package manager for Javascript. While this project bootstraps a go application, because the generator is written in Node.js, npm is needed. Find full installation instructions [here](http://blog.npmjs.org/post/85484771375/how-to-install-npm).
* **yeoman**: [yeoman](http://yeoman.io) is nothing more than an npm package that we'll use to run our generator. Install if by running `npm install -g yo`

## quick start

```sh
$ npm install -g generator-gomicro
$ yo gomicro
```

## go libraries
This project aims to use the best and most widely used libraries. Here are the main libraries that gomicro uses:

**command line interface**: [spf13/cobra](https://github.com/spf13/cobra)  
**http router**: [gorilla/mux](https://github.com/gorilla/mux)  
**logging**: [sirupsen/logrus](https://github.com/sirupsen/logrus)  

## usage
So you know what to expect, here is a sample cli interaction:

```sh
$ yo gomicro

=====================================================================
                         Welcome to gomicro!
This Yeoman generator aims to scaffold a robust RESTful microservice.
                         Let's get started!
=====================================================================

=====================================================================
           Go workspaces like to be created a certain way
           (e.g. $GOPATH/src/github.com/frankgreco/gohttp)
      To do this, let's get some information about your project
=====================================================================

? your name Frank B Greco Jr
? your email frank@petrasphere.io
? vcs github.com
? github.com username frankgreco
? project name <%= nounSingularLower %>
? resource noun (singular) person
? resource noun (plural) people
? database type mysql

=====================================================================
                   Creating your workspace...
=====================================================================

create Documents/projects/gopath/src/github.com/frankgreco/gomicro/Makefile
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/table.sql
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/Dockerfile
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/glide.yaml
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/swagger.json
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/docker-compose.yaml
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/.gitignore
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/main.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/route/logger.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/route/router.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/route/routes.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/cmd/root.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/cmd/start.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/cmd/version.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/database/mysql.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/utils/error.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/utils/flag.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/server/server.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/handler/handler.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/handler/util.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/database/driver.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/handler/person.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/handler/people.go
create Documents/projects/gopath/src/github.com/frankgreco/gomicro/models/person.go

=====================================================================
   You're almost done! Your workspace has been created here:
            $GOPATH/github.com/frankgreco/gomicro/
To complete your setup, run the following commands in your workspace:
            $ make             (use your own database)
            $ make local-dev   (or, create a local database)
            $ ./gomicro --help  (example usage)
=====================================================================

$ cd $GOPATH/github.com/frankgreco/gomicro/
$ make local-dev
$ ./gomicro start --app-port=8080 --db-port=3306 --db-host=localhost --db-name=calls --db-user=root --db-pass=password
$ curl -X POST -H "Content-Type: application/json" -d '{
    "paramOne": "foo",
    "paramTwo": "bar"
}' "http://localhost:8080/people"

```
