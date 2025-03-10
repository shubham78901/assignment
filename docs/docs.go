// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/countries/search": {
            "get": {
                "description": "Fetches details of a country using its name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries"
                ],
                "summary": "Search country by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Country Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved country details",
                        "schema": {
                            "$ref": "#/definitions/model.Country"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Missing country name",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Country": {
            "type": "object",
            "properties": {
                "capital": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "currencies": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/model.Currency"
                    }
                },
                "name": {
                    "$ref": "#/definitions/model.Name"
                },
                "population": {
                    "type": "integer"
                }
            }
        },
        "model.Currency": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "model.Name": {
            "type": "object",
            "properties": {
                "common": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Country API",
	Description:      "This is an API for fetching country information.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
