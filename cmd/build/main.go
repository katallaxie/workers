package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"cmd/build/src/worker.ts"},
		Outdir:      "dist",
		Bundle:      false,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
