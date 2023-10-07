package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithRuntimeContext(t *testing.T) {
	context := WithRuntimeContext(context.TODO(), ObjectClass.New())
	assert.NotNil(t, context)
}

func TestGetRuntimeContext(t *testing.T) {
	context := WithRuntimeContext(context.TODO(), ObjectClass.New())
	runtimeContext := GetRuntimeContext(context)
	assert.NotNil(t, runtimeContext)
}
