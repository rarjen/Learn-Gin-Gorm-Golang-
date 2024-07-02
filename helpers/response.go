package helpers

import "net/http"

type Response struct {
	Meta Meta `json:"meta"`

	// Data nilanya bisa flexibel maka menggunakan interface kosong
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponseSuccess(message string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusOK,
		Status:  "success",
	}
	customResponse := Response{
		Meta: meta,
		Data: data,
	}
	return customResponse
}

func APIResponseUnprocessableEntity(message string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
		Status:  "unprocessable entity",
	}
	customResponse := Response{
		Meta: meta,
		Data: data,
	}
	return customResponse
}

func APIResponseNotFound(message string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusNotFound,
		Status:  "not found",
	}
	customResponse := Response{
		Meta: meta,
		Data: data,
	}
	return customResponse
}

func APIResponseCreated(message string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusCreated,
		Status:  "created",
	}
	customResponse := Response{
		Meta: meta,
		Data: data,
	}
	return customResponse
}

func ApiResponseBadRequest(message string) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusBadRequest,
		Status:  "bad request",
	}
	customResponse := Response{
		Meta: meta,
		Data: nil,
	}
	return customResponse
}

func APIResponseUnauthorized(message string) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusUnauthorized,
		Status:  "unauthorized",
	}
	customResponse := Response{
		Meta: meta,
		Data: nil,
	}
	return customResponse
}

func APIResponseForbidden(message string) Response {
	meta := Meta{
		Message: message,
		Code:    http.StatusForbidden,
		Status:  "forbidden",
	}
	customResponse := Response{
		Meta: meta,
		Data: nil,
	}
	return customResponse
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonRequest := Response{
		Meta: meta,
		Data: data,
	}
	return jsonRequest
}
