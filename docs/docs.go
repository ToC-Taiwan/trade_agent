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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/balance": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "GetAllBalance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbagent.Balance"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "ImportBalance",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbagent.Balance"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "DeletaAllBalance",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cache": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cache"
                ],
                "summary": "GetTradeDayTargets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/cache/{key}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cache"
                ],
                "summary": "GetTradeDayTargets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/config": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "GetAllConfig",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.Config"
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetAllOrder",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbagent.OrderStatus"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "ImportOrder",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbagent.OrderStatus"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "DeletaAllOrder",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/targets": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Targets"
                ],
                "summary": "GetTradeDayTargets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbagent.Target"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Analyze": {
            "type": "object",
            "properties": {
                "close_change_ratio_high": {
                    "type": "number"
                },
                "close_change_ratio_low": {
                    "type": "number"
                },
                "in_out_ratio": {
                    "type": "number"
                },
                "max_loss": {
                    "type": "number"
                },
                "open_close_change_ratio_high": {
                    "type": "number"
                },
                "open_close_change_ratio_low": {
                    "type": "number"
                },
                "out_in_ratio": {
                    "type": "number"
                },
                "rsi_high": {
                    "type": "number"
                },
                "rsi_low": {
                    "type": "number"
                },
                "rsi_min_count": {
                    "type": "integer"
                },
                "tick_analyze_max_period": {
                    "type": "number"
                },
                "tick_analyze_min_period": {
                    "type": "number"
                },
                "volume_pr_high": {
                    "type": "number"
                },
                "volume_pr_low": {
                    "type": "number"
                }
            }
        },
        "config.Config": {
            "type": "object",
            "properties": {
                "analyze": {
                    "$ref": "#/definitions/config.Analyze"
                },
                "database": {
                    "$ref": "#/definitions/config.Database"
                },
                "mqtt": {
                    "$ref": "#/definitions/config.MQTT"
                },
                "quota": {
                    "$ref": "#/definitions/config.Quota"
                },
                "schedule": {
                    "$ref": "#/definitions/config.Schedule"
                },
                "server": {
                    "$ref": "#/definitions/config.Server"
                },
                "switch": {
                    "$ref": "#/definitions/config.Switch"
                },
                "target_cond": {
                    "$ref": "#/definitions/config.TargetCond"
                },
                "trade": {
                    "$ref": "#/definitions/config.Trade"
                }
            }
        },
        "config.Database": {
            "type": "object",
            "properties": {
                "database": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "passwd": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "time_zone": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "config.MQTT": {
            "type": "object",
            "properties": {
                "ca_path": {
                    "type": "string"
                },
                "cert_path": {
                    "type": "string"
                },
                "client_id": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "key_path": {
                    "type": "string"
                },
                "passwd": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "config.Quota": {
            "type": "object",
            "properties": {
                "fee_discount": {
                    "type": "number"
                },
                "trade_fee_ratio": {
                    "type": "number"
                },
                "trade_quota": {
                    "type": "integer"
                },
                "trade_tax_ratio": {
                    "type": "number"
                }
            }
        },
        "config.Schedule": {
            "type": "object",
            "properties": {
                "clean_event": {
                    "type": "string"
                },
                "restart_sinopac": {
                    "type": "string"
                }
            }
        },
        "config.Server": {
            "type": "object",
            "properties": {
                "http_port": {
                    "type": "string"
                },
                "run_mode": {
                    "type": "string"
                },
                "sinopac_srv_host": {
                    "type": "string"
                },
                "sinopac_srv_port": {
                    "type": "string"
                }
            }
        },
        "config.Switch": {
            "type": "object",
            "properties": {
                "buy": {
                    "type": "boolean"
                },
                "buy_later": {
                    "type": "boolean"
                },
                "forward_max": {
                    "type": "integer"
                },
                "mean_time_forward": {
                    "type": "integer"
                },
                "mean_time_reverse": {
                    "type": "integer"
                },
                "reverse_max": {
                    "type": "integer"
                },
                "sell": {
                    "type": "boolean"
                },
                "sell_first": {
                    "type": "boolean"
                },
                "simulation": {
                    "type": "boolean"
                }
            }
        },
        "config.TargetCond": {
            "type": "object",
            "properties": {
                "black_category": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "black_stock": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "limit_price_high": {
                    "type": "number"
                },
                "limit_price_low": {
                    "type": "number"
                },
                "limit_volume": {
                    "type": "integer"
                },
                "real_time_targets_count": {
                    "type": "integer"
                }
            }
        },
        "config.Trade": {
            "type": "object",
            "properties": {
                "history_close_period": {
                    "type": "integer"
                },
                "history_kbar_period": {
                    "type": "integer"
                },
                "history_tick_period": {
                    "type": "integer"
                },
                "hold_time_from_open": {
                    "type": "number"
                },
                "total_open_time": {
                    "type": "number"
                },
                "trade_in_end_time": {
                    "type": "number"
                },
                "trade_in_wait_time": {
                    "type": "integer"
                },
                "trade_out_end_time": {
                    "type": "number"
                },
                "trade_out_wait_time": {
                    "type": "integer"
                }
            }
        },
        "dbagent.Balance": {
            "type": "object",
            "properties": {
                "discount": {
                    "type": "integer"
                },
                "forward": {
                    "type": "integer"
                },
                "original_balance": {
                    "type": "integer"
                },
                "reverse": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "trade_count": {
                    "type": "integer"
                },
                "trade_day": {
                    "type": "string"
                }
            }
        },
        "dbagent.OrderStatus": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                },
                "order_id": {
                    "type": "string"
                },
                "order_time": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "stock": {
                    "$ref": "#/definitions/dbagent.Stock"
                },
                "stock_id": {
                    "type": "integer"
                }
            }
        },
        "dbagent.Stock": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "day_trade": {
                    "type": "boolean"
                },
                "exchange": {
                    "type": "string"
                },
                "last_close": {
                    "type": "number"
                },
                "last_close_change_pct": {
                    "type": "number"
                },
                "last_volume": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "dbagent.Target": {
            "type": "object",
            "properties": {
                "rank": {
                    "type": "integer"
                },
                "real_time_add": {
                    "type": "boolean"
                },
                "stock": {
                    "$ref": "#/definitions/dbagent.Stock"
                },
                "stock_id": {
                    "type": "integer"
                },
                "trade_day": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "routers.ErrorResponse": {
            "type": "object",
            "properties": {
                "attachment": {},
                "response": {
                    "type": "string"
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
	Version:     "2.0.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Trade Agent",
	Description: "API docs for Trade Agent",
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
	swag.Register("swagger", &s{})
}
