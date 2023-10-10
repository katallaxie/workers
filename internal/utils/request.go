package utils

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
)

// FromRequest returns a new Request from a JavaScript Request.
func FromRequest(r js.Value) (*http.Request, error) {
	url, err := url.Parse(r.Get("url").String())
	if err != nil {
		return nil, err
	}

	h := FromJS(r.Get("headers"))
	contentLength, _ := strconv.ParseInt(h.Get("Content-Length"), 10, 64)

	return &http.Request{
		Method:           r.Get("method").String(),
		URL:              url,
		Header:           h,
		ContentLength:    contentLength,
		TransferEncoding: strings.Split(h.Get("Transfer-Encoding"), ","),
		Host:             h.Get("Host"),
	}, nil
}

// RequestToJS converts a Go http.Request to a JavaScript Request object.
func RequestToJS(r *http.Request) js.Value {
	opts := NewObject()
	opts.Set("method", r.Method)
	opts.Set("url", r.URL.String())
	opts.Set("headers", FromHeaderToJS(r.Header))

	body := Undefined
	opts.Set("body", body)

	return RequestClass.New(r.URL.String(), opts)
}
