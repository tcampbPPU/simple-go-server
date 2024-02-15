package harmony

import (
	"github.com/gofiber/fiber/v2"
)

type SimplePaginationInfo struct {
	Total    int `json:"total"`
	PerPage  int `json:"per_page"`
	Current  int `json:"current"`
	LastPage int `json:"last_page"`
}

type CursorPaginationInfo struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
}

type ResponseDTO struct {
	Type    string `json:"type"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Meta    any    `json:"meta"`
	Data    any    `json:"data"`

	Pagination SimplePaginationInfo `json:"pagination"`
	Cursor     CursorPaginationInfo `json:"cursor"`
}

func Harmony(c *fiber.Ctx, status int, response ResponseDTO) error {
	return c.Status(status).JSON(response)
}

func base(c *fiber.Ctx, status int, message string, data ...interface{}) error {
	var responseType string
	switch status {
	case 200:
		responseType = "success"
	case 400:
		responseType = "client_error"
	case 500:
		responseType = "server_error"
	case 404:
		responseType = "not_found"
	case 401:
		responseType = "unauthorized"
	case 403:
		responseType = "forbidden"
	case 422:
		responseType = "unprocessable_entity"
	case 429:
		responseType = "too_many_requests"
	case 503:
		responseType = "service_unavailable"
	case 504:
		responseType = "gateway_timeout"
	case 505:
		responseType = "http_version_not_supported"
	default:
		responseType = "success"
	}

	var success bool
	if responseType == "success" {
		success = true
	} else {
		success = false
	}

	json := ResponseDTO{Type: responseType, Success: success, Message: message}

	if len(data) > 0 {
		json.Data = data[0]
	}

	return c.Status(status).JSON(json)
}

func Success(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 200, message, data...)
}

func BadRequest(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 400, message, data...)
}

func ServerError(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 500, message, data...)
}

func NotFound(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 404, message, data...)
}

func Unauthorized(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 401, message, data...)
}

func Forbidden(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 403, message, data...)
}

func TooManyRequests(c *fiber.Ctx, message string, data ...interface{}) error {
	return base(c, 429, message, data...)
}
