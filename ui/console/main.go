package console

import (
	"log"
	"time"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/world"
	"github.com/rivo/tview"
)

type UI struct {
	world world.Explorer
	name  string
}

func New() *UI {
	return &UI{name: "Console"}
}

func (u *UI) SetWorld(w world.Explorer) {
	u.world = w
}

func (u *UI) Name() string {
	return u.name
}

func (u *UI) Run() {
	maxR, maxC := u.world.Size()

	rows := make([]int, maxR)
	for r := range rows {
		rows[r] = 1
	}

	cols := make([]int, maxC)
	for c := range cols {
		cols[c] = 1
	}

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(false)
	app.SetRoot(table, true)

	updateTable(table, u.world)

	go func() {
		for {
			u.world.CheckAll()
			u.world.UpdateAll()
			updateTable(table, u.world)
			time.Sleep(time.Millisecond * 150)
			app.Draw()
		}
	}()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func updateTable(t *tview.Table, w world.Informer) {
	maxR, maxC := w.Size()
	for col := 0; col < maxC; col++ {
		for row := 0; row < maxR; row++ {
			l := world.Location{row, col}
			t.SetCellSimple(row, col, getSymbol(w.Get(l)))
		}
	}
}

func getSymbol(c lifeform.Shaper) string {
	syms := map[int]string{
		int(lifeform.ALIVE): "\u2588",
		int(lifeform.DEAD):  " ",
	}

	return syms[int(c.State())]
}
