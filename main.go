package main

import (
	"errors"
	"os"

	"github.com/mohammadne/zar/cmd"
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
		PersistentPreRunE:     isRoot,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
	}

	root.AddCommand(
		cmd.Run{}.Command(cfg),
	)

	if err := root.Execute(); err != nil {
		logger := log.NewZap(cfg.Log)
		logger.Fatal("failed to execute root command", zap.Error(err))
	}
}

// isRoot implements a cobra acceptable function and
// returns ErrNotPermitted if user is not root.
func isRoot(_ *cobra.Command, _ []string) error {
	if os.Getuid() != 0 {
		return errors.New("operation not permitted, you should be the root user")
	}
	return nil
}
