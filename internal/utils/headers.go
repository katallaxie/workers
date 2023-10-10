package utils

import (
	"net/http"
	"net/textproto"
	"strings"
	"syscall/js"
)

// Headers is a Go representation of a JavaScript Headers object.
type Headers http.Header

// ToJS converts a Go http.Header to a JavaScript Headers object.
func (h Headers) ToJS() js.Value {
	headers := HeadersClass.New()

	for k, vv := range h {
		for _, v := range vv {
			headers.Call("append", k, v)
		}
	}

	return headers
}

// Add adds the key, value pair to the header.
func (h Headers) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

// Set sets the key, value pair to the header.
func (h Headers) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

// FromJS converts a JavaScript Headers object to a Go http.Header.
func FromJS(headers js.Value) Headers {
	entries := ArrayFrom(headers.Call("entries"))
	h := make(Headers)

	for i := 0; i < entries.Length(); i++ {
		e := entries.Index(i)
		k := e.Index(0).String()
		v := e.Index(1).String()
		for _, value := range strings.Split(v, ",") {
			h.Add(k, value)
		}
	}

	return h
}
