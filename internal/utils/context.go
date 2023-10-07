package utils

import (
	"context"
	"syscall/js"
)

type runtimeCtxKey struct{}

// WithRuntimeContext returns a new context with the given runtime context.
func WithRuntimeContext(ctx context.Context, runtimeContext js.Value) context.Context {
	return context.WithValue(ctx, runtimeCtxKey{}, runtimeContext)
}

// GetRuntimeContext returns the runtime context from the given context.
func GetRuntimeContext(ctx context.Context) js.Value {
	runtimeCtx, ok := ctx.Value(runtimeCtxKey{}).(js.Value)
	if !ok {
		return Null
	}

	return runtimeCtx
}
