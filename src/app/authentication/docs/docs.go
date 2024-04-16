// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer-signin": {
            "post": {
                "description": "Signs in a user with provided credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User Sign In",
                "parameters": [
                    {
                        "description": "Sign In Credentials",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User signed in successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SignInCredentials"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer-signup": {
            "post": {
                "description": "Register a new customer and save their data in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new customer",
                "parameters": [
                    {
                        "description": "Customer data",
                        "name": "customerRecords",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Customer": {
            "type": "object",
            "required": [
                "email",
                "name",
                "pancardNumber",
                "password",
                "phoneNumber"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@gmail.com"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "John Doe"
                },
                "pancardNumber": {
                    "type": "string",
                    "example": "ABCDE1234F"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password"
                },
                "phoneNumber": {
                    "type": "integer",
                    "example": 1234567890
                }
            }
        },
        "models.SignInCredentials": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@gmail.com"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Stock Broker Application",
	Description:      "api for Stock Broker using gin and gorm",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
