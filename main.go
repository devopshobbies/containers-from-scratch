package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammadne/zar/cmd"
	"github.com/mohammadne/zar/internal/config"
	"github.com/mohammadne/zar/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	short = "short description"
	long  = `long description`
)

func main() {
	cfg := config.Load()

	logger := log.NewZap(cfg.Log)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	root := &cobra.Command{Short: short, Long: long}
	root.AddCommand(
		cmd.Run{}.Command(cfg, channel),
	)

	if err := root.Execute(); err != nil {
		logger.Fatal("failed to execute root command", zap.Error(err))
	}
}
