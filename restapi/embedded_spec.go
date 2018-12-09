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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Next is a simple single node store for retrieving key/value information",
    "title": "Next",
    "contact": {
      "name": "Steve Nicholls",
      "email": "stevexnicholls@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "https://github.com/stevexnicholls/next/blob/master/LICENSE"
    },
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/backup": {
      "get": {
        "produces": [
          "application/octet-stream"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Get a copy of the store",
        "operationId": "BackupGet",
        "responses": {
          "200": {
            "description": "Yay a backup"
          },
          "404": {
            "$ref": "#/responses/ErrorNotFound"
          },
          "default": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      }
    },
    "/kv": {
      "get": {
        "description": "Lists all the keys",
        "tags": [
          "kv"
        ],
        "operationId": "KeyList",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyValue"
              }
            }
          },
          "default": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      },
      "put": {
        "security": [
          {
            "token": [
              "admin"
            ]
          }
        ],
        "tags": [
          "kv"
        ],
        "summary": "Update value",
        "operationId": "ValueUpdate",
        "parameters": [
          {
            "name": "keyvalue",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Updated successfully",
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          },
          "404": {
            "$ref": "#/responses/ErrorNotFound"
          },
          "default": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      }
    },
    "/kv/{key}": {
      "get": {
        "tags": [
          "kv"
        ],
        "summary": "Get a value",
        "operationId": "ValueGet",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          },
          "404": {
            "$ref": "#/responses/ErrorNotFound"
          },
          "default": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      },
      "delete": {
        "tags": [
          "kv"
        ],
        "summary": "Delete an existing key",
        "operationId": "KeyDelete",
        "responses": {
          "204": {
            "description": "Deleted successfully"
          },
          "404": {
            "$ref": "#/responses/ErrorNotFound"
          },
          "default": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/Key"
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "description": "the error model is a model for all the error responses coming from kvstore\n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "cause": {
          "$ref": "#/definitions/Error"
        },
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "KeyValue": {
      "type": "object",
      "properties": {
        "key": {
          "description": "The key",
          "type": "string"
        },
        "value": {
          "description": "The value",
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "parameters": {
    "Key": {
      "minLength": 1,
      "type": "string",
      "description": "The key for a given entry",
      "name": "key",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "ErrorNotFound": {
      "description": "The entry was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ErrorResponse": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "token": {
      "type": "apiKey",
      "name": "x-api-key",
      "in": "header"
    }
  },
  "security": [
    {
      "token": []
    }
  ],
  "tags": [
    {
      "description": "key value",
      "name": "kv"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Next is a simple single node store for retrieving key/value information",
    "title": "Next",
    "contact": {
      "name": "Steve Nicholls",
      "email": "stevexnicholls@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "https://github.com/stevexnicholls/next/blob/master/LICENSE"
    },
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/backup": {
      "get": {
        "produces": [
          "application/octet-stream"
        ],
        "tags": [
          "backup"
        ],
        "summary": "Get a copy of the store",
        "operationId": "BackupGet",
        "responses": {
          "200": {
            "description": "Yay a backup"
          },
          "404": {
            "description": "The entry was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/kv": {
      "get": {
        "description": "Lists all the keys",
        "tags": [
          "kv"
        ],
        "operationId": "KeyList",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyValue"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "token": [
              "admin"
            ]
          }
        ],
        "tags": [
          "kv"
        ],
        "summary": "Update value",
        "operationId": "ValueUpdate",
        "parameters": [
          {
            "name": "keyvalue",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Updated successfully",
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          },
          "404": {
            "description": "The entry was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/kv/{key}": {
      "get": {
        "tags": [
          "kv"
        ],
        "summary": "Get a value",
        "operationId": "ValueGet",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/KeyValue"
            }
          },
          "404": {
            "description": "The entry was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "kv"
        ],
        "summary": "Delete an existing key",
        "operationId": "KeyDelete",
        "responses": {
          "204": {
            "description": "Deleted successfully"
          },
          "404": {
            "description": "The entry was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "minLength": 1,
          "type": "string",
          "description": "The key for a given entry",
          "name": "key",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "description": "the error model is a model for all the error responses coming from kvstore\n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "cause": {
          "$ref": "#/definitions/Error"
        },
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "KeyValue": {
      "type": "object",
      "properties": {
        "key": {
          "description": "The key",
          "type": "string"
        },
        "value": {
          "description": "The value",
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "parameters": {
    "Key": {
      "minLength": 1,
      "type": "string",
      "description": "The key for a given entry",
      "name": "key",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "ErrorNotFound": {
      "description": "The entry was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ErrorResponse": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "token": {
      "type": "apiKey",
      "name": "x-api-key",
      "in": "header"
    }
  },
  "security": [
    {
      "token": []
    }
  ],
  "tags": [
    {
      "description": "key value",
      "name": "kv"
    }
  ]
}`))
}
