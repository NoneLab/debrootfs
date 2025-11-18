package bootstrap

type DebBootstrap struct {
	OutPath string
}

type DebBootstrapInterface interface {
	Create(string)
}

func (d DebBootstrap) Create(path string) {
	d.OutPath = path
}
