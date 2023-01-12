//go:build linux
// +build linux

package main

import (
	"errors"
	"github.com/devopshobbies/containers-from-scratch/cmd"
	"github.com/devopshobbies/containers-from-scratch/internal"
	"github.com/devopshobbies/containers-from-scratch/internal/config"
	"github.com/devopshobbies/containers-from-scratch/pkg/log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const Perm = 0700

func main() {
	cfg := config.Load()

	root := &cobra.Command{
		Use:                   "zar [OPTIONS] COMMAND",
		Short:                 "A tiny tool for managing containers",
		PersistentPreRunE:     preRun,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
	}

	root.AddCommand(
		cmd.Run{}.Command(cfg),
	)

	if err := root.Execute(); err != nil {
		logger := log.NewZap(cfg.Log)
		logger.Fatal("failed to execute root command", zap.Error(err))
	}
}

func preRun(_ *cobra.Command, _ []string) error {
	// returns ErrNotPermitted if user is not root
	if os.Getuid() != 0 {
		return errors.New("operation not permitted, you should be the root user")
	}

	// create necessary directories
	return createDirs()
}

func createDirs() error {
	dirs := []string{internal.LayersPath, internal.ContainersPath,
		internal.NetNSPath}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, Perm); err != nil {
			return err
		}
	}

	return nil
}
