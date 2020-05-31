package ui

import (
	"github.com/Burmuley/life/ui/console"
	"github.com/Burmuley/life/ui/fyne"
)

type Fabric struct {
	uis map[string]func() UI
}

func NewFabric() *Fabric {
	f := &Fabric{make(map[string]func() UI, 0)}
	// fill known UIs
	f.Add("Console", console.New)
	f.Add("Fyne", fyne.New)
	return f
}

func (f *Fabric) Get(ui string) UI {
	return f.uis[ui]()
}

func (f *Fabric) Add(ui string, uiFunc func() UI) {
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
