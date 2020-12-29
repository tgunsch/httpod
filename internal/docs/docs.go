// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete": {
            "delete": {
                "description": "delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/get": {
            "get": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/patch": {
            "patch": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/put": {
            "put": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP Methods"
                ],
                "summary": "The request's query parameters.",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/status/{code}": {
            "get": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status codes"
                ],
                "summary": "The request's query parameters.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "return status code",
                        "name": "code",
                        "in": "path"
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Informational responses",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "300": {
                        "description": "Redirection",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Client Errors",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Errors",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status codes"
                ],
                "summary": "The request's query parameters.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "return status code",
                        "name": "code",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status codes"
                ],
                "summary": "The request's query parameters.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "return status code",
                        "name": "code",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status codes"
                ],
                "summary": "The request's query parameters.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "return status code",
                        "name": "code",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status codes"
                ],
                "summary": "The request's query parameters.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "return status code",
                        "name": "code",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}