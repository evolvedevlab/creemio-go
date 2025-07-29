package creemio

import "net/http"

type Response struct {
	Status  int
	Headers http.Header
}

func newResponse(res *http.Response) *Response {
	return &Response{Status: res.StatusCode, Headers: res.Header}
}
