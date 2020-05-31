package console

import (
	"time"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/world"
	"github.com/rivo/tview"
)

type UI struct {
	world world.WholeWorld
	name  string
}

func New() *UI {
	return &UI{name: "Console"}
}

func (u *UI) SetWorld(w world.WholeWorld) {
	u.world = w
}

func (u *UI) Name() string {
	return u.name
}

func (u *UI) Run() {
	maxX, maxY := u.world.Size()

	rows := make([]int, maxX)
	for r := range rows {
		rows[r] = 1
	}

	cols := make([]int, maxY)
	for c := range cols {
		cols[c] = 1
	}

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(false)
	app.SetRoot(table, true)

	UpdateTable(table, u.world)

	go func() {
		for {
			u.world.CheckAll()
			u.world.UpdateAll()
			UpdateTable(table, u.world)
			time.Sleep(time.Millisecond * 150)
			app.Draw()
		}
	}()

	app.Run()
}

func UpdateTable(t *tview.Table, w world.Informer) {
	maxX, maxY := w.Size()
	for dy := 0; dy < maxY; dy++ {
		for dx := 0; dx < maxX; dx++ {
			l := world.Location{dx, dy}
			t.SetCellSimple(dx, dy, GetSymbol(w.Get(l)))
		}
	}
}

func GetSymbol(c lifeform.Shaper) string {
	syms := map[int]string{
		1: "\u2588",
		2: " ",
	}

	return syms[int(c.State())]
}
