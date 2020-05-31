package lifeform

type SimpleCell struct {
	cState State // current state
	nState State // next state
}

func (c *SimpleCell) State() State {
	return c.cState
}

func (c *SimpleCell) SetNext(s State) {
	c.nState = s
}

func (c *SimpleCell) Update() {
	c.cState = c.nState
}

func NewSimpleCell(state State) *SimpleCell {
	return &SimpleCell{cState: state}
}
