package main

import (
	"context"
	"embed"
	"log"
	"os"
	"path"

	"github.com/katallaxie/pkg/filex"
	"github.com/spf13/cobra"
)

//go:embed shims wasm
var fs embed.FS

const (
	shimsDir   = "shims"
	wasmDir    = "wasm"
	entryPoint = "wasm_exec.js"
)

type cfg struct {
	Verbose    bool
	Path       string
	EntryPoint string
}

var config = cfg{
	Verbose:    false,
	Path:       "build",
	EntryPoint: "wasm_exec_go.js",
}

var rootCmd = &cobra.Command{
	Use: "gen",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&config.EntryPoint, "entrypoint", "e", config.EntryPoint, "entrypoint to use")
	rootCmd.PersistentFlags().StringVarP(&config.Path, "path", "p", config.Path, "output path")

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
}

func runRoot(_ context.Context) error {
	err := filex.MkdirAll(config.Path, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = filex.CopyEmbedFile(path.Join(wasmDir, config.EntryPoint), path.Join(config.Path, entryPoint), fs)
	if err != nil {
		return err
	}

	// copy shims
	entries, err := fs.ReadDir(shimsDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		_, err = filex.CopyEmbedFile(path.Join(shimsDir, entry.Name()), path.Join(config.Path, entry.Name()), fs)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
