package workers

import "context"

// GetEnv gets the value of the environment variable named by the key.
func GetEnv(ctx context.Context, key string) string {
	return GetEnvContext(ctx).Get(key).String()
}

// GetBinding gets the value of the binding variable named by the key.
func GetBinding(ctx context.Context, key string) string {
	return GetExecutionContext(ctx).Get(key)
}
