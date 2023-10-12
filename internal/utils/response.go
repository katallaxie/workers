package utils

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
)

// ResponseWriter ...
type ResponseWriter struct {
	StatusCode int

	Reader *io.PipeReader
	Writer *io.PipeWriter

	header http.Header
	ready  chan struct{}
	once   sync.Once
}

// NewResponseWriter ...
func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{}
}

// Ready ...
func (w *ResponseWriter) Ready() {
	w.once.Do(func() {
		close(w.ready)
	})
}

// Write ...
func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Ready()

	return w.Writer.Write(b)
}

// Header ...
func (w *ResponseWriter) Header() http.Header {
	return w.header
}

// SetHeader ...
func (w *ResponseWriter) SetHeader(key, value string) {
	w.header.Set(key, value)
}

// WriteHeader ...
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}

// ToJS ...
func (w *ResponseWriter) ToJS() js.Value {
	return ResponseToJS(newJSResponse(w.StatusCode, w.header, w.Reader))
}

// FromResponse ...
func FromResponse(r js.Value) (*http.Response, error) {
	status := r.Get("status").Int()
	promise := r.Call("text")

	body, err := Await(promise)
	if err != nil {
		return nil, err
	}

	header := FromJS(r.Get("headers"))
	contentLength, _ := strconv.ParseInt(header.Get("Content-Length"), 10, 64)

	res := &http.Response{
		Status:        strconv.Itoa(status) + " " + r.Get("statusText").String(),
		StatusCode:    status,
		Header:        header,
		Body:          io.NopCloser(strings.NewReader(body.String())),
		ContentLength: contentLength,
	}

	return res, nil
}

// ResponseToJS ...
func ResponseToJS(r *http.Response) js.Value {
	return newJSResponse(r.StatusCode, r.Header, r.Body)
}

func newJSResponse(status int, headers http.Header, body io.ReadCloser) js.Value {
	status := r.StatusCode

	if status == 0 {
		status = http.StatusOK
	}

	res := NewObject()
	res.Set("status", status)
	res.Set("statusText", http.StatusText(status))
	res.Set("headers", FromHeaderToJS(r.Header))

	if status == http.StatusSwitchingProtocols ||
		status == http.StatusNoContent ||
		status == http.StatusResetContent ||
		status == http.StatusNotModified {
		return ResponseClass.New(Null, res)
	}

	readableStream := jsutil.ConvertReaderToReadableStream(body)

	return ResponseClass.New(readableStream, res)
}
