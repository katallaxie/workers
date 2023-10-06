package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewObject(t *testing.T) {
	obj := NewObject()
	assert.NotNil(t, obj)
}

func TestNewUint8Array(t *testing.T) {
	arr := NewUint8Array(10)
	assert.NotNil(t, arr)
	assert.Equal(t, 10, arr.Length())
}
