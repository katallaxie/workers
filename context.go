package workers

import (
	"context"
	"syscall/js"

	"github.com/katallaxie/workers/internal/utils"
)

// GetEnvContext returns the value of the environment variable named by the key.
// - see: https://github.com/cloudflare/workers-types/blob/c8d9533caa4415c2156d2cf1daca75289d01ae70/index.d.ts#L566
func GetEnvContext(ctx context.Context) js.Value {
	return utils.GetRuntimeContext(ctx).Get("env")
}

// GetExecutionContext returns the execution context of the current script.
func GetExecutionContext(ctx context.Context) js.Value {
	return utils.GetRuntimeContext(ctx).Get("ctx")
}
