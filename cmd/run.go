package cmd

import (
	"fmt"

	"github.com/mohammadne/zar/internal/config"
	"github.com/spf13/cobra"
)

type Run struct {
	config *config.Config
}

func (run Run) Command(cfg *config.Config) *cobra.Command {
	run.config = cfg

	cmd := &cobra.Command{
		Use:                   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short:                 "run a command inside a new container",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:                  cobra.ExactArgs(1),
		Run:                   run.main,
	}

	flags := cmd.Flags()
	flags.String("host", "", "Container Hostname")
	flags.IntVar(&cfg.CGroups.Memory, "memory", cfg.CGroups.Memory, "Limit memory access in MB")
	flags.Float64Var(&cfg.CGroups.CPUs, "cpu", cfg.CGroups.CPUs, "Limit CPUs")
	flags.IntVar(&cfg.CGroups.Swap, "swap", cfg.CGroups.Swap, "Limit swap access in MB")
	flags.IntVar(&cfg.CGroups.PIDs, "pids", cfg.CGroups.PIDs, "Limit number of processes")
	flags.Bool("detach", false, "run command in the background")

	return cmd
}

func (run *Run) main(cmd *cobra.Command, args []string) {
	val, _ := cmd.Flags().GetInt("memory")
	fmt.Println("ARGS", val)

	fmt.Println(run.config.CGroups.Memory)
}
