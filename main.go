package main

import (
	"fmt"

	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/world"
)

func StateToText(s lifeform.State) string {
	switch s {
	case lifeform.ALIVE:
		return "ALIVE"
	case lifeform.DEAD:
		return "DEAD"
	default:
		return "UNKNOWN"
	}
}

func printWorld(w world.Informer) {
	x, y := w.Size()
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			c := w.Get(world.Location{dx, dy})
			fmt.Printf("%v\t", StateToText(c.State()))
		}
		fmt.Println()
	}
}

func checkWorld(w world.CheckInformer) {
	x, y := w.Size()
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			w.Check(world.Location{dx, dy})
		}
	}
}

func main() {
	convey := world.NewConvey(3, 3)
	_, y := convey.Size()

	for dy := 0; dy < y; dy++ {
		convey.SetLife(lifeform.NewSimpleCell(lifeform.ALIVE), world.Location{1, dy})
		convey.SetLife(lifeform.NewSimpleCell(lifeform.DEAD), world.Location{0, dy})
		convey.SetLife(lifeform.NewSimpleCell(lifeform.DEAD), world.Location{2, dy})
	}

	printWorld(convey)
	fmt.Println()

	for i := 0; i < 3; i++ {
		checkWorld(convey)
		convey.Update()
		printWorld(convey)
		fmt.Println()
	}
}
