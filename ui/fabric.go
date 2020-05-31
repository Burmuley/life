package ui

import "github.com/Burmuley/life/ui/console"

type Fabric struct {
	uis map[string]UI
}

func NewFabric() *Fabric {
	f := &Fabric{make(map[string]UI, 0)}
	f.Add(console.New())
	return f
}

func (f *Fabric) Get(ui string) UI {
	return f.uis[ui]
}

func (f *Fabric) Add(ui UI) {
	f.uis[ui.Name()] = ui
}

func (f *Fabric) List() []string {
	l := make([]string, len(f.uis))
	i := 0
	for u := range f.uis {
		l[i] = u
	}

	return l
}
