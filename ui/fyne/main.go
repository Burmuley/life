package fyne

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Burmuley/life/world"
)

const (
	cellSize = 8
)

type UI struct {
	widget.BaseWidget
	world world.Explorer
	name  string
}

func (u *UI) SetWorld(w world.Explorer) {
	u.world = w
}

func (u *UI) Name() string {
	return "Fyne"
}

func (u *UI) Run() {
	window := app.New().NewWindow("Convey's Game of Life")
	window.SetFixedSize(true)
	maxlayout := layout.NewMaxLayout()
	cont := fyne.NewContainerWithLayout(maxlayout, u)
	maxR, maxC := u.world.Size()
	cont.Resize(fyne.NewSize(maxR*cellSize, maxC*cellSize))
	window.SetContent(cont)

	go func() {
		for {
			u.world.CheckAll()
			u.world.UpdateAll()
			u.Refresh()
			time.Sleep(time.Millisecond * 150)
		}
	}()

	window.ShowAndRun()
}

func (u *UI) CreateRenderer() fyne.WidgetRenderer {
	renderer := &worldRenderer{world: u}

	render := canvas.NewRaster(renderer.draw)
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}
	renderer.ApplyTheme()
	return renderer
}

func (u *UI) pixelDensity() float64 {
	c := fyne.CurrentApp().Driver().CanvasForObject(u)
	if c == nil {
		return 1.0
	}

	pixW, _ := c.PixelCoordinateForPosition(fyne.NewPos(cellSize, cellSize))
	return float64(pixW) / float64(cellSize)
}

func New() *UI {
	ui := &UI{name: "Fyne"}
	ui.ExtendBaseWidget(ui)
	return ui
}
