// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "serj_bibox@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/submitData": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/submitData"
                ],
                "summary": "Создаёт новую запись в pereval_added",
                "parameters": [
                    {
                        "description": "карточка объекта",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pereval"
                        }
                    },
                    {
                        "description": "ID созданной записи",
                        "name": "output",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.Response"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apis.ErrResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/apis.ErrResponse"
                        }
                    }
                }
            }
        },
        "/submitData/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/submitData/:id"
                ],
                "summary": "Получает запись из pereval_added по ID записи",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apis.ErrResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/apis.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apis.ErrResponse": {
            "description": "Структура HTTP ответа об ошибке",
            "type": "object",
            "properties": {
                "code": {
                    "description": "application-specific error code",
                    "type": "integer"
                },
                "error": {
                    "description": "application-level error message, for debugging",
                    "type": "string"
                },
                "status": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        },
        "apis.Response": {
            "description": "Структура HTTP ответа: если отправка успешна, дополнительно возвращается id вставленной записи.",
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "123"
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "models.Images": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Спуск. Фото №99"
                },
                "url": {
                    "type": "string",
                    "example": "https://avatars.mds.yandex.net/i?id=a467876d3e1b1f0a84050103a206cf81-5858922-images-thumbs\u0026n=13"
                }
            }
        },
        "models.Pereval": {
            "type": "object",
            "properties": {
                "add_time": {
                    "type": "string",
                    "example": "2021-09-22 13:18:13"
                },
                "beautyTitle": {
                    "type": "string",
                    "example": "пер. "
                },
                "connect": {
                    "type": "string",
                    "example": " "
                },
                "coords": {
                    "type": "object",
                    "properties": {
                        "height": {
                            "type": "string",
                            "example": "1200"
                        },
                        "latitude": {
                            "type": "string",
                            "example": "45.3842"
                        },
                        "longitude": {
                            "type": "string",
                            "example": "7.1525"
                        }
                    }
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Images"
                    }
                },
                "level": {
                    "type": "object",
                    "properties": {
                        "autumn": {
                            "type": "string",
                            "example": "1A"
                        },
                        "spring": {
                            "type": "string",
                            "example": " "
                        },
                        "summer": {
                            "type": "string",
                            "example": "1A"
                        },
                        "winter": {
                            "type": "string",
                            "example": " "
                        }
                    }
                },
                "other_titles": {
                    "type": "string",
                    "example": "1"
                },
                "pereval_id": {
                    "type": "string",
                    "example": "125"
                },
                "title": {
                    "type": "string",
                    "example": "Туя-Ашуу"
                },
                "type": {
                    "type": "string",
                    "example": "pass"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dd@aa.ru"
                },
                "fam": {
                    "type": "string",
                    "example": "Скворцов"
                },
                "id": {
                    "type": "string",
                    "example": "11234"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "otc": {
                    "type": "string",
                    "example": "Кожедубович"
                },
                "phone": {
                    "type": "string",
                    "example": "+744434555"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ФСТР API",
	Description:      "API для взаимодействия приложения с сервером БД ФСТР.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
