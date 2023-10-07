package main

import (
	"context"
	"os"

	"github.com/katallaxie/pkg/utils/files"
	"github.com/spf13/cobra"
)

type config struct {
	TinyGo    bool
	Verbose   bool
	AssetDir  string
	CommonDir string
	OutDir    string
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
		TinyGo:    true,
		Verbose:   false,
		AssetDir:  "assets",
		CommonDir: "assets/common",
		OutDir:    "build",
	}
}

var cfg = defaultConfig()

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&cfg.TinyGo, "tinygo", cfg.TinyGo, "use tinygo compiler")

	rootCmd.SilenceErrors = true
}

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func runRoot(ctx context.Context) error {
	err := files.Clean(cfg.OutDir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
