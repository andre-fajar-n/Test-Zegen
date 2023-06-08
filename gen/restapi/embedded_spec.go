// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Zegen",
    "version": "1.0.0"
  },
  "paths": {
    "/health": {
      "get": {
        "security": [],
        "description": "Check if the App is Running",
        "tags": [
          "health"
        ],
        "summary": "Health Check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Health Check",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/author": {
      "post": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Create author",
        "tags": [
          "author"
        ],
        "summary": "Create",
        "operationId": "createAuthor",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createAuthorParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreateAuthor"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/author/{author_id}": {
      "put": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Update author",
        "tags": [
          "author"
        ],
        "summary": "Update",
        "operationId": "updateAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "author_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateAuthorParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Soft delete author",
        "tags": [
          "author"
        ],
        "summary": "Soft Delete",
        "operationId": "softDeleteAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "author_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success delete",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/book": {
      "post": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Create book",
        "tags": [
          "book"
        ],
        "summary": "Create",
        "operationId": "createBook",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createBookParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreateBook"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/book/{book_id}": {
      "put": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Update book",
        "tags": [
          "book"
        ],
        "summary": "Update",
        "operationId": "updateBook",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "book_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateBookParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Soft delete book",
        "tags": [
          "book"
        ],
        "summary": "Soft Delete",
        "operationId": "softDeleteBook",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "book_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success delete",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/login": {
      "post": {
        "security": [],
        "description": "Login",
        "tags": [
          "authentication"
        ],
        "summary": "Login",
        "operationId": "login",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success login",
            "schema": {
              "$ref": "#/definitions/successLogin"
            },
            "headers": {
              "token": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/register": {
      "post": {
        "security": [],
        "description": "Register user",
        "tags": [
          "authentication"
        ],
        "summary": "Register",
        "operationId": "register",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success register",
            "schema": {
              "$ref": "#/definitions/successRegister"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Author": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/authorData"
        },
        {
          "$ref": "#/definitions/authorForeignKey"
        }
      ]
    },
    "Book": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/bookData"
        },
        {
          "$ref": "#/definitions/bookForeignKey"
        }
      ]
    },
    "Principal": {
      "type": "object",
      "properties": {
        "expired_at": {
          "type": "string",
          "format": "date-time"
        },
        "user_id": {
          "type": "number",
          "format": "uint64"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "User": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/userData"
        }
      ]
    },
    "authorData": {
      "type": "object",
      "properties": {
        "country": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "authorForeignKey": {
      "type": "object",
      "properties": {
        "books": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/bookAuthor"
          }
        }
      }
    },
    "bookAuthor": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/bookAuthorData"
        }
      ]
    },
    "bookAuthorData": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "string"
        },
        "book_id": {
          "type": "string"
        }
      }
    },
    "bookData": {
      "type": "object",
      "properties": {
        "isbn": {
          "type": "string"
        },
        "publishedYear": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "bookForeignKey": {
      "type": "object",
      "properties": {
        "authors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/bookAuthor"
          }
        }
      }
    },
    "createAuthorParamsBody": {
      "type": "object",
      "required": [
        "name",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "createBookParamsBody": {
      "type": "object",
      "required": [
        "title",
        "isbn",
        "published_year",
        "author_ids"
      ],
      "properties": {
        "author_ids": {
          "type": "array",
          "items": {
            "type": "number",
            "minItems": 1,
            "uniqueItems": true
          }
        },
        "isbn": {
          "type": "string",
          "minLength": 1
        },
        "published_year": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "loginParamsBody": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "modelIdentifier": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "modelTrackTime": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:created_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "deleted_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:deleted_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:updated_at\"",
          "x-nullable": true,
          "x-omitempty": false
        }
      }
    },
    "registerParamsBody": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "minLength": 8
        },
        "username": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "successCreateAuthor": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successCreateAuthorAllOf1"
        }
      ]
    },
    "successCreateAuthorAllOf1": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successCreateBook": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successCreateBookAllOf1"
        }
      ]
    },
    "successCreateBookAllOf1": {
      "type": "object",
      "properties": {
        "book_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successLogin": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successLoginAllOf1"
        }
      ]
    },
    "successLoginAllOf1": {
      "type": "object",
      "properties": {
        "expired_at": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    },
    "successRegister": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successRegisterAllOf1"
        }
      ]
    },
    "successRegisterAllOf1": {
      "type": "object",
      "properties": {
        "user_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "updateAuthorParamsBody": {
      "type": "object",
      "required": [
        "name",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "updateBookParamsBody": {
      "type": "object",
      "required": [
        "title",
        "isbn",
        "published_year",
        "author_ids"
      ],
      "properties": {
        "author_ids": {
          "type": "array",
          "items": {
            "type": "number",
            "minItems": 1,
            "uniqueItems": true
          }
        },
        "isbn": {
          "type": "string",
          "minLength": 1
        },
        "published_year": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "userData": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "x-omitempty": false
        },
        "username": {
          "type": "string",
          "x-omitempty": false
        }
      }
    }
  },
  "securityDefinitions": {
    "authorization": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Zegen",
    "version": "1.0.0"
  },
  "paths": {
    "/health": {
      "get": {
        "security": [],
        "description": "Check if the App is Running",
        "tags": [
          "health"
        ],
        "summary": "Health Check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Health Check",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/author": {
      "post": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Create author",
        "tags": [
          "author"
        ],
        "summary": "Create",
        "operationId": "createAuthor",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createAuthorParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreateAuthor"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/author/{author_id}": {
      "put": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Update author",
        "tags": [
          "author"
        ],
        "summary": "Update",
        "operationId": "updateAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "author_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateAuthorParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Soft delete author",
        "tags": [
          "author"
        ],
        "summary": "Soft Delete",
        "operationId": "softDeleteAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "author_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success delete",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/book": {
      "post": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Create book",
        "tags": [
          "book"
        ],
        "summary": "Create",
        "operationId": "createBook",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createBookParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreateBook"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/book/{book_id}": {
      "put": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Update book",
        "tags": [
          "book"
        ],
        "summary": "Update",
        "operationId": "updateBook",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "book_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateBookParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authorization": []
          }
        ],
        "description": "Soft delete book",
        "tags": [
          "book"
        ],
        "summary": "Soft Delete",
        "operationId": "softDeleteBook",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "book_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success delete",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/login": {
      "post": {
        "security": [],
        "description": "Login",
        "tags": [
          "authentication"
        ],
        "summary": "Login",
        "operationId": "login",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success login",
            "schema": {
              "$ref": "#/definitions/successLogin"
            },
            "headers": {
              "token": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/register": {
      "post": {
        "security": [],
        "description": "Register user",
        "tags": [
          "authentication"
        ],
        "summary": "Register",
        "operationId": "register",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success register",
            "schema": {
              "$ref": "#/definitions/successRegister"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Author": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/authorData"
        },
        {
          "$ref": "#/definitions/authorForeignKey"
        }
      ]
    },
    "Book": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/bookData"
        },
        {
          "$ref": "#/definitions/bookForeignKey"
        }
      ]
    },
    "Principal": {
      "type": "object",
      "properties": {
        "expired_at": {
          "type": "string",
          "format": "date-time"
        },
        "user_id": {
          "type": "number",
          "format": "uint64"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "User": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/userData"
        }
      ]
    },
    "authorData": {
      "type": "object",
      "properties": {
        "country": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "authorForeignKey": {
      "type": "object",
      "properties": {
        "books": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/bookAuthor"
          }
        }
      }
    },
    "bookAuthor": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        },
        {
          "$ref": "#/definitions/bookAuthorData"
        }
      ]
    },
    "bookAuthorData": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "string"
        },
        "book_id": {
          "type": "string"
        }
      }
    },
    "bookData": {
      "type": "object",
      "properties": {
        "isbn": {
          "type": "string"
        },
        "publishedYear": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "bookForeignKey": {
      "type": "object",
      "properties": {
        "authors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/bookAuthor"
          }
        }
      }
    },
    "createAuthorParamsBody": {
      "type": "object",
      "required": [
        "name",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "createBookParamsBody": {
      "type": "object",
      "required": [
        "title",
        "isbn",
        "published_year",
        "author_ids"
      ],
      "properties": {
        "author_ids": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "isbn": {
          "type": "string",
          "minLength": 1
        },
        "published_year": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "loginParamsBody": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "modelIdentifier": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "modelTrackTime": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:created_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "deleted_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:deleted_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:updated_at\"",
          "x-nullable": true,
          "x-omitempty": false
        }
      }
    },
    "registerParamsBody": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "minLength": 8
        },
        "username": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "successCreateAuthor": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successCreateAuthorAllOf1"
        }
      ]
    },
    "successCreateAuthorAllOf1": {
      "type": "object",
      "properties": {
        "author_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successCreateBook": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successCreateBookAllOf1"
        }
      ]
    },
    "successCreateBookAllOf1": {
      "type": "object",
      "properties": {
        "book_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successLogin": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successLoginAllOf1"
        }
      ]
    },
    "successLoginAllOf1": {
      "type": "object",
      "properties": {
        "expired_at": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    },
    "successRegister": {
      "allOf": [
        {
          "$ref": "#/definitions/success"
        },
        {
          "$ref": "#/definitions/successRegisterAllOf1"
        }
      ]
    },
    "successRegisterAllOf1": {
      "type": "object",
      "properties": {
        "user_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "updateAuthorParamsBody": {
      "type": "object",
      "required": [
        "name",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "updateBookParamsBody": {
      "type": "object",
      "required": [
        "title",
        "isbn",
        "published_year",
        "author_ids"
      ],
      "properties": {
        "author_ids": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "isbn": {
          "type": "string",
          "minLength": 1
        },
        "published_year": {
          "type": "number",
          "format": "int64"
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      },
      "x-go-gen-location": "operations"
    },
    "userData": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "x-omitempty": false
        },
        "username": {
          "type": "string",
          "x-omitempty": false
        }
      }
    }
  },
  "securityDefinitions": {
    "authorization": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
