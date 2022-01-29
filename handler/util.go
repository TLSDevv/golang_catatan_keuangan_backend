package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ResponseBody struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func GetParams(r *http.Request, params string) int {
	paramString, _ := strconv.ParseInt(mux.Vars(r)[params], 10, 64)
	return int(paramString)
}

func Decode(r *http.Request, toBody interface{}) error {
	return json.NewDecoder(r.Body).Decode(toBody)
}

func Send(w http.ResponseWriter, responseBody interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return Encode(w, responseBody)
}

func Encode(w http.ResponseWriter, responseBody interface{}) error {
	return json.NewEncoder(w).Encode(responseBody)
}

func SendWithData(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resBody := ResponseBody{
		Code:   statusCode,
		Status: message,
		Data:   data,
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
