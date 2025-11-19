package system

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type VirtualEnvCmd struct {
	Cmds []string
}

type VirtualEnv struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	cmd_ch chan<- VirtualEnvCmd
}

func initShellIO(s *io.WriteCloser, v *exec.Cmd) error {
	stdin, err := v.StdinPipe()
	if err != nil {
		return err
	}

	*s = stdin
	// v.Stdin = os.Stdin
	v.Stdout = os.Stdout
	v.Stderr = os.Stderr

	return nil
}

func CreateVirtualEnvBase(shell string, chroot_path string, flag bool) *VirtualEnv {
	if shell == "" {
		return nil
	}

	ret := &VirtualEnv{
		cmd: exec.Command(shell),
	}

	ret.cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot: chroot_path,
	}

	if flag {
		ret.cmd.Dir = "/"
	}

	if err := initShellIO(&ret.stdin, ret.cmd); err != nil {
		panic(err)
	}

	return ret
}

func CreateLocalEnv(shell string) *VirtualEnv {
	return CreateVirtualEnvBase(shell, "", false)
}

func CreateVirtualEnv(shell string, chroot_path string) *VirtualEnv {
	return CreateVirtualEnvBase(shell, chroot_path, true)
}

type VirtualEnvInterface interface {
	Start()
	Done()
	CretaeChannel() chan<- VirtualEnvCmd
	CloseChannel()
	Wait() error
}

func (v *VirtualEnv) Start() error {
	return v.cmd.Start()
}

func (v *VirtualEnv) CretaeChannel() (chan<- VirtualEnvCmd, <-chan error) {
	ch := make(chan VirtualEnvCmd)
	finCh := make(chan error, 1)

	go func() {
		defer v.stdin.Close()
		for c := range ch {
			for _, cmd := range c.Cmds {
				fmt.Fprintln(v.stdin, cmd)
			}
		}
	}()

	go func() {
		err := v.Wait()
		finCh <- err
		close(finCh)
		close(ch)
	}()

	v.cmd_ch = ch

	return ch, finCh
}

func (v *VirtualEnv) CloseChannel() {
	close(v.cmd_ch)
}

func (v *VirtualEnv) Wait() error {
	return v.cmd.Wait()
}

func (v *VirtualEnv) Send(command string) {
	fmt.Fprintln(v.stdin, command)
}
