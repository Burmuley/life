package fyne

import (
	"github.com/Burmuley/life/ui"
	"github.com/Burmuley/life/world"
)

type UI struct {
	world world.WholeWorld
	name  string
}

func (u *UI) SetWorld(w world.WholeWorld) {
	panic("implement me")
}

func (u *UI) Name() string {
	panic("implement me")
}

func (u *UI) Run() {
	panic("implement me")
}

func New() ui.UI {
	return &UI{name: "Fyne"}
}
