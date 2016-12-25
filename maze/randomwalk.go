package maze

import (
	"math/rand"
	"time"
)

func RandomWalk(x int, y int) Grid2d {
	rand.Seed(time.Now().UTC().UnixNano())
	g := Grid2d{}
	g.initialize(x, y)

	cell := g.randomCell()
	unvisited := g.size() - 1
	for unvisited > 0 {
		x := rand.Intn(4)
		var neighbor *Cell
		switch x {
		case 0:
			neighbor = g.GetCell(cell.north.x, cell.north.y)
		case 1:
			neighbor = g.GetCell(cell.east.x, cell.east.y)
		case 2:
			neighbor = g.GetCell(cell.south.x, cell.south.y)
		case 3:
			neighbor = g.GetCell(cell.west.x, cell.west.y)
		}

		if neighbor != nil {
			if len(neighbor.links) < 1 {
				cell.link(neighbor, true)
				unvisited -= 1
			}

			cell = neighbor
		}
	}

	return g
}