// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/google": {
            "post": {
                "description": "2000 성공\n2001 신규 유저 닉네임 입력 화면으로 이동\n4004 missing bodies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "구글 로그인",
                "deprecated": true,
                "parameters": [
                    {
                        "description": "google email , profile_image",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GoogleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/auth/google/token": {
            "get": {
                "description": "2000 성공\n2001 신규 유저 닉네임 입력 화면으로 이동\n4004 required token\n4003 not verify id token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "구글 idtoken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idtoken",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "get": {
                "description": "2000 성공",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "로그인 -\u003e 테스트 로그인",
                "deprecated": true,
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/auth/regist": {
            "post": {
                "description": "2001 성공\n4001 missing bodies\n4002 Cannot create user\n4003 Type is not match",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "회원가입",
                "deprecated": true,
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegistDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        },
        "/auth/token": {
            "post": {
                "description": "2000 성공\n4001 required refresh_token\n4101 required login\n4004 다른 웹에서 로그인되어 있습니다 다시 로그인 해주세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "토큰",
                "parameters": [
                    {
                        "description": "refresh token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TokenDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    }
                }
            }
        },
        "/health/check": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "health check",
                "responses": {}
            }
        },
        "/holiday": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n4003 missing bodies\n4004 이미 생성된 공휴일",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "holiday"
                ],
                "summary": "공휴일 or 매매 늦게 시작하는 날 설정",
                "parameters": [
                    {
                        "description": "holiday",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.HolidayDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/interest": {
            "get": {
                "description": "2000 성공",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interest"
                ],
                "summary": "관심종목",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset count",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.InterestListResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n4001 missing bodies\n4002 cannot add ticker\n4101 required login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interest"
                ],
                "summary": "관심종목 추가",
                "parameters": [
                    {
                        "description": "add interest",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SetInterestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.InterestListResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n4001 missing bodies\n4002 cannot delete[sql]\n4101 required login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interest"
                ],
                "summary": "관심종목 삭제",
                "parameters": [
                    {
                        "description": "delete interest",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteInterestDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/interest/sale": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n4002 missing bodies\n4101 required login\n4004 Cannot sale\n4005 The time to sale has over",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interest"
                ],
                "summary": "관심종목 매도",
                "parameters": [
                    {
                        "description": "sale interest",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SaleInterestDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/ticker/chart/{symbol}": {
            "get": {
                "description": "2000 성공\n2001 empty response\n4001 missing count\n4002 date format error yyyy-mm-dd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ticker"
                ],
                "summary": "종목 차트 가져오기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "마지막 날짜",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "count",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.OneTickerChartResponse"
                            }
                        }
                    }
                }
            }
        },
        "/ticker/comment": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n4101 required login\n4003 empty comment or symbol\n4004 missing bodies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "티커 코멘트 등록",
                "parameters": [
                    {
                        "description": "ticker comment",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewCommentDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.TickerCommentsResponse"
                            }
                        }
                    }
                }
            }
        },
        "/ticker/comment/{symbol}": {
            "get": {
                "description": "2000 성공\n4101 required login\n4004 count , created_at error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "티커 코멘트 목록",
                "parameters": [
                    {
                        "type": "string",
                        "description": "symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "count",
                        "name": "count",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "created_at",
                        "name": "created_at",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.TickerCommentsResponse"
                            }
                        }
                    }
                }
            }
        },
        "/ticker/search": {
            "get": {
                "description": "2000 성공\n4001 offset and count are int type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ticker"
                ],
                "summary": "종목 검색",
                "parameters": [
                    {
                        "type": "string",
                        "description": "검색할 종목명 , 코드",
                        "name": "word",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "가지고 있는 list.length",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "count",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.SearchTickerResponse"
                            }
                        }
                    }
                }
            }
        },
        "/ticker/{symbol}": {
            "get": {
                "description": "2000 성공\n4001 required count\n4002 cannot found list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ticker"
                ],
                "summary": "종목 가져오기",
                "parameters": [
                    {
                        "type": "string",
                        "description": "symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "count",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.OneTickerResponse"
                        }
                    }
                }
            }
        },
        "/user/nickname": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "2000 성공\n2001 nickname is exists\n4001 required nickname\n4101 required login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "닉네임 변경",
                "parameters": [
                    {
                        "description": "nickname",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NicknameDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CommentUser": {
            "type": "object",
            "properties": {
                "nickname": {
                    "description": "닉네임",
                    "type": "string"
                },
                "profile_image": {
                    "description": "프로필이미지",
                    "type": "string"
                }
            }
        },
        "dto.DeleteInterestDto": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "dto.GoogleDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                }
            }
        },
        "dto.HolidayDto": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "휴일 날짜",
                    "type": "string"
                },
                "name": {
                    "description": "휴일 이름",
                    "type": "string"
                },
                "opened_at": {
                    "description": "시간 [수능일 같은 경우 10시에 개장]",
                    "type": "string"
                },
                "type": {
                    "description": "1인 경우 opened_at 이후에 가능",
                    "type": "integer"
                }
            }
        },
        "dto.InterestList": {
            "type": "object",
            "properties": {
                "close": {
                    "description": "등록한 시점의 가격",
                    "type": "number"
                },
                "date_time": {
                    "description": "등록된 시점",
                    "type": "string"
                },
                "market": {
                    "description": "0 KOSPI 1 KOSDAQ",
                    "type": "integer"
                },
                "name": {
                    "description": "종목이름",
                    "type": "string"
                },
                "percent": {
                    "description": "등록한 시점의 등락률",
                    "type": "string"
                },
                "symbol": {
                    "description": "종목코드",
                    "type": "string"
                },
                "ticker": {
                    "$ref": "#/definitions/dto.InterestTickerModel"
                },
                "type": {
                    "description": "현재 사용하지 않습니다.",
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/dto.InterestUserModel"
                },
                "volume": {
                    "description": "등록된 시점의 거래량",
                    "type": "string"
                }
            }
        },
        "dto.InterestTickerChartModel": {
            "type": "object",
            "properties": {
                "close": {
                    "description": "현재가",
                    "type": "number"
                },
                "percent": {
                    "description": "현재 등락률",
                    "type": "string"
                },
                "volume": {
                    "description": "현재 거래량",
                    "type": "string"
                }
            }
        },
        "dto.InterestTickerModel": {
            "type": "object",
            "properties": {
                "homepage": {
                    "description": "홈페이지",
                    "type": "string"
                }
            }
        },
        "dto.InterestUserModel": {
            "type": "object",
            "properties": {
                "nickname": {
                    "description": "등록한 사용자의 닉네임",
                    "type": "string"
                }
            }
        },
        "dto.NewCommentDto": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "dto.NicknameDto": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                }
            }
        },
        "dto.RegistDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SaleInterestDto": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "dto.SetInterestDto": {
            "type": "object",
            "properties": {
                "symbol": {
                    "type": "string"
                }
            }
        },
        "dto.TokenDto": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "responses.InterestListResponse": {
            "type": "object",
            "properties": {
                "recent": {
                    "$ref": "#/definitions/dto.InterestTickerChartModel"
                },
                "ticker": {
                    "$ref": "#/definitions/dto.InterestList"
                }
            }
        },
        "responses.OneTickerChartResponse": {
            "type": "object",
            "properties": {
                "close": {
                    "description": "종가(현재가)",
                    "type": "number"
                },
                "date": {
                    "description": "날짜",
                    "type": "string"
                },
                "high": {
                    "description": "고가",
                    "type": "number"
                },
                "low": {
                    "description": "저가",
                    "type": "number"
                },
                "open": {
                    "description": "시가",
                    "type": "number"
                },
                "percent": {
                    "description": "등락률",
                    "type": "string"
                },
                "volume": {
                    "description": "거래량",
                    "type": "string"
                }
            }
        },
        "responses.OneTickerResponse": {
            "type": "object",
            "properties": {
                "bps": {
                    "type": "number"
                },
                "chart": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.OneTickerChartResponse"
                    }
                },
                "div": {
                    "type": "number"
                },
                "dps": {
                    "type": "number"
                },
                "eps": {
                    "type": "number"
                },
                "homepage": {
                    "description": "홈페이지 (로고 url)",
                    "type": "string"
                },
                "industry": {
                    "type": "string"
                },
                "listing_date": {
                    "description": "상장일",
                    "type": "string"
                },
                "market": {
                    "description": "0 : 코스피 , 1 : 코스닥",
                    "type": "integer"
                },
                "market_cap": {
                    "description": "시가총액",
                    "type": "number"
                },
                "name": {
                    "description": "종목명",
                    "type": "string"
                },
                "pbr": {
                    "description": "표시",
                    "type": "number"
                },
                "per": {
                    "description": "표시",
                    "type": "number"
                },
                "representative": {
                    "description": "대표자 (표시할 필요없습니다.)",
                    "type": "string"
                },
                "sector": {
                    "description": "표시",
                    "type": "string"
                },
                "settle_month": {
                    "description": "결산월",
                    "type": "string"
                },
                "symbol": {
                    "description": "index",
                    "type": "string"
                }
            }
        },
        "responses.SearchTickerResponse": {
            "type": "object",
            "properties": {
                "homepage": {
                    "description": "홈페이지 (로고 url)",
                    "type": "string"
                },
                "market": {
                    "description": "0 코스피 , 1 코스닥",
                    "type": "integer"
                },
                "name": {
                    "description": "종목명",
                    "type": "string"
                },
                "symbol": {
                    "description": "코드",
                    "type": "string"
                }
            }
        },
        "responses.TickerCommentsResponse": {
            "type": "object",
            "properties": {
                "comment": {
                    "description": "코멘트",
                    "type": "string"
                },
                "comment_id": {
                    "description": "코멘트 uuid 수정이나 삭제시 필요할 수도 있습니다.",
                    "type": "string"
                },
                "created_at": {
                    "description": "코멘트 등록일시",
                    "type": "string"
                },
                "symbol": {
                    "description": "메세지를 뿌려줄 symbol",
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/dto.CommentUser"
                }
            }
        },
        "responses.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "access token",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "refresh token",
                    "type": "string"
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "created_at": {
                    "description": "가입일",
                    "type": "string"
                },
                "email": {
                    "description": "email 마우스 호버시에만 보이게",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "description": "nickname",
                    "type": "string"
                },
                "profile_image": {
                    "description": "google login시 받아오는 profile image",
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
