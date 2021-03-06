// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Momzor",
            "email": "m.benaida.pro@gmail.com"
        },
        "license": {
            "url": "https://www.gnu.org/licenses/quick-guide-gplv3.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/fizzbuzz": {
            "get": {
                "description": "get custom Fizzbuzz",
                "produces": [
                    "application/json"
                ],
                "summary": "Return Fizzbuzz resource",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int1 query parameter",
                        "name": "int1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "int2 query parameter",
                        "name": "int2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit query parameter",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "str1 query parameter",
                        "name": "str1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "str2 query parameter",
                        "name": "str2",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        },
        "/stats": {
            "get": {
                "description": "get The most used params",
                "produces": [
                    "application/json"
                ],
                "summary": "Return Stats resource"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "LBC test fizzbuzz",
	Description: "LBC test fizzbuzz",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
