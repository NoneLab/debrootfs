package bootstrap

import (
	"fmt"
	"os"
	"os/exec"
)

type DebBootstrap struct {
	OutPath string
}

type DebBootstrapInterface interface {
	Create(string, string, string)
}

func debootstrapVerbose(arch string, deb_ver string, path string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("Debian Version : " + deb_ver)
	fmt.Println("Arch : " + arch)
	fmt.Println("Output Path : " + path)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func (d *DebBootstrap) Create(arch string, deb_ver string, path string) error {
	debootstrapVerbose(arch, deb_ver, path)

	d.OutPath = path
	arch = "--arch=" + arch
	cmd := exec.Command("sudo", "debootstrap", arch, deb_ver, path, "http://deb.debian.org/debian")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
