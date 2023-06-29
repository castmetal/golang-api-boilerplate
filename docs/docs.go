// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Michel La Guardia",
            "url": "https://www.github.com/castmetal",
            "email": "mlaguardia@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/example": {
            "get": {
                "description": "Listing all examples that was stored in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListAll Example"
                ],
                "summary": "List all examples in database",
                "parameters": [
                    {
                        "type": "string",
                        "format": "numeric",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "numeric",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ListAllExampleResponseDTO"
                        }
                    }
                }
            },
            "post": {
                "description": "Creating an example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Example"
                ],
                "summary": "Create an example based on the name input",
                "parameters": [
                    {
                        "description": "CreateExample Data",
                        "name": "createExample",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateExampleDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateExampleResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateExampleDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 2
                }
            }
        },
        "dtos.CreateExampleResponseDTO": {
            "type": "object",
            "properties": {
                "example": {
                    "$ref": "#/definitions/dtos.ExampleResponseDTO"
                }
            }
        },
        "dtos.ExampleResponseDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dtos.ListAllExampleResponseDTO": {
            "type": "object",
            "properties": {
                "examples": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.ExampleResponseDTO"
                    }
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:8088",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Example API",
	Description:      "This is a sample server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
