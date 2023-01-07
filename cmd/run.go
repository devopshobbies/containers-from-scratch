package cmd

import (
	"fmt"

	"github.com/mohammadne/zar/internal/config"
	"github.com/mohammadne/zar/pkg/utils"
	"github.com/spf13/cobra"
)

type Run struct {
	config  *config.Config
	command string
}

func (run Run) Command(cfg *config.Config) *cobra.Command {
	run.config = cfg

	cmd := &cobra.Command{
		Use:                   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short:                 "run a command inside a new container",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:                  run.validateArgs,
		Run:                   run.main,
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.Hostname, "hostname", utils.RandomString(10), "container hostname")
	flags.IntVar(&cfg.CGroups.Memory, "memory", cfg.CGroups.Memory, "limit memory access in MB")
	flags.Float64Var(&cfg.CGroups.CPUs, "cpu", cfg.CGroups.CPUs, "limit CPUs")
	flags.IntVar(&cfg.CGroups.Swap, "swap", cfg.CGroups.Swap, "limit swap access in MB")
	flags.IntVar(&cfg.CGroups.PIDs, "pids", cfg.CGroups.PIDs, "limit number of processes")
	flags.Bool("detach", false, "run command in the background")

	return cmd
}

func (run *Run) validateArgs(_ *cobra.Command, args []string) error {
	if length := len(args); length < 1 {
		return fmt.Errorf("error, no command has been provided")
	} else if length > 1 {
		return fmt.Errorf("error, more than one command has been provided, commands: %v", args)
	}

	run.command = args[0]
	return nil
}

func (run *Run) main(cmd *cobra.Command, _ []string) {
	fmt.Println(run.command)

	val, _ := cmd.Flags().GetBool("detach")
	fmt.Println("ARGS", val)

	fmt.Println(run.config.CGroups.Memory)
}
