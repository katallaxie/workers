package utils

import (
	"net/http"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJS(t *testing.T) {
	tests := []struct {
		name string
		h    http.Header
		want js.Value
	}{
		{
			name: "empty",
			h:    http.Header{},
			want: HeadersClass.New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromHeaderToJS(tt.h)
			assert.NotNil(t, got)

			entries := ArrayFrom(got.Call("entries"))
			assert.Equal(t, len(tt.h), entries.Length())
		})
	}
}

func TestFromJS(t *testing.T) {
	tests := []struct {
		name    string
		headers js.Value
		want    http.Header
	}{
		{
			name:    "empty",
			headers: HeadersClass.New(),
			want:    http.Header{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromJS(tt.headers)
			assert.Equal(t, tt.want, got)
		})
	}
}
