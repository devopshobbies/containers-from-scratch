package cmd

import (
	"fmt"

	"github.com/devopshobbies/containers-from-scratch/internal/config"
	"github.com/devopshobbies/containers-from-scratch/pkg/random"
	"github.com/spf13/cobra"
)

type Run struct {
	config   *config.Config
	commands []string
}

func (run Run) Command(cfg *config.Config) *cobra.Command {
	run.config = cfg

	cmd := &cobra.Command{
		Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short: "run a command inside a new container",
		Args:  run.validateArgs,
		Run:   run.main,
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.Hostname, "hostname", random.RandomString(10), "container hostname")
	flags.IntVar(&cfg.CGroup.Memory, "memory", cfg.CGroup.Memory, "limit memory access in MB")
	flags.Float64Var(&cfg.CGroup.CPUs, "cpu", cfg.CGroup.CPUs, "limit CPUs")
	flags.IntVar(&cfg.CGroup.Swap, "swap", cfg.CGroup.Swap, "limit swap access in MB")
	flags.IntVar(&cfg.CGroup.PIDs, "pids", cfg.CGroup.PIDs, "limit number of processes")
	flags.Bool("detach", false, "run command in the background")

	return cmd
}

func (run *Run) validateArgs(_ *cobra.Command, args []string) error {
	if length := len(args); length < 1 {
		return fmt.Errorf("error, no command has been provided")
	}

	run.commands = args
	return nil
}

func (run *Run) main(_ *cobra.Command, _ []string) {
	fmt.Println(run.commands)

	fmt.Println(run.config.Hostname)
}
