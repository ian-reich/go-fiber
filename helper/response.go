package helper

import "github.com/gofiber/fiber/v2"

// SuccessResponse adalah struktur untuk respons sukses
type SuccessResponse struct {
	Meta struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

// CreatedResponse adalah struktur untuk respons pembuatan sukses
type CreatedResponse struct {
	Meta struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}

// ErrorResponse adalah struktur untuk respons error
type ErrorResponse struct {
	Meta struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}

// SendSuccessResponse mengirim respons sukses dengan data
func SendSuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	response := SuccessResponse{
		Meta: struct {
			Status  string `json:"status"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return c.Status(code).JSON(response)
}

// SendCreatedResponse mengirim respons pembuatan sukses
func SendCreatedResponse(c *fiber.Ctx, code int, message string) error {
	response := CreatedResponse{
		Meta: struct {
			Status  string `json:"status"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Code:    code,
			Message: message,
		},
	}

	return c.Status(code).JSON(response)
}

// SendErrorResponse mengirim respons error
func SendErrorResponse(c *fiber.Ctx, code int, message string) error {
	response := ErrorResponse{
		Meta: struct {
			Status  string `json:"status"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Status:  "error",
			Code:    code,
			Message: message,
		},
	}

	return c.Status(code).JSON(response)
}
