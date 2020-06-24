package convey

import (
	"github.com/Burmuley/life/lifeform"
	"github.com/Burmuley/life/world"
)

type Convey struct {
	grid         [][]lifeform.Shaper
	minNeighbors int
}

func (c *Convey) Size() (x, y int) {
	if len(c.grid) == 0 {
		return 0, 0
	}

	x = len(c.grid[0])
	y = len(c.grid)
	return
}

func (c *Convey) Get(l world.Location) lifeform.Shaper {
	return c.grid[l.Col][l.Row]
}

func (c *Convey) Update(l world.Location) {
	c.grid[l.Col][l.Row].Update()
}

func (c *Convey) UpdateAll() {
	for _, i := range c.grid {
		for _, k := range i {
			k.Update()
		}
	}
}

func (c *Convey) Check(l world.Location) lifeform.State {
	maxR, maxC := c.Size()                  // size of the world
	neighbors := make([]lifeform.Shaper, 0) // list of neighbors
	state := lifeform.DEAD                  // default state
	cell := c.grid[l.Col][l.Row]

	// iterate through all neighbors of the Location
	startR := l.Row - 1
	startC := l.Col - 1

	// check for min boundaries
	if startR < 0 {
		startR = l.Row
	}

	if startC < 0 {
		startC = l.Col
	}

	for col := startC; col <= l.Col+1 && col < maxC; col++ {
		for row := startR; row <= l.Row+1 && row < maxR; row++ {
			if row == l.Row && col == l.Col {
				continue
			}

			// only add ALIVE neighbors
			tCell := c.grid[col][row]
			if tCell.State() == lifeform.ALIVE {
				neighbors = append(neighbors, tCell)
			}
		}
	}

	n := len(neighbors)

	// A dead cell with exactly three live neighbors becomes a live cell (birth).
	// A live cell with two or three live neighbors stays alive (survival).
	// In all other cases, a cell dies or remains dead (overcrowding or loneliness).
	if cell.State() == lifeform.DEAD && n == c.minNeighbors {
		state = lifeform.ALIVE
	}

	if cell.State() == lifeform.ALIVE && (n == c.minNeighbors || n == c.minNeighbors-1) {
		state = lifeform.ALIVE
	}

	cell.SetNext(state)
	return state
}

func (c *Convey) CheckAll() {
	for col := range c.grid {
		for row := range c.grid[col] {
			c.Check(world.Location{row, col})
		}
	}
}

func (c *Convey) SetLife(lf lifeform.Shaper, l world.Location) {
	c.grid[l.Col][l.Row] = lf
}

func New(row, col int) *Convey {
	convey := &Convey{
		grid:         make([][]lifeform.Shaper, col),
		minNeighbors: 3,
	}

	for i := 0; i < col; i++ {
		convey.grid[i] = make([]lifeform.Shaper, row)
	}
	return convey
}

func NewFilled(row, col int, filler func(explorer world.Explorer)) *Convey {
	convey := New(row, col)
	filler(convey)

	return convey
}
