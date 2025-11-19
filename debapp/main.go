package debapp

import (
	"fmt"
)

func Main() {
	fmt.Println("Hello makedebrootfs")
	fmt.Println("This is go")

	//  Get Debian bootstrap (rootfs)
	if err := BuildBootstrap(); err != nil {
		panic(err)
	}

	if err := MountDefaultFS(); err != nil {
		panic(err)
	}
}
