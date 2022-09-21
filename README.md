# eCommerce Go API

A RESTful API project built with Golang to demonstrate some basic CRUD operations.

## Frameworks used by this project

- [Golang](https://go.dev/) - Programming language
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) - Fastest web framework according to [gin-gonic.com](https://gin-gonic.com/docs/benchmarks)
- [Gorm](https://gorm.io/) - ORM library
- [Swagger](https://github.com/swaggo/gin-swagger) - Swagger documentation
- [Logrus](https://github.com/sirupsen/logrus) - Log Library
- [SqlMock](https://github.com/DATA-DOG/go-sqlmock) - SQL Mock testing library
- [GoMock](https://github.com/golang/mock) - Mocking framework for Go

## Project structure

```
eCommerce
 ┣ config
 ┃ ┗ database.go
 ┣ controller
 ┃ ┗ person.controller.go
 ┣ docs
 ┃ ┣ docs.go
 ┃ ┣ swagger.json
 ┃ ┗ swagger.yaml
 ┣ frontend
 ┃ ┣ css
 ┃ ┃ ┣ normalize.css
 ┃ ┃ ┗ style.css
 ┃ ┣ js
 ┃ ┃ ┣ app.js
 ┃ ┃ ┗ jquery-3.6.1.min.js
 ┃ ┣ favicon.ico
 ┃ ┣ icon.png
 ┃ ┣ icon.svg
 ┃ ┗ index.html
 ┣ infrastructure
 ┃ ┣ errs
 ┃ ┃ ┣ errs.go
 ┃ ┃ ┗ http.error.go
 ┃ ┗ middleware
 ┃ ┃ ┣ correlation.id.middleware.go
 ┃ ┃ ┣ error.handler.middleware.go
 ┃ ┃ ┣ header.middleware.go
 ┃ ┃ ┗ logger.middleware.go
 ┣ model
 ┃ ┣ dto
 ┃ ┃ ┣ person.dto.go
 ┃ ┃ ┗ person.dto_test.go
 ┃ ┗ entity
 ┃ ┃ ┗ person.go
 ┣ repository
 ┃ ┣ mock
 ┃ ┃ ┗ person.repository.mock.go
 ┃ ┣ person.repository.go
 ┃ ┗ person.repository_test.go
 ┣ service
 ┃ ┣ person.service.go
 ┃ ┗ person.service_test.go
 ┣ sql
 ┃ ┗ init.sql
 ┣ .env
 ┣ .gitignore
 ┣ docker-compose.yml
 ┣ Dockerfile
 ┣ go.mod
 ┣ go.sum
 ┣ main.go
 ┗ README.md
```

## Minimum Requirements

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)

## Getting Started

### Run API

1. Start Docker container
2. Open terminal/command prompt at the root director of the project and run docker-compose up:

```shell
docker-compose up -d --built
```
3. After execution is completed, swagger page should be accessible at

```http
http://localhost:8080/swagger/index.html
```
4. Navigate to folder frontend and open index.html, this is the UI page which allows to filter person list

> #### Note:
>
> Due to my lack of experience with go, there is an issue with the cors policy that I could not fix:
>
> GET /persons returned CORS policy error but GET /persons?q=a doesn't and works as expected

### Test API

1. Open terminal/command prompt at the root director of the project and run docker-compose up:

```shell
go test ./...
```
2. It should display test results of dto, repository and service