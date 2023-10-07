package main

import (
	"context"

	"github.com/spf13/cobra"
)

type config struct {
	TinyGo  bool
	Verbose bool
}

func defaultConfig() *config {
	return &config{
		TinyGo:  true,
		Verbose: false,
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
	return nil
}

func main() {

}
