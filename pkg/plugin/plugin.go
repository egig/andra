package plugin

import (
	"plugin"
)

type Base interface {
	Name() string
}

type Plugin struct {
}

func (p *Plugin) Name() string  {
	return ""
}

func Load(file string) {
	_, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}
}
