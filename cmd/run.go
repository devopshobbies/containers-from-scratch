package cmd

import (
	"fmt"
	"os"

	"github.com/mohammadne/zar/internal/config"
	"github.com/spf13/cobra"
)

type Run struct{}

func (cmd Run) Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { cmd.main(cfg, trap) }
	return &cobra.Command{Use: "run", Short: "run heliograph call server", Run: run}
}

func (cmd *Run) main(cfg *config.Config, trap chan os.Signal) {
	fmt.Println("HELLO")
}
