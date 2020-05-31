package world

import "github.com/Burmuley/life/lifeform"

type Location struct {
	X int
	Y int
}

type Maker interface {
	SetLife(lifeform.Shaper, Location) // set a life form to the location in the world
}

type Checker interface {
	Check(Location) lifeform.State // Checks life form at the location for neighbors and sets/returns next state
}

type Processor interface {
	Update() // runs against all cells in the world and updates current state to the next
}

type Informer interface {
	Get(Location) lifeform.Shaper // get life form from the location
	Size() (int, int)             // returns size of the world
}

type CheckInformer interface {
	Checker
	Informer
}
