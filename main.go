package main

import (
	"math/rand"
	"time"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/lifeform/simplecell"
	"github.com/Burmuley/life/ui"
	"github.com/Burmuley/life/world"
)

func fillWorld(w world.Explorer) {
	rand.Seed(time.Now().UnixNano())
	n := 20000
	maxR, maxC := w.Size()

	for col := 0; col < maxC; col++ {
		for row := 0; row < maxR; row++ {
			l := world.Location{row, col}
			rnd := rand.Intn(n)
			state := lifeform.DEAD

			if rnd > n/2 {
				state = lifeform.ALIVE
			}

			w.SetLife(simplecell.New(state), l)
		}
	}
}

func main() {
	convey := world.NewConvey(30, 50)
	fillWorld(convey)

	fabric := ui.NewFabric()
	appUi := fabric.Get("Console")
	appUi.SetWorld(convey)
	appUi.Run()
}
