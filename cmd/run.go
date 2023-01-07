package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"honnef.co/go/tools/config"
)

type Run struct{}

func (cmd Run) Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { cmd.main(cfg, trap) }
	return &cobra.Command{Use: "run", Short: "run heliograph call server", Run: run}
}

func (cmd *Run) main(cfg *config.Config, trap chan os.Signal) {}
