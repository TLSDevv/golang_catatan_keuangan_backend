package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ErrInternalServer = "internal server error"
	ErrUnauthorized   = "unauthorized"
	ErrNotFound       = "not found"
	ErrValidation     = "validation error"
	ErrUserNotFound   = "user not found"
)

type ResponseBody struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetParams(r *http.Request, params string) int {
	paramString, _ := strconv.ParseInt(mux.Vars(r)[params], 10, 64)
	return int(paramString)
}

func Decode(r *http.Request, toBody interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(toBody); err != nil {
		return err
	}
	return nil
}

func Send(w http.ResponseWriter, responseBody interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return Encode(w, responseBody)
}

func SendError(w http.ResponseWriter, errorMessage string, statusCode int, errors []string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResponse := ErrorResponse{Success: false, Message: errorMessage}
	if errors != nil {
		errResponse.Errors = errors
	}

	return Encode(w, errResponse)
}

func SendSuccess(w http.ResponseWriter, successMessage string, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return Encode(w, SuccessResponse{
		Success: true,
		Message: successMessage,
		Data:    data,
	})
}

func Encode(w http.ResponseWriter, responseBody interface{}) error {
	return json.NewEncoder(w).Encode(responseBody)
}

func SendWithData(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resBody := ResponseBody{
		Code:   statusCode,
		Data:   data,
		Status: message,
	}

	if len(message) != 0 {
		resBody.Status = message
	}

	Send(w, resBody, statusCode)
}

func SendNoData(w http.ResponseWriter, statusCode int, message string) {
	resBody := ResponseBody{
		Code:   statusCode,
		Status: message,
	}

	Send(w, resBody, statusCode)
}
