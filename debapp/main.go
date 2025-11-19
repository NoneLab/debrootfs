package debapp

import (
	"fmt"

	"github.com/debrootfs/util/bootstrap"
)

func Main() {
	fmt.Println("Hello makedebrootfs")
	fmt.Println("This is go")

	bootstrap := bootstrap.DebBootstrap{}
	if err := bootstrap.Create("amd64", "bookworm", "./debroot"); err != nil {
		panic(err)
	}
}
