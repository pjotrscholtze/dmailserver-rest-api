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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Docker mailserver - OpenAPI 2.0",
    "version": "1.0.11"
  },
  "basePath": "/api",
  "paths": {
    "/email": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "listEmailAccounts",
        "responses": {
          "200": {
            "description": "Successful listing of email addresses",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/EmailAccountListItem"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "addEmailAccount",
        "parameters": [
          {
            "description": "Create a new email account",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EmailAccount"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "string"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "406": {
            "description": "Email address already exists!"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/email/{emailAddress}": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "updateEmailAddress",
        "parameters": [
          {
            "type": "string",
            "description": "Address of the email account to update",
            "name": "emailAddress",
            "in": "path",
            "required": true
          },
          {
            "description": "Password for the email account",
            "name": "password",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "Account not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "deleteEmailAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Address of the email account to delete",
            "name": "emailAddress",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "Account not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/fail2ban": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "getFail2banIps",
        "responses": {
          "200": {
            "description": "Successful listing of banned IPs by Fail2ban.",
            "schema": {
              "$ref": "#/definitions/Fail2banListItem"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "postFail2banIp",
        "parameters": [
          {
            "description": "IP address to add to fail2ban",
            "name": "ipaddress",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/fail2ban/{ip}": {
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "deleteFail2banIp",
        "parameters": [
          {
            "type": "string",
            "description": "IP address to delete",
            "name": "ip",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "IP was not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "EmailAccount": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "example@example.com"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "EmailAccountListItem": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "example@example.com"
        },
        "aliases": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "quota": {
          "$ref": "#/definitions/Quota"
        }
      }
    },
    "Fail2banListItem": {
      "type": "object",
      "properties": {
        "bannedInDovecot": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "bannedInPostfix": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "bannedInPostfixSasl": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Quota": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "string"
        },
        "usage": {
          "type": "string"
        },
        "usagePercentage": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Docker mailserver - OpenAPI 2.0",
    "version": "1.0.11"
  },
  "basePath": "/api",
  "paths": {
    "/email": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "listEmailAccounts",
        "responses": {
          "200": {
            "description": "Successful listing of email addresses",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/EmailAccountListItem"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "addEmailAccount",
        "parameters": [
          {
            "description": "Create a new email account",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EmailAccount"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "string"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "406": {
            "description": "Email address already exists!"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/email/{emailAddress}": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "updateEmailAddress",
        "parameters": [
          {
            "type": "string",
            "description": "Address of the email account to update",
            "name": "emailAddress",
            "in": "path",
            "required": true
          },
          {
            "description": "Password for the email account",
            "name": "password",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "Account not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "email"
        ],
        "operationId": "deleteEmailAccount",
        "parameters": [
          {
            "type": "string",
            "description": "Address of the email account to delete",
            "name": "emailAddress",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "Account not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/fail2ban": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "getFail2banIps",
        "responses": {
          "200": {
            "description": "Successful listing of banned IPs by Fail2ban.",
            "schema": {
              "$ref": "#/definitions/Fail2banListItem"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "postFail2banIp",
        "parameters": [
          {
            "description": "IP address to add to fail2ban",
            "name": "ipaddress",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/fail2ban/{ip}": {
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "fail2ban"
        ],
        "operationId": "deleteFail2banIp",
        "parameters": [
          {
            "type": "string",
            "description": "IP address to delete",
            "name": "ip",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation"
          },
          "404": {
            "description": "IP was not found"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "EmailAccount": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "example@example.com"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "EmailAccountListItem": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "example@example.com"
        },
        "aliases": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "quota": {
          "$ref": "#/definitions/Quota"
        }
      }
    },
    "Fail2banListItem": {
      "type": "object",
      "properties": {
        "bannedInDovecot": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "bannedInPostfix": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "bannedInPostfixSasl": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Quota": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "string"
        },
        "usage": {
          "type": "string"
        },
        "usagePercentage": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  }
}`))
}