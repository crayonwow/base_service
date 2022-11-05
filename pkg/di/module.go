package di

import (
	"go.uber.org/dig"
)

type (
	dependency struct {
		constructor interface{}
		ops         []dig.ProvideOption
	}

	Module []dependency

	Container struct {
		c *dig.Container
	}
)

func (m Module) Append(n Module) Module {
	return append(m, n...)
}

func NewDependency(c interface{}, ops ...dig.ProvideOption) dependency {
	return dependency{
		constructor: c,
		ops:         ops,
	}
}

func NewModule(deps ...dependency) Module {
	return deps
}

func NewContainer() *Container {
	c := &Container{
		c: dig.New(),
	}

	err := c.c.Provide(appAdapter)
	if err != nil {
		panic(err)
	}

	return c
}

func (c *Container) Provide(mods ...Module) *Container {
	for _, m := range mods {
		for _, d := range m {
			err := c.c.Provide(d.constructor, d.ops...)
			if err != nil {
				panic(err)
			}
		}
	}

	return c
}

func (c *Container) Invoke(fs ...interface{}) *Container {
	for _, f := range fs {
		err := c.c.Invoke(f)
		if err != nil {
			panic(err)
		}
	}
	return c
}

func (c *Container) Run() {
	c.Invoke(main)
}
