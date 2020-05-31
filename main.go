package main

import (
	"math/rand"
	"time"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/lifeform/simplecell"
	"github.com/Burmuley/life/ui"
	"github.com/Burmuley/life/world"
)

func fillWorld(w world.WholeWorld) {
	rand.Seed(time.Now().UnixNano())
	n := 20000
	maxX, maxY := w.Size()

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			l := world.Location{x, y}
			r := rand.Intn(n)
			s := lifeform.DEAD

			if r > n/2 {
				s = lifeform.ALIVE
			}

			w.SetLife(simplecell.New(s), l)
		}
	}
}

func main() {
	convey := world.NewConvey(30, 50)
	fillWorld(convey)

	fabric := ui.NewFabric()
	ui := fabric.Get("Console")
	ui.SetWorld(convey)
	ui.Run()
}
