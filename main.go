//go:build linux
// +build linux

package main

import (
	"errors"
	"os"

	"github.com/mohammadne/zar/cmd"
	"github.com/mohammadne/zar/internal"
	"github.com/mohammadne/zar/internal/config"
	"github.com/mohammadne/zar/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

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
	os.MkdirAll(internal.LayersPath, 0700)
	os.MkdirAll(internal.ContainersPath, 0700)
	os.MkdirAll(internal.NetNSPath, 0700)

	return nil
}
