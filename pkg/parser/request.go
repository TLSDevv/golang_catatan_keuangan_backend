package parser

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	jsonContentType = "application/json"
	multipartType   = "multipart/data"
	jsonCharset     = "utf-8"
)

type RequestParser interface {
	ParseJSON(r *http.Request, i interface{}) error
	ParseForm(r *http.Request, i interface{}) error
}

type requestParser struct{}

func NewRequestParser() RequestParser {
	return &requestParser{}
}

func (req *requestParser) ParseJSON(r *http.Request, body interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, jsonContentType) {
		return json.NewDecoder(r.Body).Decode(&body)
	}

	return errors.New("no supported type")
}

func (req *requestParser) ParseForm(r *http.Request, body interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, multipartType) {
		return json.NewDecoder(r.Body).Decode(&body)
	}

	return errors.New("no supported type")
}
