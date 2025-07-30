package creemio

import (
	"net/http"
	"net/url"
)

type Response struct {
	RequestURL *url.URL
	Status     int
	Headers    http.Header
}

func newResponse(res *http.Response) *Response {
	return &Response{
		RequestURL: res.Request.URL,
		Status:     res.StatusCode,
		Headers:    res.Header,
	}
}
