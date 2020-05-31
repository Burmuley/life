package world

import "github.com/Burmuley/life/lifeform"

type Convey struct {
	grid         [][]lifeform.Shaper
	minNeighbors int
	maxNeighbors int
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

func (c *Convey) Update() {
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

	// iterate through all neighbors of the Location
	sx := l.X - 1
	sy := l.Y - 1

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
			if c.grid[y][x].State() == lifeform.ALIVE {
				neighbors = append(neighbors, c.grid[y][x])
			}
		}
	}

	n := len(neighbors)

	if n >= c.minNeighbors && n < c.maxNeighbors {
		state = lifeform.ALIVE
	}

	c.grid[l.Y][l.X].SetNext(state)
	return state
}

func (c *Convey) SetLife(lf lifeform.Shaper, l Location) {
	c.grid[l.Y][l.X] = lf
}

func NewConvey(x, y int) *Convey {
	convey := &Convey{
		grid:         make([][]lifeform.Shaper, y),
		minNeighbors: 2,
		maxNeighbors: 3,
	}

	for i := 0; i < y; i++ {
		convey.grid[i] = make([]lifeform.Shaper, x)
	}
	return convey
}
