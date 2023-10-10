package utils

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJS(t *testing.T) {
	tests := []struct {
		name string
		h    Headers
		want js.Value
	}{
		{
			name: "empty",
			h:    Headers{},
			want: HeadersClass.New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.h.ToJS()
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
		want    Headers
	}{
		{
			name:    "empty",
			headers: HeadersClass.New(),
			want:    Headers{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromJS(tt.headers)
			assert.Equal(t, tt.want, got)
		})
	}
}
