package di

import (
	"go.uber.org/dig"
)

type (
	appIn struct {
		dig.In

		Applications []Application `group:"applications"`
	}

	appOut struct {
		dig.Out

		Application ApplicationPool
	}
)

func appAdapter(in appIn) appOut {
	return appOut{
		Application: in.Applications,
	}
}
