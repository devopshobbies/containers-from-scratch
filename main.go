package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammadne/zar/cmd"
	"github.com/mohammadne/zar/internal/config"
	"github.com/mohammadne/zar/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger := log.NewZap(cfg.Log)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	root := &cobra.Command{
		Use:                   "zar [OPTIONS] COMMAND",
		Short:                 "A tiny tool for managing containers",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PersistentPreRunE:     isRoot,
	}

	root.AddCommand(
		cmd.Run{}.Command(cfg, channel),
	)

	if err := root.Execute(); err != nil {
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
