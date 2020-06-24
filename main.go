package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/lifeform/simplecell"
	"github.com/Burmuley/life/ui"
	"github.com/Burmuley/life/world"
	"github.com/Burmuley/life/world/convey"
)

func fillWorld(w world.Explorer) {
	rand.Seed(time.Now().UnixNano())
	n := 200000
	maxR, maxC := w.Size()

	for col := 0; col < maxC; col++ {
		for row := 0; row < maxR; row++ {
			l := world.Location{
				Row: row,
				Col: col,
			}
			rnd := rand.Intn(n)
			state := lifeform.DEAD

			if rnd > n/2 {
				state = lifeform.ALIVE
			}

			w.SetLife(simplecell.New(state), l)
		}
	}
}

var cmdui string

func main() {
	var appUi ui.UI
	flag.StringVar(&cmdui, "ui", "console", "choose UI: 'console' or 'gui'")
	flag.Parse()

	conveyWorld := convey.NewFilled(30, 50, fillWorld)

	fabric := ui.NewFabric()

	switch cmdui {
	case "console":
		appUi = fabric.Get("Console")
	case "gui":
		appUi = fabric.Get("Fyne")
	default:
		log.Fatal(fmt.Sprintf("Unknown UI provided: %s", cmdui))
	}

	appUi.SetWorld(conveyWorld)
	appUi.Run()
}
