package maze

import (
	"math/rand"
	"time"
)

func binary() {
	rand.Seed( time.Now().UTC().UnixNano())
	g := Grid2d{}
	g.initialize(9, 9)
	for i, row := range(g.grid) {
		for j, _ := range(row) {
			cell := g.getCell(i, j)
			neighbors := make([]*Cell, 0)
			north := g.getCell(cell.north.x, cell.north.y)
			if north != nil {
				neighbors = append(neighbors, north)
			}
			east := g.getCell(cell.east.x, cell.east.y)
			if east != nil {
				neighbors = append(neighbors, east)
			}
			if len(neighbors) > 0 {
				index := rand.Intn(len(neighbors))
				cell.link(neighbors[index], true)
			}
		}
	}
	g.display()
}
