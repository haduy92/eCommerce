# eCommerce Go API

A RESTful API project built with Golang to demonstrate some basic CRUD operations.

## Frameworks used by this project

- [Golang](https://go.dev/) - Programming language
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) - Fastest web framework according to [gin-gonic.com](https://gin-gonic.com/docs/benchmarks)
- [Gorm](https://gorm.io/) - ORM library
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
3. After execution is completed, open a browser and access swagger page at

```http
http://localhost:8080/swagger/index.html
```
### Test API

1. Open terminal/command prompt at the root director of the project and run docker-compose up:

```shell
go test ./...
```
2. It should display test results of dto, repository and service