# Simple Go Backend 

This is a simple boilerplate backend built with Go in REST API.

## How to run this project

To run this project you need Go installed in your system. Follow the instructions [here](https://go.dev/doc/install) to install Golang.

After installing Go, install all the dependencies using `go get ./...`

Then run this project using `go run .`

The default port is set to `2211` you can change the port to anything. 

## Default routes 

GET    `/todos/all`:      Gives back all the todos
GET    `/todos/:id`:      When given a vaild ID returns details of that todo
POST   `/todos/create`:   Creates a new Todo when supplied with the todo description (with the key desc) as a json body. This sets the isDone key to be false by default and gives it an ID
PUT    `/todos/update`:   Supplied with the isDone and description update along with the id in the body, the todo gets updated. Note: you need to pass all the required elements to form a todo, id, isDone and desc to properly execute this route.
DELETE `/todos/delete`:   Deletes a todo when supplied with the ID in the body as JSON

## Future of this project 

Connecting a database and a frontend to this project / API.

## Dependencies used 

1. Gin-gonic - https://gin-gonic.com/ - A defacto backend library for Go
2. Shortid - https://pkg.go.dev/github.com/teris-io/shortid#section-readme - A library to generate shortid for project. Seems sometimes it generates non URL friendly IDs, will give a look into it. 
