package ui

import (
	"github.com/Burmuley/life/ui/console"
	"github.com/Burmuley/life/ui/fyne"
)

const (
	uiConsole string = "Console"
	uiFyne    string = "Fyne"
)

func getUI(name string) UI {
	switch name {
	case uiConsole:
		return console.New()
	case uiFyne:
		return fyne.New()
	default:
		return nil

	}
}

type fabricFunc func(string) UI

type Fabric struct {
	uis map[string]fabricFunc
}

func NewFabric() *Fabric {
	f := &Fabric{make(map[string]fabricFunc, 0)}
	// fill known UIs
	f.Add(uiConsole, getUI)
	f.Add(uiFyne, getUI)
	return f
}

func (f *Fabric) Get(ui string) UI {
	return f.uis[ui](ui)
}

func (f *Fabric) Add(ui string, uiFunc fabricFunc) {
	f.uis[ui] = uiFunc
}

func (f *Fabric) List() []string {
	l := make([]string, len(f.uis))
	i := 0
	for u := range f.uis {
		l[i] = u
		i++
	}

	return l
}
