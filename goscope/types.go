// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

const (
	LogDashboardMode = iota
	RequestDashboardMode
)

type ExceptionRecord struct {
	Error string `json:"error"`
	Time  int    `json:"time"`
	UID   string `json:"uid"`
}

type SummarizedRequest struct {
	Method         string `json:"method"`
	Path           string `json:"path"`
	Time           int    `json:"time"`
	UID            string `json:"uid"`
	ResponseStatus int    `json:"response_status"`
}

type RecordByURI struct {
	UID string `uri:"id" binding:"required"`
}

type SummarizedResponse struct {
	RequestUID string `json:"request_uid"`
	ClientIP   string `json:"client_ip"`
	Path       string `json:"path"`
	Status     string `json:"status"`
	Time       int    `json:"time"`
	UID        string `json:"uid"`
}

type DetailedResponse struct {
	Body       string `json:"body"`
	ClientIP   string `json:"client_ip"`
	Headers    string `json:"headers"`
	Path       string `json:"path"`
	Size       int    `json:"size"`
	Status     string `json:"status"`
	Time       int    `json:"time"`
	RequestUID string `json:"request_uid"`
	UID        string `json:"uid"`
}

type DetailedRequest struct {
	Body      string `json:"body"`
	ClientIP  string `json:"client_ip"`
	Headers   string `json:"headers"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Referrer  string `json:"referrer"`
	Time      int    `json:"time"`
	UID       string `json:"uid"`
	URL       string `json:"url"`
	UserAgent string `json:""`
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// HTTP request body object.
func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type SearchRequestPayload struct {
	Query string `json:"query"`
}
