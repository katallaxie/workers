package utils

import (
	"net/http"
	"strings"
	"syscall/js"
)

// ToJS converts a Go http.Header to a JavaScript Headers object.
func FromHeaderToJS(h http.Header) js.Value {
	headers := HeadersClass.New()

	for k, vv := range h {
		for _, v := range vv {
			headers.Call("append", k, v)
		}
	}

	return headers
}

// FromJS converts a JavaScript Headers object to a Go http.Header.
func FromJS(headers js.Value) http.Header {
	entries := ArrayFrom(headers.Call("entries"))
	h := make(http.Header)

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
