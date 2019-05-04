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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HackPrague 2019 API\n",
    "title": "HackPrague 2019",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/grade/{profile_id}": {
      "get": {
        "tags": [
          "data"
        ],
        "summary": "Get anonymous grade list",
        "operationId": "getGradeList",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Anonymous grade list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sample"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "data"
        ],
        "summary": "Post new grade",
        "operationId": "postNewGrade",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "Submitted"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/grade/{profile_id}/{user_id}": {
      "get": {
        "tags": [
          "data"
        ],
        "summary": "Get grade list for selected user",
        "operationId": "getUserGradeList",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Grade list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sample"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/profile": {
      "get": {
        "tags": [
          "profile"
        ],
        "summary": "Get profile list",
        "operationId": "getProfile",
        "responses": {
          "200": {
            "description": "Profile list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/profile"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      }
    },
    "profile": {
      "type": "object",
      "properties": {
        "description": {
          "description": "Profile description",
          "type": "string"
        },
        "id": {
          "description": "Profile ID",
          "type": "string"
        },
        "name": {
          "description": "Profile name",
          "type": "string"
        }
      }
    },
    "sample": {
      "type": "object",
      "properties": {
        "grade": {
          "type": "number",
          "format": "integer"
        },
        "lat": {
          "description": "Lat",
          "type": "number"
        },
        "lon": {
          "description": "Lon",
          "type": "number"
        },
        "ts": {
          "type": "number",
          "format": "integer"
        },
        "user_id": {
          "description": "User ID",
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HackPrague 2019 API\n",
    "title": "HackPrague 2019",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/grade/{profile_id}": {
      "get": {
        "tags": [
          "data"
        ],
        "summary": "Get anonymous grade list",
        "operationId": "getGradeList",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Anonymous grade list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sample"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "data"
        ],
        "summary": "Post new grade",
        "operationId": "postNewGrade",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "Submitted"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/grade/{profile_id}/{user_id}": {
      "get": {
        "tags": [
          "data"
        ],
        "summary": "Get grade list for selected user",
        "operationId": "getUserGradeList",
        "parameters": [
          {
            "type": "string",
            "name": "profile_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Grade list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/sample"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/profile": {
      "get": {
        "tags": [
          "profile"
        ],
        "summary": "Get profile list",
        "operationId": "getProfile",
        "responses": {
          "200": {
            "description": "Profile list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/profile"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      }
    },
    "profile": {
      "type": "object",
      "properties": {
        "description": {
          "description": "Profile description",
          "type": "string"
        },
        "id": {
          "description": "Profile ID",
          "type": "string"
        },
        "name": {
          "description": "Profile name",
          "type": "string"
        }
      }
    },
    "sample": {
      "type": "object",
      "properties": {
        "grade": {
          "type": "number",
          "format": "integer"
        },
        "lat": {
          "description": "Lat",
          "type": "number"
        },
        "lon": {
          "description": "Lon",
          "type": "number"
        },
        "ts": {
          "type": "number",
          "format": "integer"
        },
        "user_id": {
          "description": "User ID",
          "type": "string"
        }
      }
    }
  }
}`))
}
