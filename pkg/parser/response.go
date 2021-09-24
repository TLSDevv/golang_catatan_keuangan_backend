// package parser

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"strconv"

// 	pkgerr "heltek-go/pkg/error"
// )

// const (
// 	successCode = "0000"
// 	errorCode   = "9999"
// )

// type Response struct {
// 	Code string      `json:"code"`
// 	Msg  string      `json:"message"`
// 	Data interface{} `json:"data"`
// 	Meta *Meta       `json:"meta,omitempty"`
// }

// type Meta struct {
// 	Next     string `json:"next"`
// 	Current  string `json:"current"`
// 	Previous string `json:"previous"`
// }

// type JSONResponder interface {
// 	Write(w http.ResponseWriter, status int, data interface{})
// 	SuccessNoData(w http.ResponseWriter, status int, location string)
// 	Error(w http.ResponseWriter, apiErr error)
// 	SuccessWithData(w http.ResponseWriter, msg string, data interface{})
// 	SuccessWithDataPagination(w http.ResponseWriter, msg string, data interface{}, meta *Meta)
// 	PreparePaginationMeta(r *http.Request, totalRetrieved, pageNum, pageRow int) *Meta
// }

// type jsonResponder struct {
// 	contentType string
// }

// func NewJSONResponder() JSONResponder {
// 	return &jsonResponder{
// 		contentType: fmt.Sprintf("%s; charset=%s", jsonContentType, jsonCharset),
// 	}
// }

// func (c *jsonResponder) Write(w http.ResponseWriter, status int, data interface{}) {
// 	w.Header().Set("Content-Type", c.contentType)
// 	w.WriteHeader(status)
// 	if data == nil {
// 		return
// 	}

// 	content, _ := json.Marshal(data)
// 	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
// 	_, _ = w.Write(content)
// }

// func (c *jsonResponder) SuccessWithData(w http.ResponseWriter, msg string, data interface{}) {
// 	content := Response{
// 		Code: successCode,
// 		Msg:  msg,
// 		Data: data,
// 	}
// 	c.Write(w, http.StatusOK, content)
// }

// func (c *jsonResponder) SuccessWithDataPagination(w http.ResponseWriter, msg string, data interface{}, meta *Meta) {
// 	content := Response{
// 		Code: successCode,
// 		Msg:  msg,
// 		Data: data,
// 		Meta: meta,
// 	}
// 	c.Write(w, http.StatusOK, content)
// }

// func (c *jsonResponder) SuccessNoData(w http.ResponseWriter, status int, location string) {
// 	if location != "" {
// 		w.Header().Set("Location", location)
// 	}
// 	c.Write(w, status, nil)
// }

// func (c *jsonResponder) Error(w http.ResponseWriter, err error) {
// 	apiErr, ok := err.(pkgerr.APIError)
// 	if ok {
// 		c.Write(w, apiErr.StatusCode, Response{
// 			Code: apiErr.Code,
// 			Msg:  apiErr.Message,
// 		})
// 	} else {
// 		c.Write(w, http.StatusInternalServerError, Response{
// 			Code: errorCode,
// 			Msg:  "We failed to process your request",
// 		})
// 	}
// }

// func (c *jsonResponder) PreparePaginationMeta(r *http.Request, totalData, pageNum, pageRow int) *Meta {
// 	newURL, err := url.Parse(r.URL.Host)
// 	if err != nil {
// 		return nil
// 	}
// 	newURL.Path = r.URL.Path

// 	newURL.RawQuery = r.URL.Query().Encode()

// 	meta := &Meta{}
// 	meta.Current = c.defineNewURL(*newURL, pageNum, pageRow)

// 	if totalData != 0 {
// 		if totalData == pageRow {
// 			// return empty when nothing to show next paging
// 			meta.Next = c.defineNewURL(*newURL, pageNum+1, pageRow)
// 		}
// 		if pageNum > 1 {
// 			meta.Previous = c.defineNewURL(*newURL, pageNum-1, pageRow)
// 		}
// 	}

// 	return meta
// }

// func (c *jsonResponder) defineNewURL(n url.URL, page, perPage int) string {
// 	v := url.Values{}
// 	v = n.Query()
// 	v.Set("page", strconv.Itoa(page))
// 	v.Set("per_page", strconv.Itoa(perPage))

// 	n.RawQuery = v.Encode()

// 	return n.String()
// }
