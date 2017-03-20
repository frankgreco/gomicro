> a yeoman generator for golang rest microservices

## why gomicro
Whether you're a beginner with go or not, valuable time is spent setting up a new workspace. Here are some of the reasons to use gomicro to setup yours:

**more than boilerplate**: most of the time, the majority of code when writing multiple RESTful microservices is the same. Things such as CORS, auth, logging, cli, testing, database drivers, error handling, etc. remain the same while the only thing changing is the data model. gomicro aims to take advantage of this by providing a standardized project utilizing the most widely used libraries so that all you have to worry about is the data model.   
**unit testing**: gomicro provides robust unit testing for all your code   
**containerized**: use the provided Dockerfile to containerize your application. Also, use the provided configuration to quickly spin up a lightweight database of your choice using Docker to enhance local development.   
**command line interface**: use a cli to dynamically configure your application at runtime  
**multiple backends**: choose from a variety of backends to be automatically generated for you.

## dependencies
* **go**: since this is a template for a project written in go we should install it. Find full installation instructions [here](https://golang.org/doc/install).
* **glide**: [glide](https://glide.sh) is a popular dependency management tool for go. This project uses it to manage dependencies in a fine grained way. Find full installation instructions [here](https://glide.sh).
* **gopath**: in order to properly setup a new go project, it is important that we have a proper gopath configured. Find out how to do so [here](https://golang.org/doc/code.html#GOPATH).
* **npm**: [npm](https://www.npmjs.com) is a package manager for Javascript. While this project bootstraps a go application, because the generator is written in Node.js, npm is needed. Find full installation instructions [here](http://blog.npmjs.org/post/85484771375/how-to-install-npm).
* **yeoman**: [yeoman](http://yeoman.io) is nothing more than an npm package that we'll use to run our generator. Install if by running `npm install -g yo`

## quick start

```sh
$ npm install -g generator-gomicro
$ yo gomicro
```

## usage
