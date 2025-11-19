package system

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type VirtualEnv struct {
	cmd   *exec.Cmd
	stdin io.WriteCloser
}

func initShellIO(s *io.WriteCloser, v *exec.Cmd) error {
	stdin, err := v.StdinPipe()
	if err != nil {
		return err
	}

	s = &stdin
	v.Stdout = os.Stdout
	v.Stderr = os.Stderr

	return nil
}

func CreateVirtualEnv(shell string, chroot_path string) *VirtualEnv {
	if shell == "" {
		return nil
	}

	ret := &VirtualEnv{
		cmd: exec.Command(shell),
	}

	ret.cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot: chroot_path,
	}

	if err := initShellIO(&ret.stdin, ret.cmd); err != nil {
		panic(err)
	}

	return ret
}

type VirtualEnvInterface interface {
	Start()
	Done()
	Send(func())
	Wait() error
}

func (v *VirtualEnv) Start() error {
	return v.cmd.Start()
}

func (v *VirtualEnv) Send(command string) {
	fmt.Fprintln(v.stdin, command)
}
