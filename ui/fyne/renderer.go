package fyne

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/world"
)

type worldRenderer struct {
	render  *canvas.Raster
	objects []fyne.CanvasObject
	img     *image.RGBA

	aliveColor color.Color
	deadColor  color.Color

	world *UI
}

func (w *worldRenderer) Layout(s fyne.Size) {
	w.render.Resize(s)
}

func (w *worldRenderer) MinSize() fyne.Size {
	wd, ht := w.world.world.Size()
	pixDensity := w.world.pixelDensity()
	wd = int(float64(wd*cellSize) / pixDensity)
	ht = int(float64(ht*cellSize) / pixDensity)
	return fyne.NewSize(ht, wd)
}

func (w *worldRenderer) Refresh() {
	canvas.Refresh(w.render)
}

func (w *worldRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (w *worldRenderer) Objects() []fyne.CanvasObject {
	return w.objects
}

func (w *worldRenderer) Destroy() {
}

func (w *worldRenderer) ApplyTheme() {
	w.aliveColor = theme.TextColor()
	w.deadColor = theme.BackgroundColor()
}

func (w *worldRenderer) draw(width, height int) image.Image {
	img := w.img
	if img == nil || img.Bounds().Size().X != width || img.Bounds().Size().Y != height {
		img = image.NewRGBA(image.Rect(0, 0, height, width))
		w.img = img
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cellState := w.getCoordState(x, y)
			img.Set(x, y, w.getStateColor(cellState))
		}
	}

	return img
}

func (w *worldRenderer) getStateColor(state lifeform.State) color.Color {
	switch state {
	case lifeform.ALIVE:
		return w.aliveColor
	case lifeform.DEAD:
		return w.deadColor
	}

	return nil
}

func (w *worldRenderer) getCoordState(x, y int) lifeform.State {
	maxC, maxR := w.world.world.Size()
	maxX := w.img.Bounds().Size().X
	maxY := w.img.Bounds().Size().Y
	dx := math.Round(float64(maxX) / float64(maxC))
	dy := math.Round(float64(maxY) / float64(maxR))

	col := int(x / int(dx))
	row := int(y / int(dy))

	return w.world.world.Get(world.Location{
		Row: row,
		Col: col,
	}).State()
}
