package simplecell

import "github.com/Burmuley/life/lifeform"

type SimpleCell struct {
	cState lifeform.State // current state
	nState lifeform.State // next state
}

func (c *SimpleCell) State() lifeform.State {
	return c.cState
}

func (c *SimpleCell) SetNext(s lifeform.State) {
	c.nState = s
}

func (c *SimpleCell) Update() {
	c.cState = c.nState
}

func New(state lifeform.State) *SimpleCell {
	return &SimpleCell{cState: state}
}
