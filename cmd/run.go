package cmd

import (
	"fmt"

	"github.com/mohammadne/zar/internal/config"
	"github.com/spf13/cobra"
)

type Run struct{}

func (run Run) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short:                 "run a command inside a new container",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:                  cobra.ExactArgs(1),
		Run:                   run.main,
	}

	cfg := config.Load()

	flags := cmd.Flags()
	flags.StringP("host", "", "", "Container Hostname")
	flags.IntP("memory", "m", cfg.CGroups.Memory, "Limit memory access in MB")
	flags.IntP("swap", "s", 20, "Limit swap access in MB")
	flags.Float64P("cpu", "c", cfg.CGroups.CPU, "Limit CPUs")
	flags.IntP("pids", "p", 128, "Limit number of processes")
	flags.BoolP("detach", "d", false, "run command in the background")

	return cmd
}

func (run *Run) main(cmd *cobra.Command, args []string) {
	cfg := config.Load()
	_ = cfg
	val, _ := cmd.Flags().GetInt("memory")
	fmt.Println("ARGS", val)
}
