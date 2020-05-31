package lifeform

type State int

const (
	ALIVE State = iota + 1
	DEAD
)

type Shaper interface {
	State() State  // returns current state of the life form
	SetNext(State) // sets next state for the lifeform
	Update()       // updates current state to the next state
}
