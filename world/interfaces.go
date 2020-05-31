package world

import "github.com/Burmuley/life/lifeform"

type Location struct {
	Row int
	Col int
}

type Maker interface {
	SetLife(lifeform.Shaper, Location) // set a life form to the location in the world
}

type Checker interface {
	Check(Location) lifeform.State // Checks life form at the location for neighbors and sets/returns next state
	CheckAll()
}

type Processor interface {
	Update(Location) // runs against all cells in the world and updates current state to the next
	UpdateAll()
}

type Informer interface {
	Get(Location) lifeform.Shaper // get life form from the location
	Size() (int, int)             // returns size of the world
}

type CheckInformer interface {
	Checker
	Informer
}

type Explorer interface {
	Informer
	Checker
	Maker
	Processor
}
