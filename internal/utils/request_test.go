package utils

import (
	"net/http"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromRequest(t *testing.T) {
	tests := []struct {
		name string
		in   js.Value
		out  *http.Request
	}{
		{
			name: "empty",
			in:   RequestClass.New("http://localhost:8080/"),
			out:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := FromRequest(tt.in)
			assert.NoError(t, err)
			assert.NotNil(t, r)
		})
	}
}

func TestRequestToJS(t *testing.T) {
	tests := []struct {
		name string
		in   *http.Request
		out  js.Value
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RequestToJS(tt.in)
			assert.Equal(t, tt.out, r)
		})
	}
}
