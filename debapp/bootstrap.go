package debapp

import (
	"fmt"

	"github.com/debrootfs/util"
	"github.com/debrootfs/util/bootstrap"
	"github.com/debrootfs/util/system"
)

type BootstrapError struct {
	Msg string
}

func (e *BootstrapError) Error() string {
	return fmt.Sprintf("Bootstrap Error: %s", e.Msg)
}

func BuildBootstrap() error {
	bootstrap := bootstrap.DebBootstrap{}
	exist, err := util.PathExists("./debroot")

	if err != nil {
		return err
	}

	if !exist {
		if err := bootstrap.Create("amd64", "bookworm", "./debroot"); err != nil {
			return err
		}
	}

	return nil
}

func MountDefaultFS() error {
	e := system.CreateLocalEnv("/bin/bash")
	if e == nil {
		return &BootstrapError{
			Msg: "Failed to Create Virtual Environment",
		}
	}

	if err := e.Start(); err != nil {
		panic(err)
	}

	ch, finCh := e.CretaeChannel()

	go func() {
		ch <- system.VirtualEnvCmd{
			Cmds: []string{
				"sudo mount --bind /dev ./debroot/dev",
				"sudo mount -t proc proc ./debroot/proc",
				"sudo mount -t sysfs sys ./debroot/sys",
				"sudo mount -t tmpfs tmpfs ./debroot/run",
				"exit",
			},
		}
	}()

	if err := <-finCh; err != nil {
		fmt.Println("Test Error: ", err)
	}

	return nil
}
