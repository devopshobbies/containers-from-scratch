//go:build linux
// +build linux

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()

	default:
		panic("help")
	}
}

func run() {
	log.Printf("Running %v \n", os.Args[2:])

	//nolint:gosec // we will panic if user arguments are invalid.
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	log.Printf("Running %v as %d\n", os.Args[2:], os.Getegid())

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/tmp/alpine-rootfs/"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	//nolint:gosec // we will panic if user arguments are invalid.
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())

	must(syscall.Unmount("/proc", 0))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
