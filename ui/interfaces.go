package ui

import "github.com/Burmuley/life/world"

type UI interface {
	SetWorld(world world.Explorer)
	Name() string
	Run()
}
