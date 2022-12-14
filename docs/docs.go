// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://tos.santoshk.dev",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/persons": {
            "get": {
                "description": "Responds with the list of matched persons as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Persons"
                ],
                "summary": "Search persons by name or email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Not required.",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.PersonGetDto"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Takes a person JSON and store in DB. Return No Content.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Persons"
                ],
                "summary": "Update an existed person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PersonUpdateDto JSON",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PersonUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a person JSON and store in DB. Return saved ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Persons"
                ],
                "summary": "Create a new person",
                "parameters": [
                    {
                        "description": "PersonCreateDto JSON",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PersonCreateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Takes a person ID and removes it from the database. Return No Content.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Persons"
                ],
                "summary": "Delete an existed person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    }
                }
            }
        },
        "/persons/{id}": {
            "get": {
                "description": "Returns the person whose ID value matches the id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Persons"
                ],
                "summary": "Get single person by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PersonGetDto"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.PersonCreateDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.PersonGetDto": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.PersonUpdateDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "errs.ErrResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/errs.ServiceError"
                }
            }
        },
        "errs.ServiceError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "eCommerce API",
	Description:      "An eCommerce service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
