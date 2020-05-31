package world

import "github.com/Burmuley/life/lifeform"

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

func (c *Convey) Get(l Location) lifeform.Shaper {
	return c.grid[l.Y][l.X]
}

func (c *Convey) Update(l Location) {
	c.grid[l.Y][l.X].Update()
}

func (c *Convey) UpdateAll() {
	for _, i := range c.grid {
		for _, k := range i {
			k.Update()
		}
	}
}

func (c *Convey) Check(l Location) lifeform.State {
	maxX, maxY := c.Size()                  // size of the world
	neighbors := make([]lifeform.Shaper, 0) // list of neighbors
	state := lifeform.DEAD                  // default state
	cell := c.grid[l.Y][l.X]

	// iterate through all neighbors of the Location
	sx := l.X - 1
	sy := l.Y - 1

	// check for min boundaries
	if sx < 0 {
		sx = l.X
	}

	if sy < 0 {
		sy = l.Y
	}

	for y := sy; y <= l.Y+1 && y < maxY; y++ {
		for x := sx; x <= l.X+1 && x < maxX; x++ {
			if x == l.X && y == l.Y {
				continue
			}

			// only add ALIVE neighbors
			tCell := c.grid[y][x]
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
	for y := range c.grid {
		for x := range c.grid[y] {
			c.Check(Location{x, y})
		}
	}
}

func (c *Convey) SetLife(lf lifeform.Shaper, l Location) {
	c.grid[l.Y][l.X] = lf
}

func NewConvey(x, y int) *Convey {
	convey := &Convey{
		grid:         make([][]lifeform.Shaper, y),
		minNeighbors: 3,
	}

	for i := 0; i < y; i++ {
		convey.grid[i] = make([]lifeform.Shaper, x)
	}
	return convey
}
