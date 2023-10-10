package main

import (
	"context"
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
)

type config struct {
	Verbose bool
}

func (c *config) Cwd() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func defaultConfig() *config {
	return &config{
		Verbose: false,
	}
}

var cfg = defaultConfig()

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "verbose output")

	rootCmd.SilenceErrors = true
}

var rootCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func runRoot(ctx context.Context) error {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"cmd/build/src/worker.ts"},
		Outdir:      "dist",
		Bundle:      true,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
	})

	if len(result.Errors) > 0 {
		return fmt.Errorf("failed to build: %v", result.Errors)
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
