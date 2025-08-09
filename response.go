package creemio

import (
	"net/http"
	"net/url"
)

type Response struct {
	RequestURL *url.URL
	Status     int
	Headers    http.Header
	Body       []byte
}

func newResponse(r *http.Response, body []byte) *Response {
	return &Response{
		RequestURL: r.Request.URL,
		Status:     r.StatusCode,
		Headers:    r.Header,
		Body:       body,
	}
}
