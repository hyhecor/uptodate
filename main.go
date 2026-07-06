package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	binaryPath string
	exts       string
	all        bool

	rootCMD = cobra.Command{
		Use:   "uptodate <source-root>",
		Short: "Check whether a binary is up to date",
		Long: `uptodate checks whether the given binary is up to date
with respect to the source files under source-root.

It walks all files under source-root and compares their
modification time against the binary.

Exit code:
  0 - binary is up to date
  1 - rebuild is required
  2 - error`,
		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			rootDir := args[0]

			allowed := map[string]bool{}
			if exts != "" {
				for _, e := range strings.Split(exts, ",") {
					e = strings.TrimSpace(e)
					if !strings.HasPrefix(e, ".") {
						e = "." + e
					}
					allowed[e] = true
				}
			}

			binInfo, err := os.Stat(binaryPath)
			if err != nil {
				slog.Error("failed to stat binary", "error", err)
				os.Exit(2)
			}

			binTime := binInfo.ModTime()
			needsRebuild := false

			err = filepath.Walk(rootDir, func(path string, file fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if file.IsDir() {
					return nil
				}

				if !all && len(allowed) > 0 {
					ext := filepath.Ext(file.Name())
					if !allowed[ext] {
						return nil
					}
				}

				if file.ModTime().After(binTime) {
					needsRebuild = true
					return filepath.SkipAll
				}

				return nil
			})

			if err != nil {
				slog.Error("walk failed", "error", err)
				os.Exit(2)
			}

			if needsRebuild {
				fmt.Println("rebuild is required")
				os.Exit(1)
			}

			fmt.Println("up to date")
			os.Exit(0)
		},
	}
)

func init() {
	rootCMD.Flags().StringVarP(&binaryPath, "binary", "b", "", "Path to the binary to compare against")

	rootCMD.Flags().StringVarP(&exts, "ext", "e", exts, "Comma-separated file extensions to include. If not set, all files are considered. (e.g. go,mod,sum)")

	rootCMD.Flags().BoolVar(&all, "all", all, "Include all files (override ext filter)")

	_ = rootCMD.MarkFlagRequired("binary")
}

func main() {
	_ = rootCMD.Execute()
}
