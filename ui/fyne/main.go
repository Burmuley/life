package fyne

import (
	"github.com/Burmuley/life/world"
)

type UI struct {
	world world.Explorer
	name  string
}

func (u *UI) SetWorld(w world.Explorer) {
	panic("implement me")
}

func (u *UI) Name() string {
	panic("implement me")
}

func (u *UI) Run() {
	panic("implement me")
}

func New() *UI {
	return &UI{name: "Fyne"}
}
